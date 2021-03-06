// Code generated by MockGen. DO NOT EDIT.
// Source: pod.go

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/strugglehonor/KCS/internal/model"
)

// MockPodSrv is a mock of PodSrv interface.
type MockPodSrv struct {
	ctrl     *gomock.Controller
	recorder *MockPodSrvMockRecorder
}

// MockPodSrvMockRecorder is the mock recorder for MockPodSrv.
type MockPodSrvMockRecorder struct {
	mock *MockPodSrv
}

// NewMockPodSrv creates a new mock instance.
func NewMockPodSrv(ctrl *gomock.Controller) *MockPodSrv {
	mock := &MockPodSrv{ctrl: ctrl}
	mock.recorder = &MockPodSrvMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodSrv) EXPECT() *MockPodSrvMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPodSrv) Create(ctx context.Context, pod *model.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, pod)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockPodSrvMockRecorder) Create(ctx, pod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPodSrv)(nil).Create), ctx, pod)
}

// Delete mocks base method.
func (m *MockPodSrv) Delete(ctx context.Context, pod *model.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, pod)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPodSrvMockRecorder) Delete(ctx, pod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPodSrv)(nil).Delete), ctx, pod)
}

// GetPodByName mocks base method.
func (m *MockPodSrv) GetPodByName(ctx context.Context, podName string) (*model.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodByName", ctx, podName)
	ret0, _ := ret[0].(*model.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodByName indicates an expected call of GetPodByName.
func (mr *MockPodSrvMockRecorder) GetPodByName(ctx, podName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodByName", reflect.TypeOf((*MockPodSrv)(nil).GetPodByName), ctx, podName)
}

// InsertPod mocks base method.
func (m *MockPodSrv) InsertPod(ctx context.Context, ID, name, namespace, status, creator, restartPolicy, imagePullPolicy, dnsPolicy, ip string, memLimit, cpulimit, gpulimit, lifelimit int32, env map[interface{}]interface{}, hostIPC, hostNetwork, hostPID bool, node model.Node, cluster model.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertPod", ctx, ID, name, namespace, status, creator, restartPolicy, imagePullPolicy, dnsPolicy, ip, memLimit, cpulimit, gpulimit, lifelimit, env, hostIPC, hostNetwork, hostPID, node, cluster)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertPod indicates an expected call of InsertPod.
func (mr *MockPodSrvMockRecorder) InsertPod(ctx, ID, name, namespace, status, creator, restartPolicy, imagePullPolicy, dnsPolicy, ip, memLimit, cpulimit, gpulimit, lifelimit, env, hostIPC, hostNetwork, hostPID, node, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertPod", reflect.TypeOf((*MockPodSrv)(nil).InsertPod), ctx, ID, name, namespace, status, creator, restartPolicy, imagePullPolicy, dnsPolicy, ip, memLimit, cpulimit, gpulimit, lifelimit, env, hostIPC, hostNetwork, hostPID, node, cluster)
}

// Update mocks base method.
func (m *MockPodSrv) Update(ctx context.Context, pod *model.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, pod)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPodSrvMockRecorder) Update(ctx, pod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPodSrv)(nil).Update), ctx, pod)
}
