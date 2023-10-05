package nifitui

import (
	"dam4rus/nifi-tui/internal/pkg/nifiapi"
	"slices"
)

type component interface {
	GetId() string
	CreateService(apiClient *nifiapi.APIClient) entityService
}

type connectionSource interface {
	component
	GetGroupId() string
	GetConnectionSourceType() string
}

type connectionDestination interface {
	component
	GetGroupId() string
	GetConnectionDestinationType() string
}

type runnableComponent interface {
	GetState() string
	CreateStateChangerService(apiclient *nifiapi.APIClient) componentStateChanger
}

type processGroupEntity struct {
	id            string
	name          string
	parentGroupId string
}

func newProcessGroupEntity(entity nifiapi.ProcessGroupEntity) *processGroupEntity {
	return &processGroupEntity{
		id:            entity.Component.GetId(),
		name:          entity.Component.GetName(),
		parentGroupId: entity.Component.GetParentGroupId(),
	}
}

func (pg *processGroupEntity) GetId() string {
	return pg.id
}

func (pg *processGroupEntity) GetGroupId() string {
	return pg.parentGroupId
}

func (pg *processGroupEntity) CreateService(apiClient *nifiapi.APIClient) entityService {
	return newProcessGroupService(apiClient, pg.id)
}

func (pg *processGroupEntity) GetConnectionSourceType() string {
	return "OUTPUT_PORT"
}

func (pg *processGroupEntity) GetConnectionDestinationType() string {
	return "INPUT_PORT"
}

func (pg *processGroupEntity) GetDisplayName() string {
	return SymbolProcessGroup + pg.name
}

type processorEntity struct {
	id            string
	name          string
	parentGroupId string
	relationships []relationship
	state         string
}

func newProcessorEntity(entity nifiapi.ProcessorEntity) *processorEntity {
	component := entity.Component
	var relationships []relationship
	for _, relationship := range entity.Component.Relationships {
		relationships = append(relationships, newRelationship(relationship))
	}
	return &processorEntity{
		id:            component.GetId(),
		name:          component.GetName(),
		parentGroupId: component.GetParentGroupId(),
		relationships: relationships,
		state:         component.GetState(),
	}
}

func (p *processorEntity) GetId() string {
	return p.id
}

func (p *processorEntity) GetGroupId() string {
	return p.parentGroupId
}

func (p *processorEntity) CreateService(apiClient *nifiapi.APIClient) entityService {
	return newProcessorService(apiClient, p.id)
}

func (p *processorEntity) GetConnectionSourceType() string {
	return "PROCESSOR"
}

func (p *processorEntity) GetConnectionDestinationType() string {
	return p.GetConnectionSourceType()
}

func (p *processorEntity) GetDisplayName() string {
	return SymbolProcessor + p.name + " " + componentStateSymbol(p.state)
}

func (p *processorEntity) CreateStateChangerService(apiClient *nifiapi.APIClient) componentStateChanger {
	return newProcessorService(apiClient, p.id)
}

func (p *processorEntity) GetState() string {
	return p.state
}

func (p *processorEntity) getRelationshipNames() []string {
	var relationshipNames []string
	for _, relationship := range p.relationships {
		relationshipNames = append(relationshipNames, relationship.name)
	}
	return relationshipNames
}

type relationship struct {
	name          string
	autoTerminate bool
	retry         bool
}

func newRelationship(dto nifiapi.RelationshipDTO) relationship {
	return relationship{
		name:          dto.GetName(),
		autoTerminate: dto.GetAutoTerminate(),
		retry:         dto.GetRetry(),
	}
}

type funnelEntity struct {
	id            string
	parentGroupId string
}

func newFunnelEntity(entity nifiapi.FunnelEntity) *funnelEntity {
	return &funnelEntity{
		id:            entity.Component.GetId(),
		parentGroupId: entity.Component.GetParentGroupId(),
	}
}

func (f *funnelEntity) GetId() string {
	return f.id
}

func (f *funnelEntity) GetGroupId() string {
	return f.parentGroupId
}

func (f *funnelEntity) CreateService(apiClient *nifiapi.APIClient) entityService {
	return newFunnelService(apiClient, f.id)
}

func (f *funnelEntity) GetConnectionSourceType() string {
	return "FUNNEL"
}

func (f *funnelEntity) GetConnectionDestinationType() string {
	return f.GetConnectionSourceType()
}

func (f *funnelEntity) GetDisplayName() string {
	return "Ⴤ"
}

type inputPortEntity struct {
	id            string
	name          string
	parentGroupId string
	isRemote      bool
	state         string
}

func newInputPort(entity nifiapi.PortEntity) *inputPortEntity {
	return &inputPortEntity{
		id:            entity.Component.GetId(),
		name:          entity.Component.GetName(),
		parentGroupId: entity.Component.GetParentGroupId(),
		isRemote:      false,
		state:         entity.Component.GetState(),
	}
}

func (ip *inputPortEntity) GetId() string {
	return ip.id
}

func (ip *inputPortEntity) GetGroupId() string {
	return ip.parentGroupId
}

func (ip *inputPortEntity) CreateService(apiClient *nifiapi.APIClient) entityService {
	return newInputPortService(apiClient, ip.id)
}

func (ip *inputPortEntity) GetConnectionSourceType() string {
	if ip.isRemote {
		return "REMOTE_INPUT_PORT"
	}
	return "INPUT_PORT"
}

func (ip *inputPortEntity) GetConnectionDestinationType() string {
	return ip.GetConnectionSourceType()
}

func (ip *inputPortEntity) GetDisplayName() string {
	return SymbolInputPort + ip.name + " " + componentStateSymbol(ip.state)
}

func (ip *inputPortEntity) GetState() string {
	return ip.state
}

func (ip *inputPortEntity) CreateStateChangerService(apiclient *nifiapi.APIClient) componentStateChanger {
	return newInputPortService(apiclient, ip.id)
}

type outputPortEntity struct {
	id            string
	name          string
	parentGroupId string
	isRemote      bool
	state         string
}

func newOutputPortEntity(entity nifiapi.PortEntity) *outputPortEntity {
	return &outputPortEntity{
		id:            entity.Component.GetId(),
		name:          entity.Component.GetName(),
		parentGroupId: entity.Component.GetParentGroupId(),
		isRemote:      false,
		state:         entity.Component.GetState(),
	}
}

func (op *outputPortEntity) GetId() string {
	return op.id
}

func (op *outputPortEntity) GetGroupId() string {
	return op.parentGroupId
}

func (op *outputPortEntity) CreateService(apiClient *nifiapi.APIClient) entityService {
	return newOutputPortService(apiClient, op.id)
}

func (op *outputPortEntity) GetConnectionSourceType() string {
	if op.isRemote {
		return "REMOTE_OUTPUT_PORT"
	}
	return "OUTPUT_PORT"
}

func (op *outputPortEntity) GetConnectionDestinationType() string {
	return op.GetConnectionSourceType()
}

func (op *outputPortEntity) GetDisplayName() string {
	return SymbolOutputPort + op.name + " " + componentStateSymbol(op.state)
}

func (op *outputPortEntity) GetState() string {
	return op.state
}

func (op *outputPortEntity) CreateStateChangerService(apiclient *nifiapi.APIClient) componentStateChanger {
	return newOutputPortService(apiclient, op.id)
}

type connectionEntity struct {
	id              string
	sourceType      string
	sourceId        string
	sourceName      string
	destinationType string
	destinationId   string
	desinationName  string
}

func newConnectionEntity(entity nifiapi.ConnectionEntity) *connectionEntity {
	source := entity.Component.Source
	destination := entity.Component.Destination
	return &connectionEntity{
		id:              entity.Component.GetId(),
		sourceType:      source.GetType(),
		sourceId:        source.GetId(),
		sourceName:      source.GetName(),
		destinationType: destination.GetType(),
		destinationId:   destination.GetId(),
		desinationName:  destination.GetName(),
	}
}

func (c *connectionEntity) GetId() string {
	return c.id
}

func (c *connectionEntity) CreateService(apiClient *nifiapi.APIClient) entityService {
	return newConnectionService(apiClient, c.id)
}

func (c *connectionEntity) GetDisplayName() string {
	sourceSymbol := getSymbolOfConnectionSourceType(c.sourceType)
	destinationSymbol := getSymbolOfConnectionSourceType(c.destinationType)
	return sourceSymbol + c.sourceName + "[" + c.sourceId + "]" +
		" -> " + destinationSymbol + c.desinationName + "[" + c.destinationId + "]"
}

func getSymbolOfConnectionSourceType(sourceType string) string {
	switch sourceType {
	case "PROCESSOR":
		return SymbolProcessor
	case "FUNNEL":
		return SymbolFunnel
	case "INPUT_PORT":
		return SymbolInputPort
	case "OUTPUT_PORT":
		return SymbolOutputPort
	default:
		return ""
	}
}

type remoteProcessGroupEntity struct {
	id            string
	name          string
	parentGroupId string
}

func newRemoteProcessGroupEntity(entity nifiapi.RemoteProcessGroupEntity) *remoteProcessGroupEntity {
	return &remoteProcessGroupEntity{
		id:            entity.Component.GetId(),
		name:          entity.Component.GetName(),
		parentGroupId: entity.Component.GetParentGroupId(),
	}
}

func (rpg *remoteProcessGroupEntity) GetId() string {
	return rpg.id
}

func (rpg *remoteProcessGroupEntity) GetGroupId() string {
	return rpg.parentGroupId
}

func (rpg *remoteProcessGroupEntity) CreateService(apiClient *nifiapi.APIClient) entityService {
	return newRemoteProcessGroupService(apiClient, rpg.id)
}

func (rpg *remoteProcessGroupEntity) GetConnectionSourceType() string {
	return "REMOTE_OUTPUT_PORT"
}

func (rpg *remoteProcessGroupEntity) GetConnectionDestinationType() string {
	return "REMOTE_INPUT_PORT"
}

func (rpg *remoteProcessGroupEntity) GetDisplayName() string {
	return SymbolRemoteProcessGroup + rpg.name
}

type processGroupComponents struct {
	processGroups       map[string]*processGroupEntity
	processors          map[string]*processorEntity
	remoteProcessGroups map[string]*remoteProcessGroupEntity
	connections         map[string]*connectionEntity
	funnels             map[string]*funnelEntity
	inputPorts          map[string]*inputPortEntity
	outputPorts         map[string]*outputPortEntity
}

func newProcessGroupComponents() processGroupComponents {
	return processGroupComponents{
		processGroups:       make(map[string]*processGroupEntity),
		processors:          make(map[string]*processorEntity),
		remoteProcessGroups: make(map[string]*remoteProcessGroupEntity),
		connections:         make(map[string]*connectionEntity),
		funnels:             make(map[string]*funnelEntity),
		inputPorts:          make(map[string]*inputPortEntity),
		outputPorts:         make(map[string]*outputPortEntity),
	}
}

func (pgc *processGroupComponents) findComponent(id string) component {
	if processGroup := pgc.processGroups[id]; processGroup != nil {
		return processGroup
	}
	if processor := pgc.processors[id]; processor != nil {
		return processor
	}
	if remoteProcessGroup := pgc.remoteProcessGroups[id]; remoteProcessGroup != nil {
		return remoteProcessGroup
	}
	if connection := pgc.connections[id]; connection != nil {
		return connection
	}
	if funnel := pgc.funnels[id]; funnel != nil {
		return funnel
	}
	if inputPort := pgc.inputPorts[id]; inputPort != nil {
		return inputPort
	}
	if outputPort := pgc.outputPorts[id]; outputPort != nil {
		return outputPort
	}
	return nil
}

type AllowableValues []string

func (av AllowableValues) IsBoolean() bool {
	return len(av) == 2 && slices.Contains(av, "true") && slices.Contains(av, "false")
}
