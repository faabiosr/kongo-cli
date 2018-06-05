package services

import (
	"context"
	"errors"
	api "github.com/fabiorphp/kongo"
	"net/http"
)

type (
	MockServices struct {
		Error bool
	}
)

func (mock *MockServices) Create(svc *api.Service) (*api.Service, *http.Response, error) {
	if mock.Error {
		return nil, nil, errors.New("Unable to connect")
	}

	service := &api.Service{
		Id: "a1",
	}

	return service, nil, nil
}

func (mock *MockServices) CreateWithContext(ctx context.Context, svc *api.Service) (*api.Service, *http.Response, error) {
	return nil, nil, nil
}

func (mock *MockServices) CreateByURL(svc *api.Service) (*api.Service, *http.Response, error) {
	return mock.Create(svc)
}

func (mock *MockServices) CreateByURLWithContext(ctx context.Context, svc *api.Service) (*api.Service, *http.Response, error) {
	return nil, nil, nil
}

func (mock *MockServices) Delete(idOrName string) (*http.Response, error) {
	return nil, nil
}

func (mock *MockServices) DeleteWithContext(ctx context.Context, idOrName string) (*http.Response, error) {
	return nil, nil
}

func (mock *MockServices) Get(idOrName string) (*api.Service, *http.Response, error) {
	return nil, nil, nil
}

func (mock *MockServices) GetWithContext(ctx context.Context, idOrName string) (*api.Service, *http.Response, error) {
	return nil, nil, nil
}

func (mock *MockServices) List(options *api.ListServicesOptions) ([]*api.Service, *http.Response, error) {
	return nil, nil, nil
}

func (mock *MockServices) ListWithContext(ctx context.Context, options *api.ListServicesOptions) ([]*api.Service, *http.Response, error) {
	return nil, nil, nil
}

func (mock *MockServices) Update(idOrName string, svc *api.Service) (*api.Service, *http.Response, error) {
	return nil, nil, nil
}

func (mock *MockServices) UpdateWithContext(ctx context.Context, idOrName string, svc *api.Service) (*api.Service, *http.Response, error) {
	return nil, nil, nil
}

func (mock *MockServices) UpdateByURL(idOrName string, svc *api.Service) (*api.Service, *http.Response, error) {
	return nil, nil, nil
}

func (mock *MockServices) UpdateByURLWithContext(ctx context.Context, idOrName string, svc *api.Service) (*api.Service, *http.Response, error) {
	return nil, nil, nil
}
