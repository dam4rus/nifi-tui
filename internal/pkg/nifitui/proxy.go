package nifitui

import (
	"context"
	"dam4rus/nifi-tui/internal/pkg/nifiapi"
	"net/http"
	"strconv"
	"sync"
)

type nifiComponentListEntity interface {
	nifiapi.ProcessGroupsEntity |
		nifiapi.ProcessorsEntity |
		nifiapi.FunnelsEntity |
		nifiapi.InputPortsEntity |
		nifiapi.OutputPortsEntity |
		nifiapi.ConnectionsEntity |
		nifiapi.RemoteProcessGroupsEntity
}

type processGroupProxy struct {
	service        *nifiapi.ProcessGroupsAPIService
	processGroupId string
}

func newProcessGroupProxy(apiClient *nifiapi.APIClient, processGroupId string) *processGroupProxy {
	return &processGroupProxy{
		service:        apiClient.ProcessGroupsAPI,
		processGroupId: processGroupId,
	}
}

func (pgp *processGroupProxy) GetVersion() (string, *http.Response, error) {
	processGroup, response, err := pgp.service.GetProcessGroup(context.Background(), pgp.processGroupId).
		Execute()
	if err != nil {
		return "", response, err
	}
	return strconv.Itoa(int(processGroup.Revision.GetVersion())), nil, nil
}

func (pgp *processGroupProxy) Remove() (*http.Response, error) {
	version, response, err := pgp.GetVersion()
	if err != nil {
		return response, err
	}
	_, response, err = pgp.service.RemoveProcessGroup(context.Background(), pgp.processGroupId).
		Version(version).
		Execute()
	return response, err
}

type asyncProcessGroupProxy struct {
	processGroupProxy
	ctx          context.Context
	cancelFunc   context.CancelFunc
	waitGroup    sync.WaitGroup
	errorChannel chan error
}

func newAsyncProcessGroupProxy(apiClient *nifiapi.APIClient, processGroupId string) *asyncProcessGroupProxy {
	ctx, cancelFunc := context.WithCancel(context.Background())
	return &asyncProcessGroupProxy{
		processGroupProxy: *newProcessGroupProxy(apiClient, processGroupId),
		ctx:               ctx,
		cancelFunc:        cancelFunc,
		errorChannel:      make(chan error),
	}
}

func (apgg *asyncProcessGroupProxy) getProcessGroups() <-chan *nifiapi.ProcessGroupsEntity {
	return (*asyncEntitiesGetterProxy[nifiapi.ProcessGroupsEntity])(apgg).getEntities(
		func() (*nifiapi.ProcessGroupsEntity, *http.Response, error) {
			return apgg.service.GetProcessGroups(apgg.ctx, apgg.processGroupId).Execute()
		})
}

func (apgg *asyncProcessGroupProxy) getProcessors() <-chan *nifiapi.ProcessorsEntity {
	return (*asyncEntitiesGetterProxy[nifiapi.ProcessorsEntity])(apgg).getEntities(
		func() (*nifiapi.ProcessorsEntity, *http.Response, error) {
			return apgg.service.GetProcessors(apgg.ctx, apgg.processGroupId).Execute()
		})
}

func (apgg *asyncProcessGroupProxy) getConnections() <-chan *nifiapi.ConnectionsEntity {
	return (*asyncEntitiesGetterProxy[nifiapi.ConnectionsEntity])(apgg).getEntities(
		func() (*nifiapi.ConnectionsEntity, *http.Response, error) {
			return apgg.service.GetConnections(apgg.ctx, apgg.processGroupId).Execute()
		})
}

func (apgg *asyncProcessGroupProxy) getFunnels() <-chan *nifiapi.FunnelsEntity {
	return (*asyncEntitiesGetterProxy[nifiapi.FunnelsEntity])(apgg).getEntities(
		func() (*nifiapi.FunnelsEntity, *http.Response, error) {
			return apgg.service.GetFunnels(apgg.ctx, apgg.processGroupId).Execute()
		})
}

func (apgg *asyncProcessGroupProxy) getInputPorts() <-chan *nifiapi.InputPortsEntity {
	return (*asyncEntitiesGetterProxy[nifiapi.InputPortsEntity])(apgg).getEntities(
		func() (*nifiapi.InputPortsEntity, *http.Response, error) {
			return apgg.service.GetInputPorts(apgg.ctx, apgg.processGroupId).Execute()
		})
}

func (apgg *asyncProcessGroupProxy) getOutputPorts() <-chan *nifiapi.OutputPortsEntity {
	return (*asyncEntitiesGetterProxy[nifiapi.OutputPortsEntity])(apgg).getEntities(
		func() (*nifiapi.OutputPortsEntity, *http.Response, error) {
			return apgg.service.GetOutputPorts(apgg.ctx, apgg.processGroupId).Execute()
		})
}

func (apgg *asyncProcessGroupProxy) getRemoteProcessGroups() <-chan *nifiapi.RemoteProcessGroupsEntity {
	return (*asyncEntitiesGetterProxy[nifiapi.RemoteProcessGroupsEntity])(apgg).getEntities(
		func() (*nifiapi.RemoteProcessGroupsEntity, *http.Response, error) {
			return apgg.service.GetRemoteProcessGroups(apgg.ctx, apgg.processGroupId).Execute()
		})
}

type asyncEntitiesGetterProxy[T nifiComponentListEntity] asyncProcessGroupProxy

func (aclg *asyncEntitiesGetterProxy[T]) getEntities(getter func() (*T, *http.Response, error)) <-chan *T {
	entitiesChannel := make(chan *T, 1)
	go func(ctx context.Context) {
		defer aclg.waitGroup.Done()
		entities, _, err := getter()
		if err != nil {
			aclg.errorChannel <- err
			return
		}
		entitiesChannel <- entities
	}(aclg.ctx)
	return entitiesChannel
}

type processorProxy struct {
	service     *nifiapi.ProcessorsAPIService
	processorId string
}

func newProcessorProxy(apiClient *nifiapi.APIClient, processorId string) *processorProxy {
	return &processorProxy{
		service:     apiClient.ProcessorsAPI,
		processorId: processorId,
	}
}

func (pp *processorProxy) GetVersion() (string, *http.Response, error) {
	processor, response, err := pp.service.GetProcessor(context.Background(), pp.processorId).
		Execute()
	if err != nil {
		return "", response, err
	}
	return strconv.Itoa(int(processor.Revision.GetVersion())), nil, nil
}

func (pp *processorProxy) Remove() (*http.Response, error) {
	version, response, err := pp.GetVersion()
	if err != nil {
		return response, err
	}
	_, response, err = pp.service.DeleteProcessor(context.Background(), pp.processorId).
		Version(version).
		Execute()
	return response, err
}

type connectionProxy struct {
	service      *nifiapi.ConnectionsAPIService
	connectionId string
}

func newConnectionProxy(apiClient *nifiapi.APIClient, connectionId string) *connectionProxy {
	return &connectionProxy{
		service:      apiClient.ConnectionsAPI,
		connectionId: connectionId,
	}
}

func (cp *connectionProxy) GetVersion() (string, *http.Response, error) {
	connection, response, err := cp.service.GetConnection(context.Background(), cp.connectionId).
		Execute()
	if err != nil {
		return "", response, err
	}
	return strconv.Itoa(int(connection.Revision.GetVersion())), nil, nil
}

func (cp *connectionProxy) Remove() (*http.Response, error) {
	version, response, err := cp.GetVersion()
	if err != nil {
		return response, err
	}
	_, response, err = cp.service.DeleteConnection(context.Background(), cp.connectionId).
		Version(version).
		Execute()
	return response, err
}

type funnelProxy struct {
	service  *nifiapi.FunnelAPIService
	funnelId string
}

func newFunnelProxy(apiClient *nifiapi.APIClient, funnelId string) *funnelProxy {
	return &funnelProxy{
		service:  apiClient.FunnelAPI,
		funnelId: funnelId,
	}
}

func (fp *funnelProxy) GetVersion() (string, *http.Response, error) {
	connection, response, err := fp.service.GetFunnel(context.Background(), fp.funnelId).
		Execute()
	if err != nil {
		return "", response, err
	}
	return strconv.Itoa(int(connection.Revision.GetVersion())), nil, nil
}

func (fp *funnelProxy) Remove() (*http.Response, error) {
	version, response, err := fp.GetVersion()
	if err != nil {
		return response, err
	}
	_, response, err = fp.service.RemoveFunnel(context.Background(), fp.funnelId).
		Version(version).
		Execute()
	return response, err
}

type inputPortProxy struct {
	service     *nifiapi.InputPortsAPIService
	inputPortId string
}

func newInputPortProxy(apiClient *nifiapi.APIClient, inputPortId string) *inputPortProxy {
	return &inputPortProxy{
		service:     apiClient.InputPortsAPI,
		inputPortId: inputPortId,
	}
}

func (ipp *inputPortProxy) GetVersion() (string, *http.Response, error) {
	connection, response, err := ipp.service.GetInputPort(context.Background(), ipp.inputPortId).
		Execute()
	if err != nil {
		return "", response, err
	}
	return strconv.Itoa(int(connection.Revision.GetVersion())), nil, nil
}

func (ipp *inputPortProxy) Remove() (*http.Response, error) {
	version, response, err := ipp.GetVersion()
	if err != nil {
		return response, err
	}
	_, response, err = ipp.service.RemoveInputPort(context.Background(), ipp.inputPortId).
		Version(version).
		Execute()
	return response, err
}

type outputPortProxy struct {
	service      *nifiapi.OutputPortsAPIService
	outputPortId string
}

func newOutputPortProxy(apiClient *nifiapi.APIClient, outputPortId string) *outputPortProxy {
	return &outputPortProxy{
		service:      apiClient.OutputPortsAPI,
		outputPortId: outputPortId,
	}
}

func (opp *outputPortProxy) GetVersion() (string, *http.Response, error) {
	connection, response, err := opp.service.GetOutputPort(context.Background(), opp.outputPortId).
		Execute()
	if err != nil {
		return "", response, err
	}
	return strconv.Itoa(int(connection.Revision.GetVersion())), nil, nil
}

func (opp *outputPortProxy) Remove() (*http.Response, error) {
	version, response, err := opp.GetVersion()
	if err != nil {
		return response, err
	}
	_, response, err = opp.service.RemoveOutputPort(context.Background(), opp.outputPortId).
		Version(version).
		Execute()
	return response, err
}

type remoteProcessGroupProxy struct {
	service              *nifiapi.RemoteProcessGroupsAPIService
	remoteProcessGroupId string
}

func newRemoteProcessGroupProxy(apiClient *nifiapi.APIClient, remoteProcessGroupId string) *remoteProcessGroupProxy {
	return &remoteProcessGroupProxy{
		service:              apiClient.RemoteProcessGroupsAPI,
		remoteProcessGroupId: remoteProcessGroupId,
	}
}

func (rpg *remoteProcessGroupProxy) GetVersion() (string, *http.Response, error) {
	connection, response, err := rpg.service.GetRemoteProcessGroup(context.Background(), rpg.remoteProcessGroupId).
		Execute()
	if err != nil {
		return "", response, err
	}
	return strconv.Itoa(int(connection.Revision.GetVersion())), nil, nil
}

func (rpg *remoteProcessGroupProxy) Remove() (*http.Response, error) {
	version, response, err := rpg.GetVersion()
	if err != nil {
		return response, err
	}
	_, response, err = rpg.service.RemoveRemoteProcessGroup(context.Background(), rpg.remoteProcessGroupId).
		Version(version).
		Execute()
	return response, err
}
