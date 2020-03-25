// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/solo-projects/projects/grpcserver/server/service/virtualservicesvc/mutation (interfaces: Mutator)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	mutation "github.com/solo-io/solo-projects/projects/grpcserver/server/service/virtualservicesvc/mutation"
)

// MockMutator is a mock of Mutator interface.
type MockMutator struct {
	ctrl     *gomock.Controller
	recorder *MockMutatorMockRecorder
}

// MockMutatorMockRecorder is the mock recorder for MockMutator.
type MockMutatorMockRecorder struct {
	mock *MockMutator
}

// NewMockMutator creates a new mock instance.
func NewMockMutator(ctrl *gomock.Controller) *MockMutator {
	mock := &MockMutator{ctrl: ctrl}
	mock.recorder = &MockMutatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMutator) EXPECT() *MockMutatorMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMutator) Create(arg0 *core.ResourceRef, arg1 mutation.Mutation) (*v1.VirtualService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*v1.VirtualService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMutatorMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMutator)(nil).Create), arg0, arg1)
}

// Update mocks base method.
func (m *MockMutator) Update(arg0 *core.ResourceRef, arg1 mutation.Mutation) (*v1.VirtualService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*v1.VirtualService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMutatorMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMutator)(nil).Update), arg0, arg1)
}
