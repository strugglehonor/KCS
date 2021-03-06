// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package v1 is a generated GoMock package.
package v1

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Cluster mocks base method.
func (m *MockService) Cluster() ClusterSrv {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cluster")
	ret0, _ := ret[0].(ClusterSrv)
	return ret0
}

// Cluster indicates an expected call of Cluster.
func (mr *MockServiceMockRecorder) Cluster() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockService)(nil).Cluster))
}

// Deployment mocks base method.
func (m *MockService) Deployment() DeploymentSrv {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deployment")
	ret0, _ := ret[0].(DeploymentSrv)
	return ret0
}

// Deployment indicates an expected call of Deployment.
func (mr *MockServiceMockRecorder) Deployment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deployment", reflect.TypeOf((*MockService)(nil).Deployment))
}

// Pod mocks base method.
func (m *MockService) Pod() PodSrv {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pod")
	ret0, _ := ret[0].(PodSrv)
	return ret0
}

// Pod indicates an expected call of Pod.
func (mr *MockServiceMockRecorder) Pod() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pod", reflect.TypeOf((*MockService)(nil).Pod))
}
