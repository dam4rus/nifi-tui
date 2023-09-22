package nifitui

import (
	"context"
	"dam4rus/nifi-tui/internal/pkg/nifiapi"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type mode int

const (
	ModeNone mode = iota
	ModeAdd
	ModeConnect
)

const (
	SymbolProcessGroup       = "📁"
	SymbolProcessor          = "⊞"
	SymbolRemoteProcessGroup = "☁️"
	SymbolFunnel             = "Ⴤ"
	SymbolInputPort          = "⇥"
	SymbolOutputPort         = "↦"
	SymbolConnection         = "->"
)

type processGroupScreen struct {
	app             *App
	processGroupId  string
	list            *tview.List
	helpFlex        *tview.Flex
	mode            mode
	components      processGroupComponents
	sourceComponent connectableComponent
}

func newProcessGroupScreen(app *App, processGroupId string) *processGroupScreen {
	instance := processGroupScreen{
		app:            app,
		processGroupId: processGroupId,
		list:           tview.NewList(),
		helpFlex:       tview.NewFlex(),
	}

	instance.helpFlex.SetDirection(tview.FlexColumn)
	instance.setMode(ModeNone)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(instance.list, 0, 1, true).
		AddItem(instance.helpFlex, 1, 0, false)

	app.pages.AddPage(processGroupId, flex, true, false)

	instance.list.SetTitle(processGroupId).
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch instance.mode {
			case ModeNone:
				if event.Key() == tcell.KeyESC {
					if processGroupId != "root" {
						app.pages.RemovePage(processGroupId)
						return nil
					}
				}
				switch event.Rune() {
				case 'r':
					instance.list.Clear()
					instance.loadComponents()
					return nil
				case 'a':
					instance.setMode(ModeAdd)
					return nil
				case 'c':
					if selectedComponent := instance.getSelectedComponent(); selectedComponent != nil {
						if selectedComponent.GetComponentType() == ComponentTypeProcessor {
							instance.setMode(ModeConnect)
							instance.sourceComponent = selectedComponent
						}
					}
					return nil
				case 'd':
					if selectedComponent := instance.getSelectedComponent(); selectedComponent != nil {
						instance.deleteComponent(selectedComponent)
					}

					return nil
				}
			case ModeAdd:
				switch event.Rune() {
				case 'g':
					instance.buildAndShowAddProcessGroupForm()
				case 'p':
					instance.buildAndShowAddProcessorForm()
				}
				instance.setMode(ModeNone)
				return nil
			case ModeConnect:
				if event.Key() == tcell.KeyESC {
					instance.setMode(ModeNone)
				}
			}

			return event
		})
	return &instance
}

func (pgs *processGroupScreen) enter() {
	pgs.loadComponents()
}

func (pgs *processGroupScreen) loadComponents() {
	loadingPage := pgs.app.enterLoadingScreen(pgs.processGroupId, "Loading process group")

	proxy := newAsyncProcessGroupProxy(pgs.app.apiClient, pgs.processGroupId)
	pgs.app.cancelFunc = proxy.cancelFunc
	proxy.waitGroup.Add(7)
	processGroupsChannel := proxy.getProcessGroups()
	processorsChannel := proxy.getProcessors()
	connectionsChannel := proxy.getConnections()
	funnelsChannel := proxy.getFunnels()
	inputPortsChannel := proxy.getInputPorts()
	outputPortsChannel := proxy.getOutputPorts()
	remoteProcessGroupsChannel := proxy.getRemoteProcessGroups()

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
		remoteProcessGroups := <-remoteProcessGroupsChannel
		pgs.components.processGroups = processGroups.ProcessGroups
		pgs.components.processors = processors.Processors
		pgs.components.connections = connections.Connections
		pgs.components.funnels = funnels.Funnels
		pgs.components.inputPorts = inputPorts.InputPorts
		pgs.components.outputPorts = outputPorts.OutputPorts
		pgs.components.remoteProcessGroups = remoteProcessGroups.RemoteProcessGroups
		pgs.app.app.QueueUpdateDraw(func() {
			pgs.buildAndShowComponents()
		})
	}(proxy.ctx)
}

func (pgs *processGroupScreen) buildAndShowComponents() {
	frontPage, _ := pgs.app.pages.GetFrontPage()
	if frontPage != "Loading" {
		return
	}
	pgs.list.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		pgs.updateHelp()
	})
	for _, processGroup := range pgs.components.processGroups {
		processGroup := processGroup
		pgs.list.AddItem(SymbolProcessGroup+processGroup.Component.GetName(), processGroup.Component.GetId(), 0, func() {
			childProcessGroupScreen := newProcessGroupScreen(pgs.app, *processGroup.Id)
			childProcessGroupScreen.enter()
		})
	}
	for _, processor := range pgs.components.processors {
		processor := processor
		pgs.list.AddItem(SymbolProcessor+processor.Component.GetName(), processor.Component.GetId(), 0, func() {
			switch pgs.mode {
			case ModeNone:
				pgs.onProcessorSelected(&processor)
			case ModeConnect:
				if pgs.sourceComponent.GetComponentType() == ComponentTypeProcessor {
					pgs.buildAndShowConnectProcessorToProcessor(pgs.sourceComponent.GetId(), processor.Component.GetId())
				}
				pgs.setMode(ModeNone)
			}
		})
	}
	for _, remoteProcessGroup := range pgs.components.remoteProcessGroups {
		remoteProcessGroup := remoteProcessGroup
		pgs.list.AddItem(SymbolRemoteProcessGroup+remoteProcessGroup.Component.GetName(), remoteProcessGroup.Component.GetId(), 0, nil)
	}
	for _, funnel := range pgs.components.funnels {
		funnel := funnel
		pgs.list.AddItem(SymbolFunnel, funnel.Component.GetId(), 0, nil)
	}
	for _, inputPort := range pgs.components.inputPorts {
		inputPort := inputPort
		pgs.list.AddItem(SymbolInputPort+inputPort.Component.GetName(), inputPort.Component.GetId(), 0, nil)
	}
	for _, outputPort := range pgs.components.outputPorts {
		outputPort := outputPort
		pgs.list.AddItem(SymbolOutputPort+outputPort.Component.GetName(), outputPort.Component.GetId(), 0, nil)
	}

	for _, connection := range pgs.components.connections {
		connection := connection
		sourceComponent := pgs.components.findComponent(connection.GetSourceType(), connection.GetSourceId())
		destComponent := pgs.components.findComponent(connection.GetDestinationType(), connection.GetDestinationId())
		if sourceComponent == nil || destComponent == nil {
			continue
		}
		primaryText := SymbolConnection + sourceComponent.GetName() + "[" + sourceComponent.GetId() + "]" +
			SymbolConnection + destComponent.GetName() + "[" + destComponent.GetId() + "]"
		pgs.list.AddItem(primaryText, connection.Component.GetId(), 0, nil)
	}
	pgs.app.pages.RemovePage("Loading")
	pgs.app.pages.SwitchToPage(pgs.processGroupId)
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

func (pgs *processGroupScreen) buildAndShowAddProcessGroupForm() {
	var newProcessGroupName string
	form := tview.NewForm().
		AddInputField("Process group name", "", 0, nil, func(text string) {
			newProcessGroupName = text
		}).
		AddButton("Create", func() {
			pgs.createProcessGroup(newProcessGroupName)
			pgs.app.pages.RemovePage("CreateProcessGroup")
		}).
		AddButton("Cancel", func() {
			pgs.app.pages.RemovePage("CreateProcessGroup")
		})
	form.SetTitle("Create process group").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyEsc:
				pgs.app.pages.RemovePage("CreateProcessGroup")
				return nil
			case tcell.KeyCtrlS:
				pgs.createProcessGroup(newProcessGroupName)
				pgs.app.pages.RemovePage("CreateProcessGroup")
				return nil
			}
			return event
		})

	grid := tview.NewGrid().
		SetColumns(0, 100, 0).
		SetRows(0, 7, 0).
		AddItem(form, 1, 1, 1, 1, 0, 0, true)
	pgs.app.pages.AddPage("CreateProcessGroup", grid, true, true)
}

func (pgs *processGroupScreen) buildAndShowAddProcessorForm() {
	var newProcessorName string
	var newProcessorType string

	processorTypeInputField := tview.NewInputField()
	processorTypeInputField.SetLabel("Processor type").
		SetAutocompleteFunc(func(currentText string) (entries []string) {
			var processorTypes []string
			for _, processorType := range pgs.app.processorTypes {
				if strings.Contains(processorType.GetType(), currentText) {
					processorTypes = append(processorTypes, processorType.GetType())
				}
			}
			return processorTypes
		}).
		SetAutocompletedFunc(func(text string, index, source int) bool {
			processorTypeInputField.SetText(text)
			return true
		}).
		SetChangedFunc(func(text string) {
			newProcessorType = text
		})

	form := tview.NewForm().
		AddFormItem(processorTypeInputField).
		AddInputField("Processor name", "", 0, nil, func(text string) {
			newProcessorName = text
		}).
		AddButton("Create", func() {
			pgs.createProcessor(newProcessorType, newProcessorName)
			pgs.app.pages.RemovePage("CreateProcessor")
		}).
		AddButton("Cancel", func() {
			pgs.app.pages.RemovePage("CreateProcessor")
		})
	form.SetTitle("Create process group").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyEsc:
				pgs.app.pages.RemovePage("CreateProcessor")
				return nil
			case tcell.KeyCtrlS:
				pgs.createProcessor(newProcessorType, newProcessorName)
				pgs.app.pages.RemovePage("CreateProcessor")
				return nil
			}
			return event
		})

	grid := tview.NewGrid().
		SetColumns(0, 100, 0).
		SetRows(0, 9, 0).
		AddItem(form, 1, 1, 1, 1, 0, 0, true)
	pgs.app.pages.AddPage("CreateProcessor", grid, true, true)
}

func (pgs *processGroupScreen) buildAndShowConnectProcessorToProcessor(sourceId, destinationId string) {
	var sourceProcessor *nifiapi.ProcessorEntity
	for _, processor := range pgs.components.processors {
		if processor.GetId() == sourceId {
			sourceProcessor = &processor
			break
		}
	}
	var selectedRelationships []string
	form := tview.NewForm()
	for _, relationship := range sourceProcessor.Component.GetRelationships() {
		form.AddCheckbox(relationship.GetName(), false, func(checked bool) {
			if checked {
				selectedRelationships = append(selectedRelationships, relationship.GetName())
			} else {
				if index := slices.Index(selectedRelationships, relationship.GetName()); index >= 0 {
					selectedRelationships = slices.Delete(selectedRelationships, index, index+1)
				}
			}
		})
	}
	form.AddButton("Create", func() {
		pgs.createProcessorToProcessorConnection(sourceId, destinationId, selectedRelationships)
		pgs.app.pages.RemovePage("CreateConnection")
	}).
		AddButton("Cancel", func() {
			pgs.app.pages.RemovePage("CreateConnection")
		})
	form.SetTitle("Create connection").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyEsc:
				pgs.app.pages.RemovePage("CreateConnection")
				return nil
			case tcell.KeyCtrlS:
				pgs.createProcessorToProcessorConnection(sourceId, destinationId, selectedRelationships)
				pgs.app.pages.RemovePage("CreateConnection")
				return nil
			}
			return event
		})

	grid := tview.NewGrid().
		SetColumns(0, 100, 0).
		SetRows(0, 5+(len(sourceProcessor.Component.GetRelationships())*2), 0).
		AddItem(form, 1, 1, 1, 1, 0, 0, true)
	pgs.app.pages.AddPage("CreateConnection", grid, true, true)
}

func (pgs *processGroupScreen) createProcessGroup(processGroupName string) {
	version := int64(0)
	processGroupEntity := nifiapi.ProcessGroupEntity{
		Component: &nifiapi.ProcessGroupDTO{
			Name: &processGroupName,
		},
		Revision: &nifiapi.RevisionDTO{
			Version: &version,
		},
	}
	_, response, err := pgs.app.apiClient.ProcessGroupsAPI.CreateProcessGroup(context.Background(), pgs.processGroupId).
		Body(processGroupEntity).
		Execute()
	if err != nil {
		pgs.app.showErrorDialog(response, err)
		return
	}
	pgs.list.Clear()
	pgs.loadComponents()
}

func (pgs *processGroupScreen) createProcessor(processorType, processorName string) {
	version := int64(0)
	processorEntity := nifiapi.ProcessorEntity{
		Component: &nifiapi.ProcessorDTO{
			Name: &processorName,
			Type: &processorType,
		},
		Revision: &nifiapi.RevisionDTO{
			Version: &version,
		},
	}
	_, response, err := pgs.app.apiClient.ProcessGroupsAPI.CreateProcessor(context.Background(), pgs.processGroupId).
		Body(processorEntity).
		Execute()
	if err != nil {
		pgs.app.showErrorDialog(response, err)
		return
	}
	pgs.list.Clear()
	pgs.loadComponents()
}

func (pgs *processGroupScreen) createProcessorToProcessorConnection(sourceId, destinationId string, relationships []string) {
	version := int64(0)
	connectionEntity := nifiapi.ConnectionEntity{
		Component: &nifiapi.ConnectionDTO{
			SelectedRelationships: relationships,
			Source: &nifiapi.ConnectableDTO{
				Id:      sourceId,
				GroupId: pgs.processGroupId,
				Type:    "PROCESSOR",
			},
			Destination: &nifiapi.ConnectableDTO{
				Id:      destinationId,
				GroupId: pgs.processGroupId,
				Type:    "PROCESSOR",
			},
		},
		Revision: &nifiapi.RevisionDTO{
			Version: &version,
		},
	}
	_, response, err := pgs.app.apiClient.ProcessGroupsAPI.CreateConnection(context.Background(), pgs.processGroupId).
		Body(connectionEntity).
		Execute()
	if err != nil {
		pgs.app.showErrorDialog(response, err)
		return
	}
	pgs.list.Clear()
	pgs.loadComponents()
}

func (pgs *processGroupScreen) deleteComponent(component connectableComponent) {
	response, err := component.IntoProxy(pgs.app.apiClient).Remove()
	if err != nil {
		pgs.app.showErrorDialog(response, err)
		return
	}
	pgs.list.Clear()
	pgs.loadComponents()
}

func (pgs *processGroupScreen) deleteRemoteProcessGroup(remoteProcessGroupId string) {
	remoteProcessGroup, response, err := pgs.app.apiClient.RemoteProcessGroupsAPI.GetRemoteProcessGroup(context.Background(), remoteProcessGroupId).
		Execute()
	if err != nil {
		pgs.app.showErrorDialog(response, err)
		return
	}
	_, response, err = pgs.app.apiClient.RemoteProcessGroupsAPI.RemoveRemoteProcessGroup(context.Background(), remoteProcessGroupId).
		Version(strconv.Itoa(int(remoteProcessGroup.Revision.GetVersion()))).
		Execute()
	if err != nil {
		pgs.app.showErrorDialog(response, err)
		return
	}
	pgs.list.Clear()
	pgs.loadComponents()
}

func (pgs *processGroupScreen) setMode(mode mode) {
	pgs.mode = mode
	pgs.sourceComponent = nil
	pgs.updateHelp()
}

func (pgs *processGroupScreen) updateHelp() {
	switch pgs.mode {
	case ModeNone:
		pgs.helpFlex.Clear()
		if pgs.processGroupId == "root" {
			pgs.helpFlex.AddItem(tview.NewTextView().SetText("[q]Quit"), len("[q]Quit")+1, 0, false)
		}
		pgs.helpFlex.AddItem(tview.NewTextView().SetText("[r]Refresh"), len("[r]Refresh")+1, 0, false).
			AddItem(tview.NewTextView().SetText("[a]Add component"), len("[a]Add component")+1, 1, false).
			AddItem(tview.NewTextView().SetText("[d]Delete"), len("[d]Delete")+1, 0, false)

		if selectedComponent := pgs.getSelectedComponent(); selectedComponent != nil {
			switch selectedComponent.GetComponentType() {
			case ComponentTypeProcessGroup:
				pgs.helpFlex.AddItem(tview.NewTextView().SetText("[Enter]Enter process group"), len("[Enter]Enter process group")+1, 0, false)
			case ComponentTypeProcessor:
				pgs.helpFlex.AddItem(tview.NewTextView().SetText("[c]Connect"), len("[c]Connect")+1, 0, false).
					AddItem(tview.NewTextView().SetText("[Enter]Processor details"), len("[Enter]Processor details")+1, 0, false)
			case ComponentTypeFunnel:
				pgs.helpFlex.AddItem(tview.NewTextView().SetText("[Enter]Funnel details"), len("[Enter]Funnel details")+1, 0, false)
			case ComponentTypeInputPort:
				pgs.helpFlex.AddItem(tview.NewTextView().SetText("[Enter]Input port details"), len("[Enter]Input port details")+1, 0, false)
			case ComponentTypeOutputPort:
				pgs.helpFlex.AddItem(tview.NewTextView().SetText("[Enter]Output port details"), len("[Enter]Output port details")+1, 0, false)
			}
		}
	case ModeAdd:
		pgs.helpFlex.Clear().
			AddItem(tview.NewTextView().SetText("Add:"), 5, 0, false).
			AddItem(tview.NewTextView().SetText("[g]Process Group"), len("[g]Process Group")+1, 0, false).
			AddItem(tview.NewTextView().SetText("[p]Processor"), len("[p]Processor")+1, 0, false)
	case ModeConnect:
		pgs.helpFlex.Clear().
			AddItem(tview.NewTextView().SetText("Select connection destination"), 0, 1, false)
	}
}

func (pgs *processGroupScreen) getSelectedComponent() connectableComponent {
	if pgs.list.GetItemCount() == 0 {
		return nil
	}
	if selectedItemIndex := pgs.list.GetCurrentItem(); selectedItemIndex >= 0 {
		selectedItemText, selectedItemId := pgs.list.GetItemText(selectedItemIndex)
		if strings.HasPrefix(selectedItemText, SymbolProcessGroup) {
			return &processGroup{id: selectedItemId}
		} else if strings.HasPrefix(selectedItemText, SymbolProcessor) {
			return &processor{id: selectedItemId}
		} else if strings.HasPrefix(selectedItemText, SymbolFunnel) {
			return &funnel{id: selectedItemId}
		} else if strings.HasPrefix(selectedItemText, SymbolInputPort) {
			return &inputPort{id: selectedItemId}
		} else if strings.HasPrefix(selectedItemText, SymbolOutputPort) {
			return &outputPort{id: selectedItemId}
		} else if strings.HasPrefix(selectedItemText, SymbolConnection) {
			return &connection{id: selectedItemId}
		} else if strings.HasPrefix(selectedItemText, SymbolRemoteProcessGroup) {
			return &remoteProcessGroup{id: selectedItemId}
		}
		return nil
	}
	return nil
}