package node

import (
	"context"
	"errors"
	api "github.com/fabiorphp/kongo"
	"net/http"
)

type (
	MockNode struct {
		Error bool
	}
)

func (mock *MockNode) Info() (*api.NodeInfo, *http.Response, error) {
	if mock.Error {
		return nil, nil, errors.New("Unable to connect")
	}

	info := &api.NodeInfo{
		Configuration: &api.NodeInfoConfiguration{},
		Plugins:       &api.NodeInfoPlugins{},
		Timers:        &api.NodeInfoTimers{},
	}

	return info, nil, nil
}

func (mock *MockNode) InfoWithContext(ctx context.Context) (*api.NodeInfo, *http.Response, error) {
	return nil, nil, nil
}

func (mock *MockNode) Status() (*api.NodeStatus, *http.Response, error) {
	if mock.Error {
		return nil, nil, errors.New("Unable to connect")
	}

	status := &api.NodeStatus{
		Database: &api.NodeStatusDatabase{
			Reachable: true,
		},
		Server: &api.NodeStatusServer{
			ConnectionsAccepted: 10,
			ConnectionsActive:   10,
			ConnectionsHandled:  10,
			ConnectionsReading:  10,
			ConnectionsWaiting:  10,
			ConnectionsWriting:  10,
			TotalRequests:       10,
		},
	}

	return status, nil, nil
}

func (mock *MockNode) StatusWithContext(ctx context.Context) (*api.NodeStatus, *http.Response, error) {
	return nil, nil, nil
}
