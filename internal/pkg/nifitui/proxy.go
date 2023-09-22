package nifitui

import (
	"context"
	"dam4rus/nifi-tui/internal/pkg/nifiapi"
	"sync"
)

type processGroupProxy struct {
	app            *App
	ctx            context.Context
	cancelFunc     context.CancelFunc
	waitGroup      sync.WaitGroup
	errorChannel   chan error
	processGroupId string
}

func (a *App) newProcessGroupProxy(processGroupId string) *processGroupProxy {
	ctx, cancelFunc := context.WithCancel(context.Background())
	return &processGroupProxy{
		app:            a,
		ctx:            ctx,
		cancelFunc:     cancelFunc,
		errorChannel:   make(chan error),
		processGroupId: processGroupId,
	}
}

func (pgp *processGroupProxy) getProcessGroups() <-chan *nifiapi.ProcessGroupsEntity {
	processGroupsChannel := make(chan *nifiapi.ProcessGroupsEntity, 1)
	go func(ctx context.Context) {
		defer pgp.waitGroup.Done()
		processGroups, _, err := pgp.app.apiClient.ProcessGroupsAPI.GetProcessGroups(ctx, pgp.processGroupId).Execute()
		if err != nil {
			pgp.errorChannel <- err
			return
		}
		processGroupsChannel <- processGroups
	}(pgp.ctx)
	return processGroupsChannel
}

func (pgp *processGroupProxy) getProcessors() <-chan *nifiapi.ProcessorsEntity {
	processorsChannel := make(chan *nifiapi.ProcessorsEntity, 1)
	go func(ctx context.Context) {
		defer pgp.waitGroup.Done()
		processors, _, err := pgp.app.apiClient.ProcessGroupsAPI.GetProcessors(ctx, pgp.processGroupId).Execute()
		if err != nil {
			pgp.errorChannel <- err
			return
		}
		processorsChannel <- processors
	}(pgp.ctx)
	return processorsChannel
}

func (pgp *processGroupProxy) getConnections() <-chan *nifiapi.ConnectionsEntity {
	connectionsChannel := make(chan *nifiapi.ConnectionsEntity, 1)
	go func(ctx context.Context) {
		defer pgp.waitGroup.Done()
		connections, _, err := pgp.app.apiClient.ProcessGroupsAPI.GetConnections(ctx, pgp.processGroupId).Execute()
		if err != nil {
			pgp.errorChannel <- err
			return
		}
		connectionsChannel <- connections
	}(pgp.ctx)
	return connectionsChannel
}

func (pgp *processGroupProxy) getFunnels() <-chan *nifiapi.FunnelsEntity {
	funnelsChannel := make(chan *nifiapi.FunnelsEntity, 1)
	go func(ctx context.Context) {
		defer pgp.waitGroup.Done()
		funnels, _, err := pgp.app.apiClient.ProcessGroupsAPI.GetFunnels(ctx, pgp.processGroupId).Execute()
		if err != nil {
			pgp.errorChannel <- err
			return
		}
		funnelsChannel <- funnels
	}(pgp.ctx)
	return funnelsChannel
}

func (pgp *processGroupProxy) getInputPorts() <-chan *nifiapi.InputPortsEntity {
	inputPortsChannel := make(chan *nifiapi.InputPortsEntity, 1)
	go func(ctx context.Context) {
		defer pgp.waitGroup.Done()
		inputPorts, _, err := pgp.app.apiClient.ProcessGroupsAPI.GetInputPorts(ctx, pgp.processGroupId).Execute()
		if err != nil {
			pgp.errorChannel <- err
			return
		}
		inputPortsChannel <- inputPorts
	}(pgp.ctx)
	return inputPortsChannel
}

func (pgp *processGroupProxy) getOutputPorts() <-chan *nifiapi.OutputPortsEntity {
	outputPortsChannel := make(chan *nifiapi.OutputPortsEntity, 1)
	go func(ctx context.Context) {
		defer pgp.waitGroup.Done()
		outputPorts, _, err := pgp.app.apiClient.ProcessGroupsAPI.GetOutputPorts(ctx, pgp.processGroupId).Execute()
		if err != nil {
			pgp.errorChannel <- err
			return
		}
		outputPortsChannel <- outputPorts
	}(pgp.ctx)
	return outputPortsChannel
}
