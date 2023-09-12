package app

import (
	"dam4rus/nifi-tui/internal/pkg/nifiapi"
	"slices"
)

type processGroupComponents struct {
	id            string
	processGroups *nifiapi.ProcessGroupsEntity
	processors    *nifiapi.ProcessorsEntity
	connections   *nifiapi.ConnectionsEntity
	funnels       *nifiapi.FunnelsEntity
	inputPorts    *nifiapi.InputPortsEntity
	outputPorts   *nifiapi.OutputPortsEntity
}

func (pgc *processGroupComponents) findComponent(componentType, id string) Component {
	switch componentType {
	case "PROCESSOR":
		return pgc.findProcessorComponent(id)
	case "FUNNEL":
		return pgc.findFunnelComponent(id)
	case "INPUT_PORT":
		return pgc.findInputPortComponent(id)
	case "OUTPUT_PORT":
		return pgc.findOutputPortComponent(id)
	}
	return nil
}

func (pgc *processGroupComponents) findProcessorComponent(id string) Component {
	for _, processor := range pgc.processors.Processors {
		if processor.Component.GetId() == id {
			return processor.Component
		}
	}
	return nil
}

func (pgc *processGroupComponents) findFunnelComponent(id string) Component {
	for _, funnel := range pgc.funnels.Funnels {
		if funnel.Component.GetId() == id {
			return &FunnelComponent{
				Id: funnel.Component.GetId(),
			}
		}
	}
	return nil
}

func (pgc *processGroupComponents) findInputPortComponent(id string) Component {
	for _, inputPort := range pgc.inputPorts.InputPorts {
		if inputPort.Component.GetId() == id {
			return inputPort.Component
		}
	}
	return nil
}

func (pgc *processGroupComponents) findOutputPortComponent(id string) Component {
	for _, outputPort := range pgc.outputPorts.OutputPorts {
		if outputPort.Component.GetId() == id {
			return outputPort.Component
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

type AllowableValues []string

func (av AllowableValues) IsBoolean() bool {
	return len(av) == 2 && slices.Contains(av, "true") && slices.Contains(av, "false")
}