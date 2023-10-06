package nifitui

import (
	"context"
	"net/http"
	"strconv"
	"sync"

	"github.com/dam4rus/nifi-tui/internal/pkg/nifiapi"
)

type entityService interface {
	Remove() (*http.Response, error)
}

type portGetter interface {
	GetInputPorts(ctx context.Context) ([]inputPortEntity, *http.Response, error)
	GetOutputPorts(ctx context.Context) ([]outputPortEntity, *http.Response, error)
}

type componentStateChanger interface {
	Start() (*http.Response, error)
	Stop() (*http.Response, error)
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

func (pgp *processGroupService) GetInputPorts(ctx context.Context) ([]inputPortEntity, *http.Response, error) {
	inputPorts, response, err := pgp.service.GetInputPorts(ctx, pgp.processGroupId).Execute()
	if err != nil {
		return nil, response, err
	}
	var result []inputPortEntity
	for _, inputPort := range inputPorts.InputPorts {
		result = append(result, *newInputPort(inputPort))
	}
	return result, nil, nil
}

func (pgp *processGroupService) GetOutputPorts(ctx context.Context) ([]outputPortEntity, *http.Response, error) {
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

func (pgp *processGroupService) getComponentsAsync(ctx context.Context) *getComponentsTask {
	result := newGetComponentsTask()
	result.waitGroup.Add(7)
	go func() {
		defer result.waitGroup.Done()
		processGroups, _, err := pgp.getProcessGroups(ctx)
		if err != nil {
			result.errorChannel <- err
			return
		}
		result.processGroupsChannel <- processGroups
	}()
	go func() {
		defer result.waitGroup.Done()
		processors, _, err := pgp.getProcessors(ctx)
		if err != nil {
			result.errorChannel <- err
			return
		}
		result.processorsChannel <- processors
	}()
	go func() {
		defer result.waitGroup.Done()
		connections, _, err := pgp.getConnections(ctx)
		if err != nil {
			result.errorChannel <- err
			return
		}
		result.connectionsChannel <- connections
	}()
	go func() {
		defer result.waitGroup.Done()
		funnels, _, err := pgp.getFunnels(ctx)
		if err != nil {
			result.errorChannel <- err
			return
		}
		result.funnelsChannel <- funnels
	}()
	go func() {
		defer result.waitGroup.Done()
		inputPorts, _, err := pgp.GetInputPorts(ctx)
		if err != nil {
			result.errorChannel <- err
			return
		}
		result.inputPortsChannel <- inputPorts
	}()
	go func() {
		defer result.waitGroup.Done()
		outputPorts, _, err := pgp.GetOutputPorts(ctx)
		if err != nil {
			result.errorChannel <- err
			return
		}
		result.outputPortsChannel <- outputPorts
	}()
	go func() {
		defer result.waitGroup.Done()
		remoteProcessGroups, _, err := pgp.getRemoteProcessGroups(ctx)
		if err != nil {
			result.errorChannel <- err
			return
		}
		result.remoteProcessGroupsChannel <- remoteProcessGroups
	}()
	return &result
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

func (pgp *processGroupService) createFunnel() (*nifiapi.FunnelEntity, *http.Response, error) {
	funnelEntity := nifiapi.FunnelEntity{
		Component: &nifiapi.FunnelDTO{},
		Revision:  createDefaultRevision(),
	}
	return pgp.service.CreateFunnel(context.Background(), pgp.processGroupId).
		Body(funnelEntity).
		Execute()
}

func (pgp *processGroupService) createInputPort(portName string) (*nifiapi.PortEntity, *http.Response, error) {
	portEntity := nifiapi.PortEntity{
		Component: &nifiapi.PortDTO{
			Name: &portName,
		},
		Revision: createDefaultRevision(),
	}
	return pgp.service.CreateInputPort(context.Background(), pgp.processGroupId).
		Body(portEntity).
		Execute()
}

func (pgp *processGroupService) createOutputPort(portName string) (*nifiapi.PortEntity, *http.Response, error) {
	portEntity := nifiapi.PortEntity{
		Component: &nifiapi.PortDTO{
			Name: &portName,
		},
		Revision: createDefaultRevision(),
	}
	return pgp.service.CreateOutputPort(context.Background(), pgp.processGroupId).
		Body(portEntity).
		Execute()
}

func (pgp *processGroupService) createRemoteProcessGroup(remoteProcessGroupName string,
	uri string) (*nifiapi.RemoteProcessGroupEntity, *http.Response, error) {
	remoteProcessGroupEntity := nifiapi.RemoteProcessGroupEntity{
		Component: &nifiapi.RemoteProcessGroupDTO{
			Name:      &remoteProcessGroupName,
			TargetUri: &uri,
		},
		Revision: createDefaultRevision(),
	}
	return pgp.service.CreateRemoteProcessGroup(context.Background(), pgp.processGroupId).
		Body(remoteProcessGroupEntity).
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

func (pp *processorService) Start() (*http.Response, error) {
	return pp.setState("RUNNING")
}

func (pp *processorService) Stop() (*http.Response, error) {
	return pp.setState("STOPPED")
}

func (pp *processorService) setState(state string) (*http.Response, error) {
	processor, response, err := pp.service.GetProcessor(context.Background(), pp.processorId).
		Execute()
	if err != nil {
		return response, err
	}
	_, response, err = pp.service.UpdateRunStatus(context.Background(), pp.processorId).
		Body(nifiapi.ProcessorRunStatusEntity{
			Revision: processor.Revision,
			State:    &state,
		}).
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

func (ipp *inputPortService) Start() (*http.Response, error) {
	return ipp.changeState("RUNNING")
}

func (ipp *inputPortService) Stop() (*http.Response, error) {
	return ipp.changeState("STOPPED")
}

func (ipp *inputPortService) changeState(state string) (*http.Response, error) {
	outputPort, response, err := ipp.service.GetInputPort(context.Background(), ipp.inputPortId).
		Execute()
	if err != nil {
		return response, err
	}
	_, response, err = ipp.service.UpdateRunStatus(context.Background(), ipp.inputPortId).
		Body(nifiapi.PortRunStatusEntity{
			Revision: outputPort.Revision,
			State:    &state,
		}).
		Execute()
	return response, err
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

func (opp *outputPortService) Start() (*http.Response, error) {
	return opp.changeState("RUNNING")
}

func (opp *outputPortService) Stop() (*http.Response, error) {
	return opp.changeState("STOPPED")
}

func (opp *outputPortService) changeState(state string) (*http.Response, error) {
	outputPort, response, err := opp.service.GetOutputPort(context.Background(), opp.outputPortId).
		Execute()
	if err != nil {
		return response, err
	}
	_, response, err = opp.service.UpdateRunStatus(context.Background(), opp.outputPortId).
		Body(nifiapi.PortRunStatusEntity{
			Revision: outputPort.Revision,
			State:    &state,
		}).
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

func (rpg *remoteProcessGroupService) GetInputPorts(ctx context.Context) ([]inputPortEntity, *http.Response, error) {
	remoteProcessGroup, response, err := rpg.service.GetRemoteProcessGroup(ctx, rpg.remoteProcessGroupId).Execute()
	if err != nil {
		return nil, response, err
	}
	var result []inputPortEntity
	for _, inputPort := range remoteProcessGroup.Component.Contents.GetInputPorts() {
		result = append(result, inputPortEntity{
			id:            inputPort.GetId(),
			name:          inputPort.GetName(),
			parentGroupId: inputPort.GetGroupId(),
			isRemote:      true,
		})
	}
	return result, nil, nil
}

func (rpg *remoteProcessGroupService) GetOutputPorts(ctx context.Context) ([]outputPortEntity, *http.Response, error) {
	remoteProcessGroup, response, err := rpg.service.GetRemoteProcessGroup(ctx, rpg.remoteProcessGroupId).Execute()
	if err != nil {
		return nil, response, err
	}
	var result []outputPortEntity
	for _, outputPort := range remoteProcessGroup.Component.Contents.GetOutputPorts() {
		result = append(result, outputPortEntity{
			id:            outputPort.GetId(),
			name:          outputPort.GetName(),
			parentGroupId: outputPort.GetGroupId(),
			isRemote:      true,
		})
	}
	return result, nil, nil
}

func createDefaultRevision() *nifiapi.RevisionDTO {
	version := int64(0)
	return &nifiapi.RevisionDTO{
		Version: &version,
	}
}

type getComponentsTask struct {
	errorChannel               chan error
	processGroupsChannel       chan []processGroupEntity
	processorsChannel          chan []processorEntity
	connectionsChannel         chan []connectionEntity
	funnelsChannel             chan []funnelEntity
	inputPortsChannel          chan []inputPortEntity
	outputPortsChannel         chan []outputPortEntity
	remoteProcessGroupsChannel chan []remoteProcessGroupEntity
	waitGroup                  sync.WaitGroup
}

func newGetComponentsTask() getComponentsTask {
	return getComponentsTask{
		errorChannel:               make(chan error),
		processGroupsChannel:       make(chan []processGroupEntity, 1),
		processorsChannel:          make(chan []processorEntity, 1),
		connectionsChannel:         make(chan []connectionEntity, 1),
		funnelsChannel:             make(chan []funnelEntity, 1),
		inputPortsChannel:          make(chan []inputPortEntity, 1),
		outputPortsChannel:         make(chan []outputPortEntity, 1),
		remoteProcessGroupsChannel: make(chan []remoteProcessGroupEntity, 1),
	}
}

func (gct *getComponentsTask) await() (*processGroupComponents, error) {
	gct.waitGroup.Wait()
	select {
	case err := <-gct.errorChannel:
		return nil, err
	default:
	}
	components := newProcessGroupComponents()
	for _, processGroup := range <-gct.processGroupsChannel {
		processGroup := processGroup
		components.processGroups[processGroup.id] = &processGroup
	}
	for _, processor := range <-gct.processorsChannel {
		processor := processor
		components.processors[processor.id] = &processor
	}
	for _, connection := range <-gct.connectionsChannel {
		connection := connection
		components.connections[connection.id] = &connection
	}
	for _, funnel := range <-gct.funnelsChannel {
		funnel := funnel
		components.funnels[funnel.id] = &funnel
	}
	for _, inputPort := range <-gct.inputPortsChannel {
		inputPort := inputPort
		components.inputPorts[inputPort.id] = &inputPort
	}
	for _, outputPort := range <-gct.outputPortsChannel {
		outputPort := outputPort
		components.outputPorts[outputPort.id] = &outputPort
	}
	for _, remoteProcessGroup := range <-gct.remoteProcessGroupsChannel {
		remoteProcessGroup := remoteProcessGroup
		components.remoteProcessGroups[remoteProcessGroup.id] = &remoteProcessGroup
	}
	return &components, nil
}
