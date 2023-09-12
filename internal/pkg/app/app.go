package app

import (
	"context"
	"crypto/tls"
	"dam4rus/nifi-tui/internal/pkg/nifiapi"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"
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
	waitGroup.Add(6)
	processGroupsChannel := make(chan *nifiapi.ProcessGroupsEntity, 1)
	processorsChannel := make(chan *nifiapi.ProcessorsEntity, 1)
	connectionsChannel := make(chan *nifiapi.ConnectionsEntity, 1)
	funnelsChannel := make(chan *nifiapi.FunnelsEntity, 1)
	inputPortsChannel := make(chan *nifiapi.InputPortsEntity, 1)
	outputPortsChannel := make(chan *nifiapi.OutputPortsEntity, 1)
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
		defer waitGroup.Done()
		connections, _, err := a.apiClient.ProcessGroupsAPI.GetConnections(ctx, processGroupId).Execute()
		if err != nil {
			errorChannel <- err
			return
		}
		connectionsChannel <- connections
	}(ctx)
	go func(ctx context.Context) {
		defer waitGroup.Done()
		funnels, _, err := a.apiClient.ProcessGroupsAPI.GetFunnels(ctx, processGroupId).Execute()
		if err != nil {
			errorChannel <- err
			return
		}
		funnelsChannel <- funnels
	}(ctx)
	go func(ctx context.Context) {
		defer waitGroup.Done()
		inputPorts, _, err := a.apiClient.ProcessGroupsAPI.GetInputPorts(ctx, processGroupId).Execute()
		if err != nil {
			errorChannel <- err
			return
		}
		inputPortsChannel <- inputPorts
	}(ctx)
	go func(ctx context.Context) {
		defer waitGroup.Done()
		outputPorts, _, err := a.apiClient.ProcessGroupsAPI.GetOutputPorts(ctx, processGroupId).Execute()
		if err != nil {
			errorChannel <- err
			return
		}
		outputPortsChannel <- outputPorts
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
		connections := <-connectionsChannel
		funnels := <-funnelsChannel
		inputPorts := <-inputPortsChannel
		outputPorts := <-outputPortsChannel
		a.app.QueueUpdateDraw(func() {
			frontPage, _ := a.pages.GetFrontPage()
			if frontPage != "Loading" {
				return
			}
			list := tview.NewList()
			list.SetTitle(processGroupId).
				SetTitleAlign(tview.AlignLeft).
				SetBorder(true).
				SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
					if processGroupId != "root" {
						if event.Key() == tcell.KeyESC {
							a.pages.RemovePage(processGroupId)
							return nil
						}
					}
					return event
				})

			for _, processGroup := range processGroups.ProcessGroups {
				processGroup := processGroup
				list.AddItem("📁"+processGroup.Component.GetName(), processGroup.Component.GetId(), 0, func() {
					a.EnterProcessGroup(*processGroup.Id)
				})
			}
			for _, processor := range processors.Processors {
				processor := processor
				list.AddItem("⊞"+processor.Component.GetName(), processor.Component.GetId(), 0, func() {
					a.onProcessorSelected(&processor)
				})
			}
			for _, funnel := range funnels.Funnels {
				funnel := funnel
				list.AddItem("Ⴤ", funnel.Component.GetId(), 0, nil)
			}
			for _, inputPort := range inputPorts.InputPorts {
				inputPort := inputPort
				list.AddItem("⇥"+inputPort.Component.GetName(), inputPort.Component.GetId(), 0, nil)
			}
			for _, outputPort := range outputPorts.OutputPorts {
				outputPort := outputPort
				list.AddItem("↦"+outputPort.Component.GetName(), outputPort.Component.GetId(), 0, nil)
			}

			findComponent := func(componentType, id string) Component {
				switch componentType {
				case "PROCESSOR":
					for _, processor := range processors.Processors {
						if processor.Component.GetId() == id {
							return processor.Component
						}
					}
				case "FUNNEL":
					for _, funnel := range funnels.Funnels {
						if funnel.Component.GetId() == id {
							return &FunnelComponent{
								Id: funnel.Component.GetId(),
							}
						}
					}
				case "INPUT_PORT":
					for _, inputPort := range inputPorts.InputPorts {
						if inputPort.Component.GetId() == id {
							return inputPort.Component
						}
					}
				case "OUTPUT_PORT":
					for _, outputPort := range outputPorts.OutputPorts {
						if outputPort.Component.GetId() == id {
							return outputPort.Component
						}
					}
				}
				return nil
			}

			for _, connection := range connections.Connections {
				connection := connection
				sourceComponent := findComponent(connection.GetSourceType(), connection.GetSourceId())
				destComponent := findComponent(connection.GetDestinationType(), connection.GetDestinationId())
				if sourceComponent == nil || destComponent == nil {
					continue
				}
				addConnectionItem(list, sourceComponent, destComponent)
			}
			a.pages.RemovePage("Loading")
			a.pages.AddAndSwitchToPage(processGroupId, list, true)
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
	a.pages.AddAndSwitchToPage("Loading", loadingPage, true)
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
		app:                  a,
		processor:            processor,
		pages:                tview.NewPages(),
		propertyListView:     tview.NewList(),
		relationshipsFlex:    tview.NewFlex(),
		relationshipListView: tview.NewList(),
	}
	processorDetailsScreen.propertyListView.SetTitle(*processor.Component.Name).
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true).
		SetInputCapture(processorDetailsScreen.handleInput)

	for k, v := range *processor.Component.Config.Properties {
		k, v := k, v
		descriptor := (*processor.Component.Config.Descriptors)[k]
		displayName := *descriptor.DisplayName
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
		processorDetailsScreen.propertyListView.AddItem(displayName, value, 0, func() { processorDetailsScreen.onPropertySelected(k) })
	}

	processorDetailsScreen.relationshipListView.SetTitle("Relationships").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)
	for i, relationship := range processor.Component.Relationships {
		i, relationship := i, relationship
		processorDetailsScreen.relationshipListView.AddItem(relationship.GetName(), buildRelationshipSecondaryText(relationship), 0, func() {
			processorDetailsScreen.onRelationshipSelected(i)
		})
	}

	rightColumn := tview.NewFlex().
		SetDirection(tview.FlexRow)
	rightColumn.SetTitle("Retry").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)
	processorDetailsScreen.relationshipsFlex.
		SetDirection(tview.FlexColumn).
		AddItem(processorDetailsScreen.relationshipListView, 0, 1, true).
		AddItem(rightColumn, 0, 1, false)
	processorDetailsScreen.relationshipsFlex.SetInputCapture(processorDetailsScreen.handleInput)
	processorDetailsScreen.pages.
		AddAndSwitchToPage("Properties", processorDetailsScreen.propertyListView, true).
		AddPage("Relationships", processorDetailsScreen.relationshipsFlex, true, false)

	a.pages.AddAndSwitchToPage(*processor.Id, processorDetailsScreen.pages, true)
}

type ProcessorDetailsScreen struct {
	app                  *App
	processor            *nifiapi.ProcessorEntity
	pages                *tview.Pages
	propertyListView     *tview.List
	relationshipsFlex    *tview.Flex
	relationshipListView *tview.List
	modified             bool
}

func (s *ProcessorDetailsScreen) save() {
	autoTerminatedRelationships := []string{}
	retriedRelationships := []string{}
	for _, relationship := range s.processor.Component.Relationships {
		if relationship.GetAutoTerminate() {
			autoTerminatedRelationships = append(autoTerminatedRelationships, *relationship.Name)
		}
		if relationship.GetRetry() {
			retriedRelationships = append(retriedRelationships, *relationship.Name)
		}
	}
	processor := nifiapi.ProcessorEntity{
		Component: &nifiapi.ProcessorDTO{
			Id: s.processor.Id,
			Config: &nifiapi.ProcessorConfigDTO{
				Properties:                  s.processor.Component.Config.Properties,
				AutoTerminatedRelationships: autoTerminatedRelationships,
				RetriedRelationships:        retriedRelationships,
			},
		},
		Revision: s.processor.Revision,
	}
	updatedProcessor, response, err := s.app.apiClient.ProcessorsAPI.UpdateProcessor(context.Background(), *processor.Component.Id).
		Body(processor).
		Execute()
	if err != nil {
		body, readErr := io.ReadAll(response.Body)
		modal := tview.NewModal().
			AddButtons([]string{"Ok"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				s.pages.RemovePage("Error")
			})
		modal.SetTitle("Error").
			SetBorder(true)
		if readErr == nil {
			modal.SetText(string(body))
		} else {
			modal.SetText(err.Error())
		}
		s.pages.AddPage("Error", modal, true, true)
		return
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
					s.pages.RemovePage("Modal")
				})
			s.pages.AddPage("Modal", modal, true, true)
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
	switch event.Rune() {
	case '1':
		s.pages.SwitchToPage("Properties")
		return nil
	case '2':
		s.pages.SwitchToPage("Relationships")
		return nil
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
		if *(*s.processor.Component.Config.Properties)[key] == "true" {
			s.setProperty(key, "false")
		} else {
			s.setProperty(key, "true")
		}
		return
	}
	var newValue string
	if value := (*s.processor.Component.Config.Properties)[key]; value != nil {
		newValue = *value
	} else {
		newValue = ""
	}

	form := tview.NewForm()

	var gridHeight int
	if len(allowableValues) > 0 {
		var allowableValueDisplayNames []string
		for _, allowableValue := range propertyDescriptor.AllowableValues {
			allowableValueDisplayNames = append(allowableValueDisplayNames, *allowableValue.AllowableValue.DisplayName)
		}
		form.AddDropDown(key, allowableValueDisplayNames, slices.Index(allowableValues, newValue), func(option string, optionIndex int) {
			if optionIndex >= 0 {
				s.setProperty(key, allowableValues[optionIndex])
				s.pages.RemovePage("Property")
			}
		})

		gridHeight = 5
	} else {
		form.AddTextArea(key, newValue, 100, 5, 0, func(text string) {
			newValue = text
		}).
			AddButton("Save", func() {
				s.setProperty(key, newValue)
				s.pages.RemovePage("Property")
			}).
			AddButton("Cancel", func() {
				s.pages.RemovePage("Property")
			})
		gridHeight = 11
	}

	form.SetBorder(true).
		SetTitle("Property").
		SetTitleAlign(tview.AlignLeft).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyEsc:
				s.pages.RemovePage("Property")
				return nil
			case tcell.KeyCtrlS:
				s.setProperty(key, newValue)
				s.pages.RemovePage("Property")
				return nil
			}
			return event
		})

	grid := tview.NewGrid().
		SetColumns(0, 100, 0).
		SetRows(0, gridHeight, 0).
		AddItem(form, 1, 1, 1, 1, 0, 0, true)

	s.pages.AddPage("Property", grid, true, true)
}

func (s *ProcessorDetailsScreen) onRelationshipSelected(index int) {
	relationship := s.processor.Component.Relationships[index]
	form := tview.NewForm().
		AddCheckbox("Auto-terminated", relationship.GetAutoTerminate(), func(checked bool) {
			relationship.SetAutoTerminate(checked)
		}).
		AddCheckbox("Retry", relationship.GetRetry(), func(checked bool) {
			relationship.SetRetry(checked)
		}).
		AddButton("Save", func() {
			s.setRelationship(index, relationship)
			s.pages.RemovePage("Relationship")
		}).
		AddButton("Cancel", func() {
			s.pages.RemovePage("Relationship")
		})
	form.SetBorder(true).
		SetTitle(*relationship.Name).
		SetTitleAlign(tview.AlignLeft).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyEsc:
				s.pages.RemovePage("Relationship")
				return nil
			case tcell.KeyCtrlS:
				s.setRelationship(index, relationship)
				s.pages.RemovePage("Relationship")
				return nil
			}
			return event
		})
	grid := tview.NewGrid().
		SetColumns(0, 100, 0).
		SetRows(0, 9, 0).
		AddItem(form, 1, 1, 1, 1, 0, 0, true)
	s.pages.AddPage("Relationship", grid, true, true)
}

func (s *ProcessorDetailsScreen) setProperty(key, newValue string) {
	(*s.processor.Component.Config.Properties)[key] = &newValue
	descriptor := (*s.processor.Component.Config.Descriptors)[key]
	displayName := descriptor.DisplayName
	var value string
	if allowableValue := findAllowableValueByValue(descriptor.AllowableValues, newValue); allowableValue != nil {
		value = *allowableValue.DisplayName
	} else {
		value = newValue
	}
	s.propertyListView.SetItemText(s.propertyListView.GetCurrentItem(), *displayName, value)
	s.modified = true
}

func (s *ProcessorDetailsScreen) setRelationship(index int, relationship nifiapi.RelationshipDTO) {
	s.processor.Component.Relationships[index] = relationship
	s.relationshipListView.SetItemText(s.relationshipListView.GetCurrentItem(), *relationship.Name, buildRelationshipSecondaryText(relationship))
	s.modified = true
}

func buildRelationshipSecondaryText(relationship nifiapi.RelationshipDTO) string {
	var relationShipProperties []string
	if relationship.GetAutoTerminate() {
		relationShipProperties = append(relationShipProperties, "Auto terminate")
	}
	if relationship.GetRetry() {
		relationShipProperties = append(relationShipProperties, "Retry")
	}
	return strings.Join(relationShipProperties, ", ")
}

func findAllowableValueByValue(allowableValues []nifiapi.AllowableValueEntity, value string) *nifiapi.AllowableValueDTO {
	for _, allowableValue := range allowableValues {
		if *allowableValue.AllowableValue.Value == value {
			return allowableValue.AllowableValue
		}
	}
	return nil
}

type FunnelComponent struct {
	Id string
}

func (f *FunnelComponent) GetId() string {
	return f.Id
}

func (f *FunnelComponent) GetName() string {
	return "Ⴤ"
}

type Component interface {
	GetId() string
	GetName() string
}

func addConnectionItem(list *tview.List, source, dest Component) {
	list.AddItem(source.GetName()+"->"+dest.GetName(), source.GetId()+"->"+dest.GetId(), 0, nil)
}
