// Code generated by MockGen. DO NOT EDIT.
// Source: cluster.go

// Package store is a generated GoMock package.
package store

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/strugglehonor/KCS/internal/model"
)

// MockClusterStore is a mock of ClusterStore interface.
type MockClusterStore struct {
	ctrl     *gomock.Controller
	recorder *MockClusterStoreMockRecorder
}

// MockClusterStoreMockRecorder is the mock recorder for MockClusterStore.
type MockClusterStoreMockRecorder struct {
	mock *MockClusterStore
}

// NewMockClusterStore creates a new mock instance.
func NewMockClusterStore(ctrl *gomock.Controller) *MockClusterStore {
	mock := &MockClusterStore{ctrl: ctrl}
	mock.recorder = &MockClusterStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterStore) EXPECT() *MockClusterStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockClusterStore) Create(ctx context.Context, cluster *model.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, cluster)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockClusterStoreMockRecorder) Create(ctx, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClusterStore)(nil).Create), ctx, cluster)
}