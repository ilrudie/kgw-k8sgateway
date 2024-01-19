// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/gloo/projects/gloo/pkg/api/v1 (interfaces: UpstreamClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
	gomock "go.uber.org/mock/gomock"
)

// MockUpstreamClient is a mock of UpstreamClient interface
type MockUpstreamClient struct {
	ctrl     *gomock.Controller
	recorder *MockUpstreamClientMockRecorder
}

// MockUpstreamClientMockRecorder is the mock recorder for MockUpstreamClient
type MockUpstreamClientMockRecorder struct {
	mock *MockUpstreamClient
}

// NewMockUpstreamClient creates a new mock instance
func NewMockUpstreamClient(ctrl *gomock.Controller) *MockUpstreamClient {
	mock := &MockUpstreamClient{ctrl: ctrl}
	mock.recorder = &MockUpstreamClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpstreamClient) EXPECT() *MockUpstreamClientMockRecorder {
	return m.recorder
}

// BaseClient mocks base method
func (m *MockUpstreamClient) BaseClient() clients.ResourceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseClient")
	ret0, _ := ret[0].(clients.ResourceClient)
	return ret0
}

// BaseClient indicates an expected call of BaseClient
func (mr *MockUpstreamClientMockRecorder) BaseClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseClient", reflect.TypeOf((*MockUpstreamClient)(nil).BaseClient))
}

// Delete mocks base method
func (m *MockUpstreamClient) Delete(arg0, arg1 string, arg2 clients.DeleteOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockUpstreamClientMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUpstreamClient)(nil).Delete), arg0, arg1, arg2)
}

// List mocks base method
func (m *MockUpstreamClient) List(arg0 string, arg1 clients.ListOpts) (v1.UpstreamList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(v1.UpstreamList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockUpstreamClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUpstreamClient)(nil).List), arg0, arg1)
}

// Read mocks base method
func (m *MockUpstreamClient) Read(arg0, arg1 string, arg2 clients.ReadOpts) (*v1.Upstream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Upstream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockUpstreamClientMockRecorder) Read(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockUpstreamClient)(nil).Read), arg0, arg1, arg2)
}

// Register mocks base method
func (m *MockUpstreamClient) Register() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockUpstreamClientMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUpstreamClient)(nil).Register))
}

// Watch mocks base method
func (m *MockUpstreamClient) Watch(arg0 string, arg1 clients.WatchOpts) (<-chan v1.UpstreamList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(<-chan v1.UpstreamList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockUpstreamClientMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockUpstreamClient)(nil).Watch), arg0, arg1)
}

// Write mocks base method
func (m *MockUpstreamClient) Write(arg0 *v1.Upstream, arg1 clients.WriteOpts) (*v1.Upstream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(*v1.Upstream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockUpstreamClientMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockUpstreamClient)(nil).Write), arg0, arg1)
}
