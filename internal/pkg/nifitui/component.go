package nifitui

import (
	"dam4rus/nifi-tui/internal/pkg/nifiapi"
	"slices"
)

type component interface {
	GetId() string
	CreateService(apiClient *nifiapi.APIClient) entityService
}

type displayableComponent interface {
	GetId() string
	GetDisplayName() string
}

type connectionSource interface {
	GetId() string
	GetGroupId() string
	GetConnectionSourceType() string
}

type connectionDestination interface {
	GetId() string
	GetGroupId() string
	GetConnectionDestinationType() string
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
	return SymbolProcessor + p.name
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
}

func newInputPort(entity nifiapi.PortEntity) *inputPortEntity {
	return &inputPortEntity{
		id:            entity.Component.GetId(),
		name:          entity.Component.GetName(),
		parentGroupId: entity.Component.GetParentGroupId(),
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
	return "INPUT_PORT"
}

func (ip *inputPortEntity) GetConnectionDestinationType() string {
	return ip.GetConnectionSourceType()
}

func (ip *inputPortEntity) GetDisplayName() string {
	return SymbolInputPort + ip.name
}

type outputPortEntity struct {
	id            string
	name          string
	parentGroupId string
}

func newOutputPortEntity(entity nifiapi.PortEntity) *outputPortEntity {
	return &outputPortEntity{
		id:            entity.Component.GetId(),
		name:          entity.Component.GetName(),
		parentGroupId: entity.Component.GetParentGroupId(),
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
	return "OUTPUT_PORT"
}

func (op *outputPortEntity) GetConnectionDestinationType() string {
	return op.GetConnectionSourceType()
}

func (op *outputPortEntity) GetDisplayName() string {
	return SymbolOutputPort + op.name
}

type connectionEntity struct {
	id              string
	sourceType      string
	sourceId        string
	destinationType string
	destinationId   string
}

func newConnectionEntity(entity nifiapi.ConnectionEntity) *connectionEntity {
	return &connectionEntity{
		id:              entity.Component.GetId(),
		sourceType:      entity.Component.Source.GetType(),
		sourceId:        entity.Component.Source.GetId(),
		destinationType: entity.Component.Destination.GetType(),
		destinationId:   entity.Component.Destination.GetId(),
	}
}

func (c *connectionEntity) GetId() string {
	return c.id
}

func (c *connectionEntity) CreateService(apiClient *nifiapi.APIClient) entityService {
	return newConnectionService(apiClient, c.id)
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

func (pgc *processGroupComponents) findConnectable(id string) connectionSource {
	if processGroup := pgc.processGroups[id]; processGroup != nil {
		return processGroup
	}
	if processor := pgc.processors[id]; processor != nil {
		return processor
	}
	if remoteProcessGroup := pgc.remoteProcessGroups[id]; remoteProcessGroup != nil {
		return remoteProcessGroup
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

func (pgc *processGroupComponents) findDisplayableConnectionSource(connection *connectionEntity) displayableComponent {
	switch connection.sourceType {
	case "PROCESSOR":
		return pgc.processGroups[connection.sourceId]
	case "FUNNEL":
		return pgc.funnels[connection.sourceId]
	case "INPUT_PORT":
		return pgc.inputPorts[connection.sourceId]
	case "OUTPUT_PORT":
		return pgc.outputPorts[connection.sourceId]
	}
	return nil
}

func (pgc *processGroupComponents) findDisplayableConnectionDestination(connection *connectionEntity) displayableComponent {
	switch connection.destinationType {
	case "PROCESSOR":
		return pgc.processGroups[connection.destinationType]
	case "FUNNEL":
		return pgc.funnels[connection.destinationType]
	case "INPUT_PORT":
		return pgc.inputPorts[connection.destinationType]
	case "OUTPUT_PORT":
		return pgc.outputPorts[connection.destinationType]
	}
	return nil
}

type AllowableValues []string

func (av AllowableValues) IsBoolean() bool {
	return len(av) == 2 && slices.Contains(av, "true") && slices.Contains(av, "false")
}
