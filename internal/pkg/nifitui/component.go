package nifitui

import (
	"dam4rus/nifi-tui/internal/pkg/nifiapi"
	"net/http"
	"slices"
)

const (
	ComponentTypeProcessGroup       = "Process Group"
	ComponentTypeProcessor          = "Processor"
	ComponentTypeFunnel             = "Funnel"
	ComponentTypeInputPort          = "Input Port"
	ComponentTypeOutputPort         = "Output Port"
	ComponentTypeConnection         = "Connection"
	ComponentTypeRemoteProcessGroup = "Remote Process Group"
)

type componentDescriptor interface {
	GetId() string
	GetName() string
}

type connectableComponent interface {
	GetComponentType() string
	GetId() string
	IntoProxy(apiClient *nifiapi.APIClient) componentProxy
}

type componentProxy interface {
	Remove() (*http.Response, error)
}

type processGroup struct {
	id string
}

func (pg *processGroup) GetComponentType() string {
	return ComponentTypeProcessGroup
}

func (pg *processGroup) GetId() string {
	return pg.id
}

func (pg *processGroup) IntoProxy(apiClient *nifiapi.APIClient) componentProxy {
	return newProcessGroupProxy(apiClient, pg.id)
}

type processor struct {
	id string
}

func (p *processor) GetComponentType() string {
	return ComponentTypeProcessor
}

func (p *processor) GetId() string {
	return p.id
}

func (p *processor) IntoProxy(apiClient *nifiapi.APIClient) componentProxy {
	return newProcessorProxy(apiClient, p.id)
}

type funnel struct {
	id string
}

func (f *funnel) GetComponentType() string {
	return ComponentTypeFunnel
}

func (f *funnel) GetId() string {
	return f.id
}

func (f *funnel) GetName() string {
	return "Ⴤ"
}

func (f *funnel) IntoProxy(apiClient *nifiapi.APIClient) componentProxy {
	return newFunnelProxy(apiClient, f.id)
}

type inputPort struct {
	id string
}

func (ip *inputPort) GetComponentType() string {
	return ComponentTypeInputPort
}

func (ip *inputPort) GetId() string {
	return ip.id
}

func (ip *inputPort) IntoProxy(apiClient *nifiapi.APIClient) componentProxy {
	return newInputPortProxy(apiClient, ip.id)
}

type outputPort struct {
	id string
}

func (op *outputPort) GetComponentType() string {
	return ComponentTypeInputPort
}

func (op *outputPort) GetId() string {
	return op.id
}

func (op *outputPort) IntoProxy(apiClient *nifiapi.APIClient) componentProxy {
	return newOutputPortProxy(apiClient, op.id)
}

type connection struct {
	id string
}

func (c *connection) GetComponentType() string {
	return ComponentTypeInputPort
}

func (c *connection) GetId() string {
	return c.id
}

func (c *connection) IntoProxy(apiClient *nifiapi.APIClient) componentProxy {
	return newConnectionProxy(apiClient, c.id)
}

type remoteProcessGroup struct {
	id string
}

func (rpg *remoteProcessGroup) GetComponentType() string {
	return ComponentTypeRemoteProcessGroup
}

func (rpg *remoteProcessGroup) GetId() string {
	return rpg.id
}

func (rpg *remoteProcessGroup) IntoProxy(apiClient *nifiapi.APIClient) componentProxy {
	return newRemoteProcessGroupProxy(apiClient, rpg.id)
}

type processGroupComponents struct {
	processGroups       []nifiapi.ProcessGroupEntity
	processors          []nifiapi.ProcessorEntity
	remoteProcessGroups []nifiapi.RemoteProcessGroupEntity
	connections         []nifiapi.ConnectionEntity
	funnels             []nifiapi.FunnelEntity
	inputPorts          []nifiapi.PortEntity
	outputPorts         []nifiapi.PortEntity
}

func (pgc *processGroupComponents) findComponent(componentType, id string) componentDescriptor {
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

func (pgc *processGroupComponents) findProcessorComponent(id string) componentDescriptor {
	for _, processor := range pgc.processors {
		if processor.Component.GetId() == id {
			return processor.Component
		}
	}
	return nil
}

func (pgc *processGroupComponents) findFunnelComponent(id string) componentDescriptor {
	for _, f := range pgc.funnels {
		if f.Component.GetId() == id {
			return &funnel{
				id: f.Component.GetId(),
			}
		}
	}
	return nil
}

func (pgc *processGroupComponents) findInputPortComponent(id string) componentDescriptor {
	for _, inputPort := range pgc.inputPorts {
		if inputPort.Component.GetId() == id {
			return inputPort.Component
		}
	}
	return nil
}

func (pgc *processGroupComponents) findOutputPortComponent(id string) componentDescriptor {
	for _, outputPort := range pgc.outputPorts {
		if outputPort.Component.GetId() == id {
			return outputPort.Component
		}
	}
	return nil
}

type AllowableValues []string

func (av AllowableValues) IsBoolean() bool {
	return len(av) == 2 && slices.Contains(av, "true") && slices.Contains(av, "false")
}
