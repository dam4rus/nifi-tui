package app

import (
	"context"
	"dam4rus/nifi-tui/internal/pkg/nifiapi"
	"io"
	"slices"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ProcessorDetailsScreen struct {
	app                  *App
	processor            *nifiapi.ProcessorEntity
	pages                *tview.Pages
	propertyListView     *tview.List
	relationshipListView *tview.List
	modified             bool
}

func newProcessorDetailsScreen(app *App, processor *nifiapi.ProcessorEntity) *ProcessorDetailsScreen {
	processorDetailsScreen := &ProcessorDetailsScreen{
		app:                  app,
		processor:            processor,
		pages:                tview.NewPages(),
		propertyListView:     tview.NewList(),
		relationshipListView: tview.NewList(),
	}
	settingsFlex := processorDetailsScreen.buildSettingsPanel()
	processorDetailsScreen.buildPropertyListPanel()
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

	relationshipsFlex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(processorDetailsScreen.relationshipListView, 0, 1, true).
		AddItem(rightColumn, 0, 1, false)
	relationshipsFlex.SetInputCapture(processorDetailsScreen.handleInput)
	processorDetailsScreen.pages.
		AddAndSwitchToPage("Settings", settingsFlex, true).
		AddPage("Properties", processorDetailsScreen.propertyListView, true, false).
		AddPage("Relationships", relationshipsFlex, true, false)

	return processorDetailsScreen
}

func (s *ProcessorDetailsScreen) save() {
	autoTerminatedRelationships, retriedRelationships := s.getRelationshipNames()
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
		s.pages.SwitchToPage("Properties")
		return nil
	case '3':
		s.pages.SwitchToPage("Relationships")
		return nil
	}
	return event
}

func (s *ProcessorDetailsScreen) onPropertySelected(key string) {
	propertyDescriptor := (*s.processor.Component.Config.Descriptors)[key]
	var allowableValues AllowableValues
	for _, allowableValue := range propertyDescriptor.AllowableValues {
		allowableValues = append(allowableValues, *allowableValue.AllowableValue.Value)
	}

	if allowableValues.IsBoolean() {
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
		form.AddDropDown(key, allowableValueDisplayNames, slices.Index(allowableValues, newValue), func(_ string, optionIndex int) {
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

func (s *ProcessorDetailsScreen) buildSettingsPanel() *tview.Flex {
	const LABEL_WIDTH = 20
	processorBundle := s.processor.Component.GetBundle()

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().
			SetLabel("Id").
			SetLabelWidth(LABEL_WIDTH).
			SetText(s.processor.Component.GetId()),
			2,
			0,
			false).
		AddItem(tview.NewTextView().
			SetLabel("Type").
			SetLabelWidth(LABEL_WIDTH).
			SetText(s.processor.Component.GetType()),
			2,
			0,
			false).
		AddItem(tview.NewTextView().
			SetLabel("Bundle").
			SetLabelWidth(LABEL_WIDTH).
			SetText(processorBundle.GetArtifact()+"-"+processorBundle.GetGroup()+" "+processorBundle.GetVersion()),
			2,
			0,
			false).
		AddItem(tview.NewTextView().
			SetLabel("Name").
			SetLabelWidth(LABEL_WIDTH).
			SetText(s.processor.Component.GetName()),
			2,
			0,
			false).
		AddItem(tview.NewTextView().
			SetLabel("Penalty Duration").
			SetLabelWidth(LABEL_WIDTH).
			SetText(s.processor.Component.Config.GetPenaltyDuration()),
			2,
			0,
			false).
		AddItem(tview.NewTextView().
			SetLabel("Yield Duration").
			SetLabelWidth(LABEL_WIDTH).
			SetText(s.processor.Component.Config.GetYieldDuration()),
			2,
			0,
			false).
		AddItem(tview.NewTextView().
			SetLabel("Bulletin Level").
			SetLabelWidth(LABEL_WIDTH).
			SetText(s.processor.Component.Config.GetBulletinLevel()),
			2,
			0,
			false)

	flex.SetTitle("Settings").
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true).
		SetBorderPadding(1, 1, 1, 1).
		SetInputCapture(s.handleInput)
	return flex
}

func (s *ProcessorDetailsScreen) buildPropertyListPanel() {
	s.propertyListView.SetTitle(s.processor.Component.GetName()).
		SetTitleAlign(tview.AlignLeft).
		SetBorder(true).
		SetInputCapture(s.handleInput)

	s.propertyListView.Clear()
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
		s.propertyListView.AddItem(displayName, value, 0, func() { s.onPropertySelected(k) })
	}
}

func (s *ProcessorDetailsScreen) getRelationshipNames() ([]string, []string) {
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

func (s *ProcessorDetailsScreen) showModifiedModal() {
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
