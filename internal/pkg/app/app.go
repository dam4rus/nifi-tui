package app

import (
	"context"
	"crypto/tls"
	"dam4rus/nifi-tui/internal/pkg/nifiapi"
	"fmt"
	"io"
	"net/http"
	"slices"
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type App struct {
	app        *tview.Application
	pages      *tview.Pages
	apiClient  *nifiapi.APIClient
	cancelFunc context.CancelFunc
}

func NewApp(url, username, password string) *App {
	configuration := nifiapi.NewConfiguration()
	configuration.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	configuration.Servers = nifiapi.ServerConfigurations{
		nifiapi.ServerConfiguration{
			URL: url,
		},
	}
	apiClient := nifiapi.NewAPIClient(configuration)
	accessToken, _, err := apiClient.AccessAPI.CreateAccessToken(context.Background()).
		Username(username).
		Password(password).
		Execute()
	if err != nil {
		panic(err)
	}
	apiClient.GetConfig().AddDefaultHeader("Authorization", "Bearer "+accessToken)
	accessStatus, _, err := apiClient.AccessAPI.GetAccessStatus(context.Background()).Execute()
	if err != nil {
		panic(err)
	}
	fmt.Printf("logged in user: %v\n", *accessStatus.AccessStatus.Identity)

	app := tview.NewApplication()
	pages := tview.NewPages()
	application := &App{
		app:       app,
		pages:     pages,
		apiClient: apiClient,
	}

	app.SetRoot(pages, true)
	app.SetInputCapture(application.inputCapture)
	return application
}

func (a *App) Run() error {
	return a.app.Run()
}

func (a *App) EnterProcessGroup(processGroupId string) {
	loadingPage := a.enterLoadingScreen(processGroupId, "Loading process group")

	ctx, cancelFunc := context.WithCancel(context.Background())
	a.cancelFunc = cancelFunc
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	processGroupsChannel := make(chan *nifiapi.ProcessGroupsEntity, 1)
	processorsChannel := make(chan *nifiapi.ProcessorsEntity, 1)
	errorChannel := make(chan error)
	go func(ctx context.Context) {
		defer waitGroup.Done()
		processGroups, _, err := a.apiClient.ProcessGroupsAPI.GetProcessGroups(ctx, processGroupId).Execute()
		if err != nil {
			errorChannel <- err
			return
		}
		processGroupsChannel <- processGroups
	}(ctx)
	go func(ctx context.Context) {
		defer waitGroup.Done()
		processors, _, err := a.apiClient.ProcessGroupsAPI.GetProcessors(ctx, processGroupId).Execute()
		if err != nil {
			errorChannel <- err
			return
		}
		processorsChannel <- processors
	}(ctx)
	go func(ctx context.Context) {
		waitGroup.Wait()
		select {
		case err := <-errorChannel:
			a.app.QueueUpdateDraw(func() {
				frontPage, _ := a.pages.GetFrontPage()
				if frontPage != "Loading" {
					return
				}
				loadingPage.SetLabel(fmt.Sprint("Failed to load process group details", err))
			})
		default:
		}
		processGroups := <-processGroupsChannel
		processors := <-processorsChannel
		a.app.QueueUpdateDraw(func() {
			frontPage, _ := a.pages.GetFrontPage()
			if frontPage != "Loading" {
				return
			}
			list := tview.NewList()
			list.SetTitle(processGroupId).
				SetTitleAlign(tview.AlignLeft).
				SetBorder(true)

			for _, processGroup := range processGroups.ProcessGroups {
				processGroup := processGroup
				list.AddItem("📁"+*processGroup.Component.Name, *processGroup.Id, 0, nil)
			}
			for _, processor := range processors.Processors {
				processor := processor
				list.AddItem("🔄"+*processor.Component.Name, *processor.Id, 0, func() {
					a.onProcessorSelected(&processor)
				})
			}
			a.pages.RemovePage("Loading")
			a.pages.AddPage(processGroupId, list, true, true)
		})
	}(ctx)
}

func (a *App) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	frontPage, _ := a.pages.GetFrontPage()
	if frontPage == "root" {
		if event.Rune() == 'q' {
			a.app.Stop()
			return nil
		}
	}
	return event
}

func (a *App) enterLoadingScreen(title, label string) *tview.TextView {
	loadingPage := tview.NewTextView().
		SetLabel(label)
	loadingPage.SetTitle(title).
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)
	a.pages.AddPage("Loading", loadingPage, true, true)
	return loadingPage
}

func (a *App) onProcessorSelected(processor *nifiapi.ProcessorEntity) {
	loadingPage := a.enterLoadingScreen(*processor.Component.Name, "Loading processor details")

	go func(ctx context.Context) {
		processorDetails, _, err := a.apiClient.ProcessorsAPI.GetProcessor(ctx, *processor.Id).Execute()
		a.app.QueueUpdateDraw(func() {
			frontPage, _ := a.pages.GetFrontPage()
			if frontPage != "Loading" {
				return
			}
			if err != nil {
				loadingPage.SetLabel("Failed to load processor details")
				return
			}
			a.pages.RemovePage("Loading")
			a.enterProcessorDetailsScreen(processorDetails)
		})
	}(context.Background())
}

func (a *App) enterProcessorDetailsScreen(processor *nifiapi.ProcessorEntity) {
	processorDetailsScreen := &ProcessorDetailsScreen{
		app:        a,
		processor:  processor,
		properties: *processor.Component.Config.Properties,
		list:       tview.NewList(),
	}
	processorDetailsScreen.list.SetTitle(*processor.Component.Name).
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true).
		SetInputCapture(processorDetailsScreen.handleInput)

	for k, v := range processorDetailsScreen.properties {
		k, v := k, v
		descriptor := (*processor.Component.Config.Descriptors)[k]
		displayName := descriptor.DisplayName
		var value string
		if v != nil {
			if allowableValue := findAllowableValueByValue(descriptor.AllowableValues, *v); allowableValue != nil {
				value = *allowableValue.DisplayName
			}
			if value == "" {
				value = *v
			}
		} else {
			value = ""
		}
		processorDetailsScreen.list.AddItem(*displayName, value, 0, func() { processorDetailsScreen.onPropertySelected(k) })
	}

	a.pages.AddPage(*processor.Id, processorDetailsScreen.list, true, true)
}

type ProcessorDetailsScreen struct {
	app        *App
	processor  *nifiapi.ProcessorEntity
	properties map[string]*string
	list       *tview.List
	modified   bool
}

func (s *ProcessorDetailsScreen) save() {
	processor := nifiapi.ProcessorEntity{
		Component: &nifiapi.ProcessorDTO{
			Id: s.processor.Id,
			Config: &nifiapi.ProcessorConfigDTO{
				Properties: &s.properties,
			},
		},
		Revision: s.processor.Revision,
	}
	updatedProcessor, response, err := s.app.apiClient.ProcessorsAPI.UpdateProcessor(context.Background(), *processor.Component.Id).
		Body(processor).
		Execute()
	if err != nil {
		body, readErr := io.ReadAll(response.Body)
		if readErr == nil {
			panic(string(body))
		} else {
			panic(err)
		}
	}
	s.processor = updatedProcessor
	s.modified = false
}

func (s *ProcessorDetailsScreen) handleInput(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyCtrlS:
		s.save()
		return nil
	case tcell.KeyEsc:
		if s.modified {
			modal := tview.NewModal().
				SetText("Processor modifications has not been saved!").
				AddButtons([]string{"Save", "Don't save", "Cancel"}).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					switch buttonIndex {
					case 0:
						s.save()
						s.app.pages.RemovePage(*s.processor.Id)
					case 1:
						s.app.pages.RemovePage(*s.processor.Id)
					case 2:
						// do nothing
					}
					s.app.pages.RemovePage("Modal")
				})
			s.app.pages.AddPage("Modal", modal, true, true)
			return nil
		} else {
			if s.app.cancelFunc != nil {
				s.app.cancelFunc()
				s.app.cancelFunc = nil
			}
			s.app.pages.RemovePage(*s.processor.Id)
			return nil
		}
	}
	return event
}

func (s *ProcessorDetailsScreen) onPropertySelected(key string) {
	propertyDescriptor := (*s.processor.Component.Config.Descriptors)[key]
	var allowableValues []string
	for _, allowableValue := range propertyDescriptor.AllowableValues {
		allowableValues = append(allowableValues, *allowableValue.AllowableValue.Value)
	}
	if len(allowableValues) == 2 && slices.Contains(allowableValues, "true") && slices.Contains(allowableValues, "false") {
		if *s.properties[key] == "true" {
			s.setProperty(key, "false")
		} else {
			s.setProperty(key, "true")
		}
		return
	}
	newValue := *s.properties[key]
	form := tview.NewForm()

	var gridHeight int
	if len(allowableValues) > 0 {
		var allowableValueDisplayNames []string
		for _, allowableValue := range propertyDescriptor.AllowableValues {
			allowableValueDisplayNames = append(allowableValueDisplayNames, *allowableValue.AllowableValue.DisplayName)
		}
		form.AddDropDown(key, allowableValueDisplayNames, slices.Index(allowableValues, *s.properties[key]), func(option string, optionIndex int) {
			if optionIndex >= 0 {
				s.setProperty(key, allowableValues[optionIndex])
				s.app.pages.RemovePage("Property")
			}
		})

		gridHeight = 5
	} else {
		form.AddTextArea(key, *s.properties[key], 100, 5, 0, func(text string) {
			newValue = text
		}).
			AddButton("Save", func() {
				s.setProperty(key, newValue)
				s.app.pages.RemovePage("Property")
			}).
			AddButton("Cancel", func() {
				s.app.pages.RemovePage("Property")
			})
		gridHeight = 11
	}

	form.SetBorder(true).
		SetTitle("Property").
		SetTitleAlign(tview.AlignLeft).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyEsc:
				s.app.pages.RemovePage("Property")
				return nil
			case tcell.KeyCtrlS:
				s.setProperty(key, newValue)
				s.app.pages.RemovePage("Property")
				return nil
			}
			return event
		})

	grid := tview.NewGrid().
		SetColumns(0, 100, 0).
		SetRows(0, gridHeight, 0).
		AddItem(form, 1, 1, 1, 1, 0, 0, true)

	s.app.pages.AddPage("Property", grid, true, true)
}

func (s *ProcessorDetailsScreen) setProperty(key, newValue string) {
	s.properties[key] = &newValue
	descriptor := (*s.processor.Component.Config.Descriptors)[key]
	displayName := descriptor.DisplayName
	var value string
	if allowableValue := findAllowableValueByValue(descriptor.AllowableValues, newValue); allowableValue != nil {
		value = *allowableValue.DisplayName
	} else {
		value = newValue
	}
	s.list.SetItemText(s.list.GetCurrentItem(), *displayName, value)
	s.modified = true
}

func findAllowableValueByValue(allowableValues []nifiapi.AllowableValueEntity, value string) *nifiapi.AllowableValueDTO {
	for _, allowableValue := range allowableValues {
		if *allowableValue.AllowableValue.Value == value {
			return allowableValue.AllowableValue
		}
	}
	return nil
}
