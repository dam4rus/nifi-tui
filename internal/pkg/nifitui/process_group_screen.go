package nifitui

import (
	"context"
	"dam4rus/nifi-tui/internal/pkg/nifiapi"
	"fmt"
	"net/http"
	"slices"
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
	SymbolStarted            = "▸"
	SymbolStopped            = "■"
)

func componentStateSymbol(state string) string {
	switch state {
	case "RUNNING":
		return SymbolStarted
	case "STOPPED":
		return SymbolStopped
	default:
		return state
	}
}

type processGroupScreen struct {
	app             *App
	processGroupId  string
	list            *tview.List
	helpFlex        *tview.Flex
	mode            mode
	components      processGroupComponents
	sourceComponent connectionSource
}

func newProcessGroupScreen(app *App, processGroupId string) *processGroupScreen {
	instance := processGroupScreen{
		app:            app,
		processGroupId: processGroupId,
		list:           tview.NewList(),
		helpFlex:       tview.NewFlex(),
	}

	instance.setMode(ModeNone)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(instance.list, 0, 1, true).
		AddItem(instance.helpFlex, 1, 0, false)

	app.pages.AddPage(processGroupId, flex, true, false)

	instance.list.SetTitle(processGroupId).
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true).
		SetInputCapture(instance.handleListInput)
	return &instance
}

func (pgs *processGroupScreen) enter() {
	pgs.loadComponents()
}

func (pgs *processGroupScreen) loadComponents() {
	loadingPage := pgs.app.enterLoadingScreen(pgs.processGroupId, "Loading process group")

	ctx, cancelFunc := context.WithCancel(context.Background())
	pgs.app.cancelFunc = cancelFunc
	task := newProcessGroupService(pgs.app.apiClient, pgs.processGroupId).getComponentsAsync(ctx)

	go func() {
		components, err := task.await()
		if err != nil {
			pgs.app.app.QueueUpdateDraw(func() {
				frontPage, _ := pgs.app.pages.GetFrontPage()
				if frontPage != "Loading" {
					return
				}
				loadingPage.SetLabel(fmt.Sprint("Failed to load process group details", err))
			})
		}
		pgs.components = *components
		pgs.app.app.QueueUpdateDraw(func() {
			pgs.buildAndShowComponents()
		})
	}()
}

func (pgs *processGroupScreen) addProcessGroupsToList() {
	for _, pg := range pgs.components.processGroups {
		pg := pg
		pgs.list.AddItem(pg.GetDisplayName(), pg.id, 0, func() {
			switch pgs.mode {
			case ModeNone:
				childProcessGroupScreen := newProcessGroupScreen(pgs.app, pg.id)
				childProcessGroupScreen.enter()
			case ModeConnect:
				pgs.buildAndShowConnectForm(pgs.sourceComponent, pg)
				pgs.setMode(ModeNone)
			}
		})
	}
}

func (pgs *processGroupScreen) addProcessorsToList() {
	for _, p := range pgs.components.processors {
		p := p
		pgs.list.AddItem(p.GetDisplayName(), p.id, 0, func() {
			switch pgs.mode {
			case ModeNone:
				pgs.onProcessorSelected(p)
			case ModeConnect:
				pgs.buildAndShowConnectForm(pgs.sourceComponent, p)
				pgs.setMode(ModeNone)
			}
		})
	}
}

func (pgs *processGroupScreen) addRemoteProcessGroupsToList() {
	for _, rpg := range pgs.components.remoteProcessGroups {
		rpg := rpg
		pgs.list.AddItem(rpg.GetDisplayName(), rpg.id, 0, func() {
			switch pgs.mode {
			case ModeConnect:
				pgs.buildAndShowConnectForm(pgs.sourceComponent, rpg)
				pgs.setMode(ModeNone)
			}
		})
	}
}

func (pgs *processGroupScreen) addFunnelsToList() {
	for _, f := range pgs.components.funnels {
		f := f
		pgs.list.AddItem(f.GetDisplayName(), f.id, 0, func() {
			switch pgs.mode {
			case ModeConnect:
				pgs.buildAndShowConnectForm(pgs.sourceComponent, f)
				pgs.setMode(ModeNone)
			}
		})
	}
}

func (pgs *processGroupScreen) addInputPortsToList() {
	for _, ip := range pgs.components.inputPorts {
		ip := ip
		pgs.list.AddItem(ip.GetDisplayName(), ip.id, 0, func() {
			switch pgs.mode {
			case ModeConnect:
				pgs.buildAndShowConnectForm(pgs.sourceComponent, ip)
				pgs.setMode(ModeNone)
			}
		})
	}
}

func (pgs *processGroupScreen) addOutputPortsToList() {
	for _, op := range pgs.components.outputPorts {
		op := op
		pgs.list.AddItem(op.GetDisplayName(), op.id, 0, func() {
			switch pgs.mode {
			case ModeConnect:
				pgs.buildAndShowConnectForm(pgs.sourceComponent, op)
				pgs.setMode(ModeNone)
			}
		})
	}
}

func (pgs *processGroupScreen) addConnectionsToList() {
	for _, connection := range pgs.components.connections {
		connection := connection
		pgs.list.AddItem(connection.GetDisplayName(), connection.id, 0, nil)
	}
}

func (pgs *processGroupScreen) buildAndShowComponents() {
	frontPage, _ := pgs.app.pages.GetFrontPage()
	if frontPage != "Loading" {
		return
	}
	pgs.list.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		pgs.updateHelp()
	})
	pgs.addProcessGroupsToList()
	pgs.addProcessorsToList()
	pgs.addRemoteProcessGroupsToList()
	pgs.addFunnelsToList()
	pgs.addInputPortsToList()
	pgs.addOutputPortsToList()
	pgs.addConnectionsToList()

	pgs.app.pages.RemovePage("Loading")
	pgs.app.pages.SwitchToPage(pgs.processGroupId)
}

func (pgs *processGroupScreen) onProcessorSelected(processor *processorEntity) {
	loadingPage := pgs.app.enterLoadingScreen(processor.name, "Loading processor details")

	go func(ctx context.Context) {
		processorDetails, _, err := pgs.app.apiClient.ProcessorsAPI.GetProcessor(ctx, processor.id).Execute()
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
	pgs.app.pages.AddAndSwitchToPage(processor.Component.GetId(), newProcessorDetailsScreen(pgs.app, processor).pages, true)
}

func (pgs *processGroupScreen) buildAndShowAddProcessGroupForm() {
	const PAGE_NAME = "CreateProcessGroup"
	var newProcessGroupName string
	form := pgs.buildAddComponentForm("Create process group", PAGE_NAME, func() {
		pgs.createProcessGroup(newProcessGroupName)
	}).AddInputField("Process group name", "", 0, nil, func(text string) {
		newProcessGroupName = text
	})
	pgs.app.pages.AddPage(PAGE_NAME, buildCenteredGrid(form), true, true)
}

func (pgs *processGroupScreen) buildAndShowAddProcessorForm() {
	const PAGE_NAME = "CreateProcessor"
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

	form := pgs.buildAddComponentForm("Create processor", PAGE_NAME, func() {
		pgs.createProcessor(newProcessorType, newProcessorName)
	}).
		AddFormItem(processorTypeInputField).
		AddInputField("Processor name", "", 0, nil, func(text string) {
			newProcessorName = text
		})
	pgs.app.pages.AddPage(PAGE_NAME, buildCenteredGrid(form), true, true)
}

func (pgs *processGroupScreen) buildAndShowAddInputPortForm() {
	const PAGE_NAME = "CreateInputPort"
	var newInputPortName string
	form := pgs.buildAddComponentForm("Create input port", PAGE_NAME, func() {
		pgs.createInputPort(newInputPortName)
	}).
		AddInputField("Input port name", "", 0, nil, func(text string) {
			newInputPortName = text
		})
	pgs.app.pages.AddPage(PAGE_NAME, buildCenteredGrid(form), true, true)
}

func (pgs *processGroupScreen) buildAndShowAddOutputPortForm() {
	const PAGE_NAME = "CreateOutputPort"
	var newOutputPortName string
	form := pgs.buildAddComponentForm("Create output port", PAGE_NAME, func() {
		pgs.createOutputPort(newOutputPortName)
	}).
		AddInputField("Output port name", "", 0, nil, func(text string) {
			newOutputPortName = text
		})
	pgs.app.pages.AddPage(PAGE_NAME, buildCenteredGrid(form), true, true)
}

func (pgs *processGroupScreen) buildAndShowAddRemoteProcessorGroupForm() {
	const PAGE_NAME = "CreateRemoteProcessGroup"
	var newRemoteProcessGroupName string
	var remoteUrl string
	form := pgs.buildAddComponentForm("Create remote process group", PAGE_NAME, func() {
		pgs.createRemoteProcessGroup(newRemoteProcessGroupName, remoteUrl)
	}).
		AddInputField("Remote process group name", "", 0, nil, func(text string) {
			newRemoteProcessGroupName = text
		}).
		AddInputField("URL", "", 0, nil, func(text string) {
			remoteUrl = text
		})
	pgs.app.pages.AddPage(PAGE_NAME, buildCenteredGrid(form), true, true)
}

func (pgs *processGroupScreen) buildAndShowConnectForm(source connectionSource, destination connectionDestination) (*http.Response, error) {
	const PAGE_NAME = "CreateConnection"
	var outputPorts []outputPortEntity
	selectedOutputPortIndex := -1
	var inputPorts []inputPortEntity
	selectedInputPortIndex := -1
	var selectedRelationships []string
	form := pgs.buildAddComponentForm("Create connection", PAGE_NAME, func() {
		if selectedOutputPortIndex >= 0 {
			source = &outputPorts[selectedOutputPortIndex]
		}
		if selectedInputPortIndex >= 0 {
			destination = &inputPorts[selectedInputPortIndex]
		}
		pgs.createConnection(source, destination, selectedRelationships)
	})
	if sourceProcessor := pgs.components.processors[source.GetId()]; sourceProcessor != nil {
		for _, relationship := range sourceProcessor.getRelationshipNames() {
			form.AddCheckbox(relationship, false, func(checked bool) {
				if checked {
					selectedRelationships = append(selectedRelationships, relationship)
				} else {
					if index := slices.Index(selectedRelationships, relationship); index >= 0 {
						selectedRelationships = slices.Delete(selectedRelationships, index, index+1)
					}
				}
			})
		}
	}

	if inputPortGetter, ok := source.CreateService(pgs.app.apiClient).(portGetter); ok {
		entities, response, err := inputPortGetter.GetInputPorts(context.Background())
		if err != nil {
			return response, err
		}
		inputPorts = entities
		if len(inputPorts) > 0 {
			var inputPortNames []string
			var areRemotePorts bool
			for _, op := range inputPorts {
				inputPortNames = append(inputPortNames, op.name)
				areRemotePorts = op.isRemote
			}
			var label string
			if areRemotePorts {
				label = "Remote Input Ports"
			} else {
				label = "Input Ports"
			}
			form.AddDropDown(label, inputPortNames, 0, func(option string, optionIndex int) {
				selectedInputPortIndex = optionIndex
			})
		}
	}

	if outputPortGetter, ok := destination.CreateService(pgs.app.apiClient).(portGetter); ok {
		entities, response, err := outputPortGetter.GetOutputPorts(context.Background())
		if err != nil {
			return response, err
		}
		outputPorts = entities
		if len(outputPorts) > 0 {
			var outputPortNames []string
			var areRemotePorts bool
			for _, op := range outputPorts {
				outputPortNames = append(outputPortNames, op.name)
				areRemotePorts = op.isRemote
			}
			var label string
			if areRemotePorts {
				label = "Remote Output Ports"
			} else {
				label = "Output Ports"
			}
			form.AddDropDown(label, outputPortNames, 0, func(option string, optionIndex int) {
				selectedOutputPortIndex = optionIndex
			})
		}
	}

	pgs.app.pages.AddPage(PAGE_NAME, buildCenteredGrid(form), true, true)
	return nil, nil
}

func (pgs *processGroupScreen) createProcessGroup(processGroupName string) {
	_, response, err := newProcessGroupService(pgs.app.apiClient, pgs.processGroupId).createProcessGroup(processGroupName)
	if err != nil {
		pgs.app.showErrorDialog(response, err)
		return
	}
	pgs.list.Clear()
	pgs.loadComponents()
}

func (pgs *processGroupScreen) createProcessor(processorType, processorName string) {
	_, response, err := newProcessGroupService(pgs.app.apiClient, pgs.processGroupId).createProcessor(processorType, processorName)
	if err != nil {
		pgs.app.showErrorDialog(response, err)
		return
	}
	pgs.list.Clear()
	pgs.loadComponents()
}

func (pgs *processGroupScreen) createFunnel() {
	_, response, err := newProcessGroupService(pgs.app.apiClient, pgs.processGroupId).createFunnel()
	if err != nil {
		pgs.app.showErrorDialog(response, err)
		return
	}
	pgs.list.Clear()
	pgs.loadComponents()
}

func (pgs *processGroupScreen) createInputPort(portName string) {
	_, response, err := newProcessGroupService(pgs.app.apiClient, pgs.processGroupId).createInputPort(portName)
	if err != nil {
		pgs.app.showErrorDialog(response, err)
		return
	}
	pgs.list.Clear()
	pgs.loadComponents()
}

func (pgs *processGroupScreen) createOutputPort(portName string) {
	_, response, err := newProcessGroupService(pgs.app.apiClient, pgs.processGroupId).createOutputPort(portName)
	if err != nil {
		pgs.app.showErrorDialog(response, err)
		return
	}
	pgs.list.Clear()
	pgs.loadComponents()
}

func (pgs *processGroupScreen) createRemoteProcessGroup(name string, url string) {
	_, response, err := newProcessGroupService(pgs.app.apiClient, pgs.processGroupId).
		createRemoteProcessGroup(name, url)
	if err != nil {
		pgs.app.showErrorDialog(response, err)
		return
	}
	pgs.list.Clear()
	pgs.loadComponents()
}

func (pgs *processGroupScreen) createConnection(source connectionSource, destination connectionDestination, relationships []string) {
	_, response, err := newProcessGroupService(pgs.app.apiClient, pgs.processGroupId).
		createConnection(source, destination, relationships)
	if err != nil {
		pgs.app.showErrorDialog(response, err)
		return
	}
	pgs.list.Clear()
	pgs.loadComponents()
}

func (pgs *processGroupScreen) deleteComponent(c component) {
	response, err := c.CreateService(pgs.app.apiClient).Remove()
	if err != nil {
		pgs.app.showErrorDialog(response, err)
		return
	}
	pgs.list.Clear()
	pgs.loadComponents()
}

func (pgs *processGroupScreen) startComponent(r runnableComponent) {
	service := r.CreateStateChangerService(pgs.app.apiClient)
	response, err := service.Start()
	if err != nil {
		pgs.app.showErrorDialog(response, err)
		return
	}
	pgs.list.Clear()
	pgs.loadComponents()
}

func (pgs *processGroupScreen) stopComponent(r runnableComponent) {
	service := r.CreateStateChangerService(pgs.app.apiClient)
	response, err := service.Stop()
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

func (pgs *processGroupScreen) handleListInput(event *tcell.EventKey) *tcell.EventKey {
	switch pgs.mode {
	case ModeNone:
		return pgs.handleNormalModeInput(event)
	case ModeAdd:
		return pgs.handleAddModeInput(event)
	case ModeConnect:
		return pgs.handleConnectModeInput(event)
	}

	return event
}

func (pgs *processGroupScreen) handleNormalModeInput(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyESC {
		if pgs.processGroupId != "root" {
			pgs.app.pages.RemovePage(pgs.processGroupId)
			return nil
		}
		return event
	}
	switch event.Rune() {
	case 'r':
		pgs.list.Clear()
		pgs.loadComponents()
		return nil
	case 'a':
		pgs.setMode(ModeAdd)
		return nil
	case 'c':
		if selectedComponent := pgs.getSelectedComponent(); selectedComponent != nil {
			if connectable, ok := selectedComponent.(connectionSource); ok {
				pgs.setMode(ModeConnect)
				pgs.sourceComponent = connectable
			}
		}
		return nil
	case 'd':
		if selectedComponent := pgs.getSelectedComponent(); selectedComponent != nil {
			pgs.deleteComponent(selectedComponent)
		}
		return nil
	case 's':
		if selectedComponent := pgs.getSelectedComponent(); selectedComponent != nil {
			if runnable, ok := selectedComponent.(runnableComponent); ok {
				switch runnable.GetState() {
				case "STOPPED":
					pgs.startComponent(runnable)
				case "RUNNING":
					pgs.stopComponent(runnable)
				}
			}
		}
		return nil
	}
	return event
}

func (pgs *processGroupScreen) handleAddModeInput(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	case 'g':
		pgs.buildAndShowAddProcessGroupForm()
	case 'p':
		pgs.buildAndShowAddProcessorForm()
	case 'f':
		pgs.createFunnel()
	case 'i':
		pgs.buildAndShowAddInputPortForm()
	case 'o':
		pgs.buildAndShowAddOutputPortForm()
	case 'r':
		pgs.buildAndShowAddRemoteProcessorGroupForm()
	}
	pgs.setMode(ModeNone)
	return nil
}

func (pgs *processGroupScreen) handleConnectModeInput(event *tcell.EventKey) *tcell.EventKey {
	if event.Key() == tcell.KeyESC {
		pgs.setMode(ModeNone)
		return nil
	}
	return event
}

func (pgs *processGroupScreen) updateHelp() {
	helpFlex := (*helpFlex)(pgs.helpFlex)
	switch pgs.mode {
	case ModeNone:
		helpFlex.clear()
		if pgs.processGroupId == "root" {
			helpFlex.addHelpText("[q]Quit")
		}

		helpFlex.addHelpText("[r]Refresh").
			addHelpText("[a]Add component").
			addHelpText("[d]Delete")

		if selectedComponent := pgs.getSelectedComponent(); selectedComponent != nil {
			switch component := selectedComponent.(type) {
			case *processGroupEntity:
				helpFlex.addHelpText("[Enter]Enter process group").
					addHelpText("[c]Connect")
			case *processorEntity:
				helpFlex.addHelpText("[Enter]Processor details").
					addHelpText("[c]Connect")
				switch component.state {
				case "STOPPED":
					helpFlex.addHelpText("[s]Start")
				case "RUNNING":
					helpFlex.addHelpText("[s]Stop")
				}
			case *funnelEntity:
				helpFlex.addHelpText("[Enter]Funnel details").
					addHelpText("[c]Connect")
			case *inputPortEntity:
				helpFlex.addHelpText("[Enter]Input port details").
					addHelpText("[c]Connect")
			case *outputPortEntity:
				helpFlex.addHelpText("[Enter]Output port details").
					addHelpText("[c]Connect")
			case *remoteProcessGroupEntity:
				helpFlex.addHelpText("[Enter]Remote process group details").
					addHelpText("[c]Connect")
			}
		}
	case ModeAdd:
		helpFlex.clear().
			addHelpText("Add:").
			addHelpText("[g]Process Group").
			addHelpText("[p]Processor").
			addHelpText("[f]Funnel").
			addHelpText("[i]Input port").
			addHelpText("[o]Output port").
			addHelpText("[r]Remote process group")
	case ModeConnect:
		helpFlex.clear().
			addHelpText("Select connection destination")
	}
}

func (pgs *processGroupScreen) getSelectedComponent() component {
	if pgs.list.GetItemCount() == 0 {
		return nil
	}
	if selectedItemIndex := pgs.list.GetCurrentItem(); selectedItemIndex >= 0 {
		_, selectedItemId := pgs.list.GetItemText(selectedItemIndex)
		return pgs.components.findComponent(selectedItemId)
	}
	return nil
}

// func (pgs *processGroupScreen) getSelectedConnectable() connectionSource {
// 	if pgs.list.GetItemCount() == 0 {
// 		return nil
// 	}
// 	if selectedItemIndex := pgs.list.GetCurrentItem(); selectedItemIndex >= 0 {
// 		_, selectedItemId := pgs.list.GetItemText(selectedItemIndex)
// 		return pgs.components.findConnectable(selectedItemId)
// 	}
// 	return nil
// }

// func (pgs *processGroupScreen) getSelectedRunnable() runnable {
// 	if pgs.list.GetItemCount() == 0 {
// 		return nil
// 	}
// 	if selectedItemIndex := pgs.list.GetCurrentItem(); selectedItemIndex >= 0 {
// 		_, selectedItemId := pgs.list.GetItemText(selectedItemIndex)
// 		return pgs.components.findComponent(selectedItemId)
// 	}
// 	return nil
// }

func (pgs *processGroupScreen) buildAddComponentForm(title string, pageName string, createFunc func()) *tview.Form {
	form := tview.NewForm().
		AddButton("Create", func() {
			createFunc()
			pgs.app.pages.RemovePage(pageName)
		}).
		AddButton("Cancel", func() {
			pgs.app.pages.RemovePage(pageName)
		})
	form.SetTitle(title).
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyEsc:
				pgs.app.pages.RemovePage(pageName)
				return nil
			}
			return event
		})
	return form
}

func buildCenteredGrid(form *tview.Form) *tview.Grid {
	return tview.NewGrid().
		SetColumns(0, 100, 0).
		SetRows(0, 5+form.GetFormItemCount()*2, 0).
		AddItem(form, 1, 1, 1, 1, 0, 0, true)
}
