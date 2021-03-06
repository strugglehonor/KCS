// Code generated by MockGen. DO NOT EDIT.
// Source: pod.go

// Package store is a generated GoMock package.
package store

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/strugglehonor/KCS/internal/model"
)

// MockPodStore is a mock of PodStore interface.
type MockPodStore struct {
	ctrl     *gomock.Controller
	recorder *MockPodStoreMockRecorder
}

// MockPodStoreMockRecorder is the mock recorder for MockPodStore.
type MockPodStoreMockRecorder struct {
	mock *MockPodStore
}

// NewMockPodStore creates a new mock instance.
func NewMockPodStore(ctrl *gomock.Controller) *MockPodStore {
	mock := &MockPodStore{ctrl: ctrl}
	mock.recorder = &MockPodStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodStore) EXPECT() *MockPodStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPodStore) Create(ctx context.Context, pod *model.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, pod)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockPodStoreMockRecorder) Create(ctx, pod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPodStore)(nil).Create), ctx, pod)
}

// Delete mocks base method.
func (m *MockPodStore) Delete(ctx context.Context, pod *model.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, pod)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPodStoreMockRecorder) Delete(ctx, pod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPodStore)(nil).Delete), ctx, pod)
}

// GetAllPod mocks base method.
func (m *MockPodStore) GetAllPod(ctx context.Context, limit, offset *int64) ([]*model.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPod", ctx, limit, offset)
	ret0, _ := ret[0].([]*model.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPod indicates an expected call of GetAllPod.
func (mr *MockPodStoreMockRecorder) GetAllPod(ctx, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPod", reflect.TypeOf((*MockPodStore)(nil).GetAllPod), ctx, limit, offset)
}

// GetPodByName mocks base method.
func (m *MockPodStore) GetPodByName(ctx context.Context, podName string) (*model.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodByName", ctx, podName)
	ret0, _ := ret[0].(*model.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodByName indicates an expected call of GetPodByName.
func (mr *MockPodStoreMockRecorder) GetPodByName(ctx, podName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodByName", reflect.TypeOf((*MockPodStore)(nil).GetPodByName), ctx, podName)
}

// InsertPod mocks base method.
func (m *MockPodStore) InsertPod(ctx context.Context, pod *model.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertPod", ctx, pod)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertPod indicates an expected call of InsertPod.
func (mr *MockPodStoreMockRecorder) InsertPod(ctx, pod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertPod", reflect.TypeOf((*MockPodStore)(nil).InsertPod), ctx, pod)
}

// Update mocks base method.
func (m *MockPodStore) Update(ctx context.Context, pod *model.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, pod)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPodStoreMockRecorder) Update(ctx, pod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPodStore)(nil).Update), ctx, pod)
}
