package nifitui

import (
	"context"
	"dam4rus/nifi-tui/internal/pkg/nifiapi"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type processGroupScreen struct {
	app            *App
	processGroupId string
}

func (pgs *processGroupScreen) enter() {
	frontPage, _ := pgs.app.pages.GetFrontPage()
	if frontPage == pgs.processGroupId {
		pgs.app.pages.RemovePage(pgs.processGroupId)
	}
	pgs.loadComponents()
}

func (pgs *processGroupScreen) loadComponents() {
	loadingPage := pgs.app.enterLoadingScreen(pgs.processGroupId, "Loading process group")

	proxy := pgs.app.newProcessGroupProxy(pgs.processGroupId)
	pgs.app.cancelFunc = proxy.cancelFunc
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
			pgs.app.app.QueueUpdateDraw(func() {
				frontPage, _ := pgs.app.pages.GetFrontPage()
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
		pgs.app.app.QueueUpdateDraw(func() {
			pgs.buildAndShowProcessGroupPage(processGroupComponents{
				id:            pgs.processGroupId,
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

func (pgs *processGroupScreen) buildAndShowProcessGroupPage(components processGroupComponents) {
	frontPage, _ := pgs.app.pages.GetFrontPage()
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
					pgs.app.pages.RemovePage(components.id)
					return nil
				}
			}
			return event
		})

	for _, processGroup := range components.processGroups.ProcessGroups {
		processGroup := processGroup
		list.AddItem("📁"+processGroup.Component.GetName(), processGroup.Component.GetId(), 0, func() {
			childProcessGroupScreen := processGroupScreen{
				app:            pgs.app,
				processGroupId: *processGroup.Id,
			}
			childProcessGroupScreen.enter()
		})
	}
	for _, processor := range components.processors.Processors {
		processor := processor
		list.AddItem("⊞"+processor.Component.GetName(), processor.Component.GetId(), 0, func() {
			pgs.onProcessorSelected(&processor)
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
	pgs.app.pages.RemovePage("Loading")
	pgs.app.pages.AddAndSwitchToPage(components.id, list, true)
}

func (pgs *processGroupScreen) onProcessorSelected(processor *nifiapi.ProcessorEntity) {
	loadingPage := pgs.app.enterLoadingScreen(*processor.Component.Name, "Loading processor details")

	go func(ctx context.Context) {
		processorDetails, _, err := pgs.app.apiClient.ProcessorsAPI.GetProcessor(ctx, *processor.Id).Execute()
		pgs.app.app.QueueUpdateDraw(func() {
			frontPage, _ := pgs.app.pages.GetFrontPage()
			if frontPage != "Loading" {
				return
			}
			if err != nil {
				loadingPage.SetLabel("Failed to load processor details")
				return
			}
			pgs.app.pages.RemovePage("Loading")
			pgs.enterProcessorDetailsScreen(processorDetails)
		})
	}(context.Background())
}

func (pgs *processGroupScreen) enterProcessorDetailsScreen(processor *nifiapi.ProcessorEntity) {
	pgs.app.pages.AddAndSwitchToPage(*processor.Id, newProcessorDetailsScreen(pgs.app, processor).pages, true)
}
