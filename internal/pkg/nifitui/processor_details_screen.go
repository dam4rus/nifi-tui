package nifitui

import (
	"context"
	"slices"
	"strconv"
	"strings"

	"github.com/dam4rus/nifi-tui/internal/pkg/nifiapi"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type processorDetailsScreen struct {
	app                *App
	processor          *nifiapi.ProcessorEntity
	pages              *tview.Pages
	settingsPanel      *settingsPanel
	schedulingPanel    *schedulingPanel
	propertiesPanel    *propertiesPanel
	relationshipsPanel *relationshipsPanel
	modified           bool
}

type settingsPanel struct {
	flex                    *tview.Flex
	nameTextView            *tview.TextView
	penaltyDurationTextView *tview.TextView
	yieldDurationTextView   *tview.TextView
	bulletinLevelTextView   *tview.TextView
	helpFlex                *tview.Flex
}

type settingsFormResult struct {
	name            string
	penaltyDuration string
	yieldDuration   string
	bulletinLevel   string
}

type schedulingPanel struct {
	flex                       *tview.Flex
	schedulingStrategyTextView *tview.TextView
	concurrentTaskTextView     *tview.TextView
	runScheduleTextView        *tview.TextView
	executionTextView          *tview.TextView
	helpFlex                   *tview.Flex
}

type schedulingFormResult struct {
	schedulingStrategy string
	concurrentTask     int
	runSchedule        string
	execution          string
}

type propertiesPanel struct {
	flex     *tview.Flex
	listView *tview.List
	helpFlex *tview.Flex
}

type relationshipsPanel struct {
	flex     *tview.Flex
	listView *tview.List
	helpFlex *tview.Flex
}

func newProcessorDetailsScreen(app *App, processor *nifiapi.ProcessorEntity) *processorDetailsScreen {
	processorDetailsScreen := &processorDetailsScreen{
		app:       app,
		processor: processor,
		pages:     tview.NewPages(),
	}
	processorDetailsScreen.settingsPanel = processorDetailsScreen.buildSettingsPanel()
	processorDetailsScreen.schedulingPanel = processorDetailsScreen.buildSchedulingPanel()
	processorDetailsScreen.propertiesPanel = processorDetailsScreen.buildPropertiesPanel()
	processorDetailsScreen.relationshipsPanel = processorDetailsScreen.buildRelationshipsPanel()

	processorDetailsScreen.pages.
		AddAndSwitchToPage("Settings", processorDetailsScreen.settingsPanel.flex, true).
		AddPage("Scheduling", processorDetailsScreen.schedulingPanel.flex, true, false).
		AddPage("Properties", processorDetailsScreen.propertiesPanel.flex, true, false).
		AddPage("Relationships", processorDetailsScreen.relationshipsPanel.flex, true, false)

	return processorDetailsScreen
}

func (s *processorDetailsScreen) save() {
	autoTerminatedRelationships, retriedRelationships := s.getRelationshipNames()
	processor := nifiapi.ProcessorEntity{
		Component: &nifiapi.ProcessorDTO{
			Id:   s.processor.Id,
			Name: s.processor.Component.Name,
			Config: &nifiapi.ProcessorConfigDTO{
				PenaltyDuration:                  s.processor.Component.Config.PenaltyDuration,
				YieldDuration:                    s.processor.Component.Config.YieldDuration,
				BulletinLevel:                    s.processor.Component.Config.BulletinLevel,
				Properties:                       s.processor.Component.Config.Properties,
				SchedulingStrategy:               s.processor.Component.Config.SchedulingStrategy,
				ConcurrentlySchedulableTaskCount: s.processor.Component.Config.ConcurrentlySchedulableTaskCount,
				SchedulingPeriod:                 s.processor.Component.Config.SchedulingPeriod,
				ExecutionNode:                    s.processor.Component.Config.ExecutionNode,
				AutoTerminatedRelationships:      autoTerminatedRelationships,
				RetriedRelationships:             retriedRelationships,
			},
		},
		Revision: s.processor.Revision,
	}
	updatedProcessor, response, err := s.app.apiClient.ProcessorsAPI.UpdateProcessor(context.Background(), *processor.Component.Id).
		Body(processor).
		Execute()
	if err != nil {
		s.app.showErrorDialog(response, err)
		return
	}
	s.processor = updatedProcessor
	s.modified = false
}

func (s *processorDetailsScreen) handleInput(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyCtrlS:
		s.save()
		return nil
	case tcell.KeyEsc:
		if s.modified {
			s.showModifiedModal()
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
		s.pages.SwitchToPage("Settings")
		return nil
	case '2':
		s.pages.SwitchToPage("Scheduling")
		return nil
	case '3':
		s.pages.SwitchToPage("Properties")
		return nil
	case '4':
		s.pages.SwitchToPage("Relationships")
		return nil
	}
	return event
}

func (s *processorDetailsScreen) onPropertySelected(key string) {
	propertyDescriptor := (*s.processor.Component.Config.Descriptors)[key]
	var allowableValues AllowableValues
	for _, allowableValue := range propertyDescriptor.AllowableValues {
		allowableValues = append(allowableValues, *allowableValue.AllowableValue.Value)
	}

	if allowableValues.IsBoolean() {
		if *(*s.processor.Component.Config.Properties)[key] == "true" {
			s.applyProperty(key, "false")
		} else {
			s.applyProperty(key, "true")
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
		form.AddDropDown(key, allowableValueDisplayNames, slices.Index(allowableValues, newValue), func(_ string, optionIndex int) {
			if optionIndex >= 0 {
				s.applyProperty(key, allowableValues[optionIndex])
				s.pages.RemovePage("Property")
			}
		})

		gridHeight = 5
	} else {
		form.AddTextArea(key, newValue, 100, 5, 0, func(text string) {
			newValue = text
		}).
			AddButton("Save", func() {
				s.applyProperty(key, newValue)
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
				s.applyProperty(key, newValue)
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

func (s *processorDetailsScreen) onRelationshipSelected(index int) {
	relationship := s.processor.Component.Relationships[index]
	form := tview.NewForm().
		AddCheckbox("Auto-terminated", relationship.GetAutoTerminate(), func(checked bool) {
			relationship.SetAutoTerminate(checked)
		}).
		AddCheckbox("Retry", relationship.GetRetry(), func(checked bool) {
			relationship.SetRetry(checked)
		}).
		AddButton("Save", func() {
			s.applyRelationship(index, relationship)
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
				s.applyRelationship(index, relationship)
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

func (s *processorDetailsScreen) applyProperty(key, newValue string) {
	(*s.processor.Component.Config.Properties)[key] = &newValue
	descriptor := (*s.processor.Component.Config.Descriptors)[key]
	displayName := descriptor.DisplayName
	var value string
	if allowableValue := findAllowableValueByValue(descriptor.AllowableValues, newValue); allowableValue != nil {
		value = *allowableValue.DisplayName
	} else {
		value = newValue
	}
	s.propertiesPanel.listView.SetItemText(s.propertiesPanel.listView.GetCurrentItem(), *displayName, value)
	s.modified = true
}

func (s *processorDetailsScreen) applyRelationship(index int, relationship nifiapi.RelationshipDTO) {
	s.processor.Component.Relationships[index] = relationship
	s.relationshipsPanel.listView.SetItemText(s.relationshipsPanel.listView.GetCurrentItem(),
		*relationship.Name,
		buildRelationshipSecondaryText(relationship))
	s.modified = true
}

func (s *processorDetailsScreen) applySettings(result settingsFormResult) {
	s.processor.Component.Name = &result.name
	s.processor.Component.Config.PenaltyDuration = &result.penaltyDuration
	s.processor.Component.Config.YieldDuration = &result.yieldDuration
	s.processor.Component.Config.BulletinLevel = &result.bulletinLevel
	s.settingsPanel.nameTextView.SetText(result.name)
	s.settingsPanel.penaltyDurationTextView.SetText(result.penaltyDuration)
	s.settingsPanel.yieldDurationTextView.SetText(result.yieldDuration)
	s.settingsPanel.bulletinLevelTextView.SetText(result.bulletinLevel)
	s.modified = true
}

func (s *processorDetailsScreen) applyScheduling(result schedulingFormResult) {
	concurrentTask := int32(result.concurrentTask)
	s.processor.Component.Config.SchedulingStrategy = &result.schedulingStrategy
	s.processor.Component.Config.ConcurrentlySchedulableTaskCount = &concurrentTask
	s.processor.Component.Config.SchedulingPeriod = &result.runSchedule
	s.processor.Component.Config.ExecutionNode = &result.execution
	s.schedulingPanel.schedulingStrategyTextView.SetText(result.schedulingStrategy)
	s.schedulingPanel.concurrentTaskTextView.SetText(strconv.Itoa(result.concurrentTask))
	s.schedulingPanel.runScheduleTextView.SetText(result.runSchedule)
	s.schedulingPanel.executionTextView.SetText(result.execution)
	s.modified = true
}

func (s *processorDetailsScreen) buildSettingsPanel() *settingsPanel {
	const LABEL_WIDTH = 20
	var settingsPanel settingsPanel

	idTextView := tview.NewTextView().
		SetLabel("Id").
		SetLabelWidth(LABEL_WIDTH).
		SetText(s.processor.Component.GetId())
	typeTextView := tview.NewTextView().
		SetLabel("Type").
		SetLabelWidth(LABEL_WIDTH).
		SetText(s.processor.Component.GetType())
	bundleTextView := tview.NewTextView().
		SetLabel("Bundle").
		SetLabelWidth(LABEL_WIDTH).
		SetText(s.processor.Component.Bundle.GetArtifact() + "-" + s.processor.Component.Bundle.GetGroup() + " " + s.processor.Component.Bundle.GetVersion())
	settingsPanel.nameTextView = tview.NewTextView().
		SetLabel("Name").
		SetLabelWidth(LABEL_WIDTH).
		SetText(s.processor.Component.GetName())

	settingsPanel.penaltyDurationTextView = tview.NewTextView().
		SetLabel("Penalty Duration").
		SetLabelWidth(LABEL_WIDTH).
		SetText(s.processor.Component.Config.GetPenaltyDuration())

	settingsPanel.yieldDurationTextView = tview.NewTextView().
		SetLabel("Yield Duration").
		SetLabelWidth(LABEL_WIDTH).
		SetText(s.processor.Component.Config.GetYieldDuration())

	settingsPanel.bulletinLevelTextView = tview.NewTextView().
		SetLabel("Bulletin Level").
		SetLabelWidth(LABEL_WIDTH).
		SetText(s.processor.Component.Config.GetBulletinLevel())

	settingsPanel.helpFlex = tview.NewFlex()
	(*helpFlex)(settingsPanel.helpFlex).
		addHelpText("[ctrl-s]Save").
		addHelpText("[e]Edit").
		addHelpText("[2]Scheduling").
		addHelpText("[3]Properties").
		addHelpText("[4]Relationships")

	contentFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(idTextView, 2, 0, false).
		AddItem(typeTextView, 2, 0, false).
		AddItem(bundleTextView, 2, 0, false).
		AddItem(settingsPanel.nameTextView, 2, 0, false).
		AddItem(settingsPanel.penaltyDurationTextView, 2, 0, false).
		AddItem(settingsPanel.yieldDurationTextView, 2, 0, false).
		AddItem(settingsPanel.bulletinLevelTextView, 2, 0, false)

	contentFlex.SetTitle("Settings").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Rune() == 'e' {
				s.buildAndShowEditSettingsForm()
				return nil
			}
			return s.handleInput(event)
		})
	settingsPanel.flex = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(contentFlex, 0, 1, true).
		AddItem(settingsPanel.helpFlex, 1, 0, false)
	return &settingsPanel
}

func (s *processorDetailsScreen) buildSchedulingPanel() *schedulingPanel {
	const LABEL_WIDTH = 20

	var schedulingPanel schedulingPanel
	schedulingPanel.schedulingStrategyTextView = tview.NewTextView().
		SetLabel("Scheduling Strategy").
		SetLabelWidth(LABEL_WIDTH).
		SetText(s.processor.Component.Config.GetSchedulingStrategy())
	schedulingPanel.concurrentTaskTextView = tview.NewTextView().
		SetLabel("Concurrent Tasks").
		SetLabelWidth(LABEL_WIDTH).
		SetText(strconv.Itoa(int(s.processor.Component.Config.GetConcurrentlySchedulableTaskCount())))
	schedulingPanel.runScheduleTextView = tview.NewTextView().
		SetLabel("Run Schedule").
		SetLabelWidth(LABEL_WIDTH).
		SetText(s.processor.Component.Config.GetSchedulingPeriod())
	schedulingPanel.executionTextView = tview.NewTextView().
		SetLabel("Execution").
		SetLabelWidth(LABEL_WIDTH).
		SetText(s.processor.Component.Config.GetExecutionNode())

	schedulingPanel.helpFlex = tview.NewFlex()
	(*helpFlex)(schedulingPanel.helpFlex).
		addHelpText("[ctrl-s]Save").
		addHelpText("[e]Edit").
		addHelpText("[1]Settings").
		addHelpText("[3]Properties").
		addHelpText("[4]Relationships")

	contentFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(schedulingPanel.schedulingStrategyTextView, 2, 0, false).
		AddItem(schedulingPanel.concurrentTaskTextView, 2, 0, false).
		AddItem(schedulingPanel.runScheduleTextView, 2, 0, false).
		AddItem(schedulingPanel.executionTextView, 2, 0, false)
	contentFlex.SetTitle("Scheduling").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Rune() == 'e' {
				s.buildAndShowEditSchedulingForm()
				return nil
			}
			return s.handleInput(event)
		})

	schedulingPanel.flex = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(contentFlex, 0, 1, true).
		AddItem(schedulingPanel.helpFlex, 1, 0, false)
	return &schedulingPanel
}

func (s *processorDetailsScreen) buildPropertiesPanel() *propertiesPanel {
	var propertiesPanel propertiesPanel
	propertiesPanel.listView = tview.NewList()
	propertiesPanel.listView.SetTitle(s.processor.Component.GetName()).
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true).
		SetInputCapture(s.handleInput)

	for k, v := range s.processor.Component.Config.GetProperties() {
		k, v := k, v
		descriptor := s.processor.Component.Config.GetDescriptors()[k]
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
		propertiesPanel.listView.AddItem(displayName, value, 0, func() { s.onPropertySelected(k) })
	}
	propertiesPanel.helpFlex = tview.NewFlex()
	(*helpFlex)(propertiesPanel.helpFlex).
		addHelpText("[ctrl-s]Save").
		addHelpText("[1]Settings").
		addHelpText("[2]Scheduling").
		addHelpText("[4]Relationships")
	propertiesPanel.flex = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(propertiesPanel.listView, 0, 1, true).
		AddItem(propertiesPanel.helpFlex, 1, 0, false)
	return &propertiesPanel
}

func (s *processorDetailsScreen) buildRelationshipsPanel() *relationshipsPanel {
	var relationshipPanel relationshipsPanel
	relationshipPanel.listView = tview.NewList()
	relationshipPanel.listView.SetTitle("Relationships").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)
	for i, relationship := range s.processor.Component.Relationships {
		i, relationship := i, relationship
		relationshipPanel.listView.AddItem(relationship.GetName(), buildRelationshipSecondaryText(relationship), 0, func() {
			s.onRelationshipSelected(i)
		})
	}

	const LABEL_WIDTH = 30
	rightColumn := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().
			SetLabel("Number of Retry Attempts").
			SetLabelWidth(LABEL_WIDTH).
			SetText(strconv.Itoa(int(s.processor.Component.Config.GetRetryCount()))),
			2,
			0,
			false).
		AddItem(tview.NewTextView().
			SetLabel("Retry Back Off Policy").
			SetLabelWidth(LABEL_WIDTH).
			SetText(s.processor.Component.Config.GetBackoffMechanism()),
			2,
			0,
			false).
		AddItem(tview.NewTextView().
			SetLabel("Retry Maximum Back Off Period").
			SetLabelWidth(LABEL_WIDTH).
			SetText(s.processor.Component.Config.GetMaxBackoffPeriod()),
			2,
			0,
			false)
	rightColumn.SetTitle("Retry").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true)

	relationshipsFlex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(relationshipPanel.listView, 0, 1, true).
		AddItem(rightColumn, 0, 1, false)
	relationshipsFlex.SetInputCapture(s.handleInput)
	relationshipPanel.helpFlex = tview.NewFlex()
	(*helpFlex)(relationshipPanel.helpFlex).
		addHelpText("[ctrl-s]Save").
		addHelpText("[1]Settings").
		addHelpText("[2]Scheduling").
		addHelpText("[3]Properties")
	relationshipPanel.flex = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(relationshipsFlex, 0, 1, true).
		AddItem(relationshipPanel.helpFlex, 1, 0, true)
	return &relationshipPanel
}

func (s *processorDetailsScreen) getRelationshipNames() ([]string, []string) {
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
	return autoTerminatedRelationships, retriedRelationships
}

func (s *processorDetailsScreen) buildAndShowEditSettingsForm() {
	result := settingsFormResult{
		name:            s.processor.Component.GetName(),
		penaltyDuration: s.processor.Component.Config.GetPenaltyDuration(),
		yieldDuration:   s.processor.Component.Config.GetYieldDuration(),
		bulletinLevel:   s.processor.Component.Config.GetBulletinLevel(),
	}
	bulletinLevels := []string{
		"DEBUG",
		"INFO",
		"WARN",
		"ERROR",
		"NONE",
	}
	form := tview.NewForm().
		AddInputField("Name", result.name, 0, nil, func(text string) {
			result.name = text
		}).
		AddInputField("Penalty Duration", result.penaltyDuration, 0, nil, func(text string) {
			result.penaltyDuration = text
		}).
		AddInputField("Yield Duration", result.yieldDuration, 0, nil, func(text string) {
			result.yieldDuration = text
		}).
		AddDropDown("Bulletin Level",
			bulletinLevels,
			slices.Index(bulletinLevels, result.bulletinLevel),
			func(option string, optionIndex int) {
				result.bulletinLevel = option
			}).
		AddButton("Save", func() {
			s.applySettings(result)
			s.pages.RemovePage("Edit Settings")
		}).
		AddButton("Cancel", func() {
			s.pages.RemovePage("Edit Settings")
		})
	form.SetTitle("Settings").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyEsc:
				s.pages.RemovePage("Edit Settings")
				return nil
			case tcell.KeyCtrlS:
				s.applySettings(result)
				s.pages.RemovePage("Edit Settings")
				return nil
			}
			return event
		})
	grid := tview.NewGrid().
		SetColumns(0, 100, 0).
		SetRows(0, 13, 0).
		AddItem(form, 1, 1, 1, 1, 0, 0, true)
	s.pages.AddPage("Edit Settings", grid, true, true)
}

func (s *processorDetailsScreen) buildAndShowEditSchedulingForm() {
	result := schedulingFormResult{
		schedulingStrategy: s.processor.Component.Config.GetSchedulingStrategy(),
		concurrentTask:     int(s.processor.Component.Config.GetConcurrentlySchedulableTaskCount()),
		runSchedule:        s.processor.Component.Config.GetSchedulingPeriod(),
		execution:          s.processor.Component.Config.GetExecutionNode(),
	}
	schedulingStrategies := []string{
		"TIMER_DRIVEN",
		"CRON_DRIVEN",
	}
	executions := []string{
		"ALL",
		"PRIMARY",
	}
	form := tview.NewForm().
		AddDropDown("Scheduling Strategy",
			schedulingStrategies,
			slices.Index(schedulingStrategies, result.schedulingStrategy),
			func(option string, optionIndex int) {
				result.schedulingStrategy = option
			}).
		AddInputField("Concurrent Tasks",
			strconv.Itoa(result.concurrentTask),
			0,
			tview.InputFieldInteger,
			func(text string) {
				// the input should be already validated by the acceptance callback
				concurrentTask, _ := strconv.Atoi(text)
				result.concurrentTask = concurrentTask
			}).
		AddInputField("Run Schedule", result.runSchedule, 0, nil, func(text string) {
			result.runSchedule = text
		}).
		AddDropDown("Execution",
			executions,
			slices.Index(executions, result.execution),
			func(option string, optionIndex int) {
				result.execution = option
			}).
		AddButton("Save", func() {
			s.applyScheduling(result)
			s.pages.RemovePage("Edit Scheduling")
		}).
		AddButton("Cancel", func() {
			s.pages.RemovePage("Edit Scheduling")
		})
	form.SetTitle("Scheduling").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyEsc:
				s.pages.RemovePage("Edit Scheduling")
				return nil
			case tcell.KeyCtrlS:
				s.applyScheduling(result)
				s.pages.RemovePage("Edit Scheduling")
				return nil
			}
			return event
		})
	grid := tview.NewGrid().
		SetColumns(0, 100, 0).
		SetRows(0, 13, 0).
		AddItem(form, 1, 1, 1, 1, 0, 0, true)
	s.pages.AddPage("Edit Scheduling", grid, true, true)
}

func (s *processorDetailsScreen) showModifiedModal() {
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
