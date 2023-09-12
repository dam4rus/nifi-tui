package app

import (
	"context"
	"crypto/tls"
	"dam4rus/nifi-tui/internal/pkg/nifiapi"
	"fmt"
	"net/http"

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

	proxy := a.newProcessGroupProxy(processGroupId)
	a.cancelFunc = proxy.cancelFunc
	proxy.waitGroup.Add(6)
	processGroupsChannel := proxy.getProcessGroups()
	processorsChannel := proxy.getProcessors()
	connectionsChannel := proxy.getConnections()
	funnelsChannel := proxy.getFunnels()
	inputPortsChannel := proxy.getInputPorts()
	outputPortsChannel := proxy.getOutputPorts()

	go func(ctx context.Context) {
		proxy.waitGroup.Wait()
		select {
		case err := <-proxy.errorChannel:
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
			a.buildAndShowProcessGroupPage(processGroupComponents{
				id:            processGroupId,
				processGroups: processGroups,
				processors:    processors,
				connections:   connections,
				funnels:       funnels,
				inputPorts:    inputPorts,
				outputPorts:   outputPorts,
			})
		})
	}(proxy.ctx)
}

func (a *App) buildAndShowProcessGroupPage(components processGroupComponents) {
	frontPage, _ := a.pages.GetFrontPage()
	if frontPage != "Loading" {
		return
	}
	list := tview.NewList()
	list.SetTitle(components.id).
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if components.id != "root" {
				if event.Key() == tcell.KeyESC {
					a.pages.RemovePage(components.id)
					return nil
				}
			}
			return event
		})

	for _, processGroup := range components.processGroups.ProcessGroups {
		processGroup := processGroup
		list.AddItem("📁"+processGroup.Component.GetName(), processGroup.Component.GetId(), 0, func() {
			a.EnterProcessGroup(*processGroup.Id)
		})
	}
	for _, processor := range components.processors.Processors {
		processor := processor
		list.AddItem("⊞"+processor.Component.GetName(), processor.Component.GetId(), 0, func() {
			a.onProcessorSelected(&processor)
		})
	}
	for _, funnel := range components.funnels.Funnels {
		funnel := funnel
		list.AddItem("Ⴤ", funnel.Component.GetId(), 0, nil)
	}
	for _, inputPort := range components.inputPorts.InputPorts {
		inputPort := inputPort
		list.AddItem("⇥"+inputPort.Component.GetName(), inputPort.Component.GetId(), 0, nil)
	}
	for _, outputPort := range components.outputPorts.OutputPorts {
		outputPort := outputPort
		list.AddItem("↦"+outputPort.Component.GetName(), outputPort.Component.GetId(), 0, nil)
	}

	for _, connection := range components.connections.Connections {
		connection := connection
		sourceComponent := components.findComponent(connection.GetSourceType(), connection.GetSourceId())
		destComponent := components.findComponent(connection.GetDestinationType(), connection.GetDestinationId())
		if sourceComponent == nil || destComponent == nil {
			continue
		}
		list.AddItem(sourceComponent.GetName()+"->"+destComponent.GetName(), sourceComponent.GetId()+"->"+destComponent.GetId(), 0, nil)
	}
	a.pages.RemovePage("Loading")
	a.pages.AddAndSwitchToPage(components.id, list, true)
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
	a.pages.AddAndSwitchToPage(*processor.Id, newProcessorDetailsScreen(a, processor).pages, true)
}
