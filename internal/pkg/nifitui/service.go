package nifitui

import (
	"context"
	"dam4rus/nifi-tui/internal/pkg/nifiapi"
	"net/http"
	"strconv"
)

type nifiComponentListEntity interface {
	[]processGroupEntity |
		[]processorEntity |
		[]funnelEntity |
		[]inputPortEntity |
		[]outputPortEntity |
		[]connectionEntity |
		[]remoteProcessGroupEntity
}

type entityService interface {
	Remove() (*http.Response, error)
}

type processGroupService struct {
	apiClient      *nifiapi.APIClient
	service        *nifiapi.ProcessGroupsAPIService
	processGroupId string
}

func newProcessGroupService(apiClient *nifiapi.APIClient, processGroupId string) *processGroupService {
	return &processGroupService{
		apiClient:      apiClient,
		service:        apiClient.ProcessGroupsAPI,
		processGroupId: processGroupId,
	}
}

func (pgp *processGroupService) Remove() (*http.Response, error) {
	version, response, err := pgp.getVersion()
	if err != nil {
		return response, err
	}
	_, response, err = pgp.service.RemoveProcessGroup(context.Background(), pgp.processGroupId).
		Version(version).
		Execute()
	return response, err
}

func (pgp *processGroupService) getVersion() (string, *http.Response, error) {
	processGroup, response, err := pgp.service.GetProcessGroup(context.Background(), pgp.processGroupId).
		Execute()
	if err != nil {
		return "", response, err
	}
	return strconv.Itoa(int(processGroup.Revision.GetVersion())), nil, nil
}

func (pgp *processGroupService) getProcessGroups(ctx context.Context) ([]processGroupEntity, *http.Response, error) {
	entities, response, err := pgp.service.GetProcessGroups(ctx, pgp.processGroupId).Execute()
	if err != nil {
		return nil, response, err
	}
	var processGroups []processGroupEntity
	for _, entity := range entities.ProcessGroups {
		processGroups = append(processGroups, *newProcessGroupEntity(entity))
	}
	return processGroups, nil, nil
}

func (pgp *processGroupService) getProcessors(ctx context.Context) ([]processorEntity, *http.Response, error) {
	entities, response, err := pgp.service.GetProcessors(ctx, pgp.processGroupId).Execute()
	if err != nil {
		return nil, response, err
	}
	var processors []processorEntity
	for _, entity := range entities.Processors {
		var relationships []relationship
		for _, relationship := range entity.Component.GetRelationships() {
			relationships = append(relationships, newRelationship(relationship))
		}
		processors = append(processors, *newProcessorEntity(entity))
	}
	return processors, nil, nil
}

func (pgp *processGroupService) getConnections(ctx context.Context) ([]connectionEntity, *http.Response, error) {
	entities, response, err := pgp.service.GetConnections(ctx, pgp.processGroupId).Execute()
	if err != nil {
		return nil, response, err
	}
	var connections []connectionEntity
	for _, entity := range entities.Connections {
		connections = append(connections, *newConnectionEntity(entity))
	}
	return connections, nil, nil
}

func (pgp *processGroupService) getFunnels(ctx context.Context) ([]funnelEntity, *http.Response, error) {
	entities, response, err := pgp.service.GetFunnels(ctx, pgp.processGroupId).Execute()
	if err != nil {
		return nil, response, err
	}
	var funnels []funnelEntity
	for _, entity := range entities.Funnels {
		funnels = append(funnels, *newFunnelEntity(entity))
	}
	return funnels, nil, nil
}

func (pgp *processGroupService) getInputPorts(ctx context.Context) ([]inputPortEntity, *http.Response, error) {
	inputPorts, response, err := pgp.service.GetInputPorts(ctx, pgp.processGroupId).Execute()
	if err != nil {
		return nil, response, err
	}
	var result []inputPortEntity
	for _, inputPort := range inputPorts.InputPorts {
		result = append(result, inputPortEntity{
			id:            inputPort.Component.GetId(),
			name:          inputPort.Component.GetName(),
			parentGroupId: inputPort.Component.GetParentGroupId(),
		})
	}
	return result, nil, nil
}

func (pgp *processGroupService) getOutputPorts(ctx context.Context) ([]outputPortEntity, *http.Response, error) {
	entities, response, err := pgp.service.GetOutputPorts(ctx, pgp.processGroupId).Execute()
	if err != nil {
		return nil, response, err
	}
	var outputPorts []outputPortEntity
	for _, entity := range entities.OutputPorts {
		outputPorts = append(outputPorts, *newOutputPortEntity(entity))
	}
	return outputPorts, nil, nil
}

func (pgp *processGroupService) getRemoteProcessGroups(ctx context.Context) ([]remoteProcessGroupEntity, *http.Response, error) {
	entities, response, err := pgp.service.GetRemoteProcessGroups(ctx, pgp.processGroupId).Execute()
	if err != nil {
		return nil, response, err
	}
	var remoteProcessGroups []remoteProcessGroupEntity
	for _, entity := range entities.RemoteProcessGroups {
		remoteProcessGroups = append(remoteProcessGroups, *newRemoteProcessGroupEntity(entity))
	}
	return remoteProcessGroups, nil, nil
}

func (pgp *processGroupService) createProcessGroup(processGroupName string) (*nifiapi.ProcessGroupEntity, *http.Response, error) {
	processGroupEntity := nifiapi.ProcessGroupEntity{
		Component: &nifiapi.ProcessGroupDTO{
			Name: &processGroupName,
		},
		Revision: createDefaultRevision(),
	}
	return pgp.service.CreateProcessGroup(context.Background(), pgp.processGroupId).
		Body(processGroupEntity).
		Execute()
}

func (pgp *processGroupService) createProcessor(processorType, processorName string) (*nifiapi.ProcessorEntity, *http.Response, error) {
	processorEntity := nifiapi.ProcessorEntity{
		Component: &nifiapi.ProcessorDTO{
			Name: &processorName,
			Type: &processorType,
		},
		Revision: createDefaultRevision(),
	}
	return pgp.service.CreateProcessor(context.Background(), pgp.processGroupId).
		Body(processorEntity).
		Execute()
}

func (pgp *processGroupService) createConnection(source connectionSource,
	destination connectionDestination,
	selectedRelationships []string,
) (*nifiapi.ConnectionEntity, *http.Response, error) {
	connectionEntity := nifiapi.ConnectionEntity{
		Component: &nifiapi.ConnectionDTO{
			SelectedRelationships: selectedRelationships,
			Source: &nifiapi.ConnectableDTO{
				Id:      source.GetId(),
				GroupId: source.GetGroupId(),
				Type:    source.GetConnectionSourceType(),
			},
			Destination: &nifiapi.ConnectableDTO{
				Id:      destination.GetId(),
				GroupId: destination.GetGroupId(),
				Type:    destination.GetConnectionDestinationType(),
			},
		},
		Revision: createDefaultRevision(),
	}
	return pgp.service.CreateConnection(context.Background(), pgp.processGroupId).
		Body(connectionEntity).
		Execute()
}

type processorService struct {
	service     *nifiapi.ProcessorsAPIService
	processorId string
}

func newProcessorService(apiClient *nifiapi.APIClient, processorId string) *processorService {
	return &processorService{
		service:     apiClient.ProcessorsAPI,
		processorId: processorId,
	}
}

func (pp *processorService) Remove() (*http.Response, error) {
	version, response, err := pp.getVersion()
	if err != nil {
		return response, err
	}
	_, response, err = pp.service.DeleteProcessor(context.Background(), pp.processorId).
		Version(version).
		Execute()
	return response, err
}

func (pp *processorService) getVersion() (string, *http.Response, error) {
	processor, response, err := pp.service.GetProcessor(context.Background(), pp.processorId).
		Execute()
	if err != nil {
		return "", response, err
	}
	return strconv.Itoa(int(processor.Revision.GetVersion())), nil, nil
}

type connectionService struct {
	service      *nifiapi.ConnectionsAPIService
	connectionId string
}

func newConnectionService(apiClient *nifiapi.APIClient, connectionId string) *connectionService {
	return &connectionService{
		service:      apiClient.ConnectionsAPI,
		connectionId: connectionId,
	}
}

func (cp *connectionService) Remove() (*http.Response, error) {
	version, response, err := cp.getVersion()
	if err != nil {
		return response, err
	}
	_, response, err = cp.service.DeleteConnection(context.Background(), cp.connectionId).
		Version(version).
		Execute()
	return response, err
}

func (cp *connectionService) getVersion() (string, *http.Response, error) {
	connection, response, err := cp.service.GetConnection(context.Background(), cp.connectionId).
		Execute()
	if err != nil {
		return "", response, err
	}
	return strconv.Itoa(int(connection.Revision.GetVersion())), nil, nil
}

type funnelService struct {
	service  *nifiapi.FunnelAPIService
	funnelId string
}

func newFunnelService(apiClient *nifiapi.APIClient, funnelId string) *funnelService {
	return &funnelService{
		service:  apiClient.FunnelAPI,
		funnelId: funnelId,
	}
}

func (fp *funnelService) Remove() (*http.Response, error) {
	version, response, err := fp.getVersion()
	if err != nil {
		return response, err
	}
	_, response, err = fp.service.RemoveFunnel(context.Background(), fp.funnelId).
		Version(version).
		Execute()
	return response, err
}

func (fp *funnelService) getVersion() (string, *http.Response, error) {
	funnel, response, err := fp.service.GetFunnel(context.Background(), fp.funnelId).
		Execute()
	if err != nil {
		return "", response, err
	}
	return strconv.Itoa(int(funnel.Revision.GetVersion())), nil, nil
}

type inputPortService struct {
	service     *nifiapi.InputPortsAPIService
	inputPortId string
}

func newInputPortService(apiClient *nifiapi.APIClient, inputPortId string) *inputPortService {
	return &inputPortService{
		service:     apiClient.InputPortsAPI,
		inputPortId: inputPortId,
	}
}

func (ipp *inputPortService) Remove() (*http.Response, error) {
	version, response, err := ipp.getVersion()
	if err != nil {
		return response, err
	}
	_, response, err = ipp.service.RemoveInputPort(context.Background(), ipp.inputPortId).
		Version(version).
		Execute()
	return response, err
}

func (ipp *inputPortService) getVersion() (string, *http.Response, error) {
	inputPort, response, err := ipp.service.GetInputPort(context.Background(), ipp.inputPortId).
		Execute()
	if err != nil {
		return "", response, err
	}
	return strconv.Itoa(int(inputPort.Revision.GetVersion())), nil, nil
}

type outputPortService struct {
	service      *nifiapi.OutputPortsAPIService
	outputPortId string
}

func newOutputPortService(apiClient *nifiapi.APIClient, outputPortId string) *outputPortService {
	return &outputPortService{
		service:      apiClient.OutputPortsAPI,
		outputPortId: outputPortId,
	}
}

func (opp *outputPortService) Remove() (*http.Response, error) {
	version, response, err := opp.getVersion()
	if err != nil {
		return response, err
	}
	_, response, err = opp.service.RemoveOutputPort(context.Background(), opp.outputPortId).
		Version(version).
		Execute()
	return response, err
}

func (opp *outputPortService) getVersion() (string, *http.Response, error) {
	outputPort, response, err := opp.service.GetOutputPort(context.Background(), opp.outputPortId).
		Execute()
	if err != nil {
		return "", response, err
	}
	return strconv.Itoa(int(outputPort.Revision.GetVersion())), nil, nil
}

type remoteProcessGroupService struct {
	service              *nifiapi.RemoteProcessGroupsAPIService
	remoteProcessGroupId string
}

func newRemoteProcessGroupService(apiClient *nifiapi.APIClient, remoteProcessGroupId string) *remoteProcessGroupService {
	return &remoteProcessGroupService{
		service:              apiClient.RemoteProcessGroupsAPI,
		remoteProcessGroupId: remoteProcessGroupId,
	}
}

func (rpg *remoteProcessGroupService) Remove() (*http.Response, error) {
	version, response, err := rpg.getVersion()
	if err != nil {
		return response, err
	}
	_, response, err = rpg.service.RemoveRemoteProcessGroup(context.Background(), rpg.remoteProcessGroupId).
		Version(version).
		Execute()
	return response, err
}

func (rpg *remoteProcessGroupService) getVersion() (string, *http.Response, error) {
	remoteProcessGroup, response, err := rpg.service.GetRemoteProcessGroup(context.Background(), rpg.remoteProcessGroupId).
		Execute()
	if err != nil {
		return "", response, err
	}
	return strconv.Itoa(int(remoteProcessGroup.Revision.GetVersion())), nil, nil
}

func (rpg *remoteProcessGroupService) getInputPorts() ([]inputPortEntity, *http.Response, error) {
	remoteProcessGroup, response, err := rpg.service.GetRemoteProcessGroup(context.Background(), rpg.remoteProcessGroupId).Execute()
	if err != nil {
		return nil, response, err
	}
	var result []inputPortEntity
	for _, inputPort := range remoteProcessGroup.Component.Contents.GetInputPorts() {
		result = append(result, inputPortEntity{
			id:            inputPort.GetId(),
			name:          inputPort.GetName(),
			parentGroupId: inputPort.GetGroupId(),
		})
	}
	return result, nil, nil
}

func (rpg *remoteProcessGroupService) getOutputPorts() ([]nifiapi.RemoteProcessGroupPortDTO, *http.Response, error) {
	outputPorts, response, err := rpg.service.GetRemoteProcessGroup(context.Background(), rpg.remoteProcessGroupId).Execute()
	if err != nil {
		return nil, response, err
	}
	return outputPorts.Component.Contents.GetOutputPorts(), nil, nil
}

func createDefaultRevision() *nifiapi.RevisionDTO {
	version := int64(0)
	return &nifiapi.RevisionDTO{
		Version: &version,
	}
}
