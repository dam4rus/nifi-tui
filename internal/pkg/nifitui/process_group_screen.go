package nifitui

import (
	"context"
	"dam4rus/nifi-tui/internal/pkg/nifiapi"
	"fmt"
	"net/http"
	"slices"
	"strings"
	"sync"

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
	sourceComponent connectionSource
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
		SetInputCapture(instance.handleListInput)
	return &instance
}

func (pgs *processGroupScreen) enter() {
	pgs.loadComponents()
}

func (pgs *processGroupScreen) loadComponents() {
	loadingPage := pgs.app.enterLoadingScreen(pgs.processGroupId, "Loading process group")

	ctx, cancelFunc := context.WithCancel(context.Background())
	errorChannel := make(chan error)
	proxy := newProcessGroupService(pgs.app.apiClient, pgs.processGroupId)
	pgs.app.cancelFunc = cancelFunc
	var waitGroup sync.WaitGroup
	waitGroup.Add(7)
	processGroupsChannel := make(chan []processGroupEntity, 1)
	processorsChannel := make(chan []processorEntity, 1)
	connectionsChannel := make(chan []connectionEntity, 1)
	funnelsChannel := make(chan []funnelEntity, 1)
	inputPortsChannel := make(chan []inputPortEntity, 1)
	outputPortsChannel := make(chan []outputPortEntity, 1)
	remoteProcessGroupsChannel := make(chan []remoteProcessGroupEntity, 1)
	go func() {
		processGroups, _, err := proxy.getProcessGroups(ctx)
		if err != nil {
			errorChannel <- err
			return
		}
		processGroupsChannel <- processGroups
	}()
	go func() {
		processors, _, err := proxy.getProcessors(ctx)
		if err != nil {
			errorChannel <- err
			return
		}
		processorsChannel <- processors
	}()
	go func() {
		connections, _, err := proxy.getConnections(ctx)
		if err != nil {
			errorChannel <- err
			return
		}
		connectionsChannel <- connections
	}()
	go func() {
		funnels, _, err := proxy.getFunnels(ctx)
		if err != nil {
			errorChannel <- err
			return
		}
		funnelsChannel <- funnels
	}()
	go func() {
		inputPorts, _, err := proxy.getInputPorts(ctx)
		if err != nil {
			errorChannel <- err
			return
		}
		inputPortsChannel <- inputPorts
	}()
	go func() {
		outputPorts, _, err := proxy.getOutputPorts(ctx)
		if err != nil {
			errorChannel <- err
			return
		}
		outputPortsChannel <- outputPorts
	}()
	go func() {
		remoteProcessGroups, _, err := proxy.getRemoteProcessGroups(ctx)
		if err != nil {
			errorChannel <- err
			return
		}
		remoteProcessGroupsChannel <- remoteProcessGroups
	}()

	go func() {
		waitGroup.Wait()
		select {
		case err := <-errorChannel:
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
		pgs.components = processGroupComponents{}
		for _, processGroup := range processGroups {
			pgs.components.processGroups[processGroup.id] = &processGroup
		}
		for _, processor := range processors {
			pgs.components.processors[processor.id] = &processor
		}
		for _, connection := range connections {
			pgs.components.connections[connection.id] = &connection
		}
		for _, funnel := range funnels {
			pgs.components.funnels[funnel.id] = &funnel
		}
		for _, inputPort := range inputPorts {
			pgs.components.inputPorts[inputPort.id] = &inputPort
		}
		for _, outputPort := range outputPorts {
			pgs.components.outputPorts[outputPort.id] = &outputPort
		}
		for _, remoteProcessGroup := range remoteProcessGroups {
			pgs.components.remoteProcessGroups[remoteProcessGroup.id] = &remoteProcessGroup
		}
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
		sourceComponent := pgs.components.findDisplayableConnectionSource(connection)
		destComponent := pgs.components.findDisplayableConnectionDestination(connection)
		if sourceComponent == nil || destComponent == nil {
			continue
		}
		primaryText := sourceComponent.GetDisplayName() + "[" + sourceComponent.GetId() + "]" +
			destComponent.GetDisplayName() + "[" + destComponent.GetId() + "]"
		pgs.list.AddItem(primaryText, connection.id, 0, nil)
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

func (pgs *processGroupScreen) buildAndShowConnectForm(source connectionSource, destination connectionDestination) (*http.Response, error) {
	formHeight := 5

	form := tview.NewForm()
	var selectedRelationships []string
	if sourceProcessor := pgs.components.processors[source.GetId()]; sourceProcessor != nil {
		for _, relationship := range sourceProcessor.getRelationshipNames() {
			formHeight += 2
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
	var inputPorts []inputPortEntity
	var isLocalInputPort bool
	selectedInputPortIndex := -1
	switch destination.(type) {
	case *processGroupEntity:
		entities, response, err := newProcessGroupService(pgs.app.apiClient, destination.GetId()).
			getInputPorts(context.Background())
		if err != nil {
			return response, err
		}
		inputPorts = entities
		isLocalInputPort = true
	case *remoteProcessGroupEntity:
		entities, response, err := newRemoteProcessGroupService(pgs.app.apiClient, destination.GetId()).
			getInputPorts()
		if err != nil {
			return response, err
		}
		inputPorts = entities
		isLocalInputPort = false
	}

	if len(inputPorts) > 0 {
		formHeight += 2
		var inputPortNames []string
		for _, ip := range inputPorts {
			inputPortNames = append(inputPortNames, ip.name)
		}
		var label string
		if isLocalInputPort {
			label = "Input Ports"
		} else {
			label = "Remote Input Ports"
		}
		form.AddDropDown(label, inputPortNames, 0, func(option string, optionIndex int) {
			selectedInputPortIndex = optionIndex
		})
	}

	form.AddButton("Create", func() {
		if selectedInputPortIndex >= 0 {
			destination = &inputPorts[selectedInputPortIndex]
		}
		pgs.createConnection(source, destination, selectedRelationships)
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
				if selectedInputPortIndex >= 0 {
					destination = &inputPorts[selectedInputPortIndex]
				}
				pgs.createConnection(source, destination, selectedRelationships)
				pgs.app.pages.RemovePage("CreateConnection")
				return nil
			}
			return event
		})

	grid := tview.NewGrid().
		SetColumns(0, 100, 0).
		SetRows(0, formHeight, 0).
		AddItem(form, 1, 1, 1, 1, 0, 0, true)
	pgs.app.pages.AddPage("CreateConnection", grid, true, true)
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
		if selectedConnectable := pgs.getSelectedConnectable(); selectedConnectable != nil {
			pgs.setMode(ModeConnect)
			pgs.sourceComponent = selectedConnectable
		}
		return nil
	case 'd':
		if selectedComponent := pgs.getSelectedComponent(); selectedComponent != nil {
			pgs.deleteComponent(selectedComponent)
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
			switch selectedComponent.(type) {
			case *processGroupEntity:
				pgs.helpFlex.AddItem(tview.NewTextView().SetText("[Enter]Enter process group"), len("[Enter]Enter process group")+1, 0, false)
			case *processorEntity:
				pgs.helpFlex.AddItem(tview.NewTextView().SetText("[c]Connect"), len("[c]Connect")+1, 0, false).
					AddItem(tview.NewTextView().SetText("[Enter]Processor details"), len("[Enter]Processor details")+1, 0, false)
			case *funnelEntity:
				pgs.helpFlex.AddItem(tview.NewTextView().SetText("[Enter]Funnel details"), len("[Enter]Funnel details")+1, 0, false)
			case *inputPortEntity:
				pgs.helpFlex.AddItem(tview.NewTextView().SetText("[Enter]Input port details"), len("[Enter]Input port details")+1, 0, false)
			case *outputPortEntity:
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

func (pgs *processGroupScreen) getSelectedConnectable() connectionSource {
	if pgs.list.GetItemCount() == 0 {
		return nil
	}
	if selectedItemIndex := pgs.list.GetCurrentItem(); selectedItemIndex >= 0 {
		_, selectedItemId := pgs.list.GetItemText(selectedItemIndex)
		return pgs.components.findConnectable(selectedItemId)
	}
	return nil
}
