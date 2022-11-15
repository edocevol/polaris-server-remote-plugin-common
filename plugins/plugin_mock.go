// Code generated by MockGen. DO NOT EDIT.
// Source: plugin.go

// Package plugins is a generated GoMock package.
package plugins

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	api "github.com/polaris-contrib/polaris-server-remote-plugin-common/api"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Call mocks base method
func (m *MockService) Call(ctx context.Context, request *api.Request) (*api.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", ctx, request)
	ret0, _ := ret[0].(*api.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call
func (mr *MockServiceMockRecorder) Call(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockService)(nil).Call), ctx, request)
}