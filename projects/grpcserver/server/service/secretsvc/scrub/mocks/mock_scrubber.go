// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/solo-projects/projects/grpcserver/server/service/secretsvc/scrub (interfaces: Scrubber)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
)

// MockScrubber is a mock of Scrubber interface.
type MockScrubber struct {
	ctrl     *gomock.Controller
	recorder *MockScrubberMockRecorder
}

// MockScrubberMockRecorder is the mock recorder for MockScrubber.
type MockScrubberMockRecorder struct {
	mock *MockScrubber
}

// NewMockScrubber creates a new mock instance.
func NewMockScrubber(ctrl *gomock.Controller) *MockScrubber {
	mock := &MockScrubber{ctrl: ctrl}
	mock.recorder = &MockScrubberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScrubber) EXPECT() *MockScrubberMockRecorder {
	return m.recorder
}

// Secret mocks base method.
func (m *MockScrubber) Secret(arg0 context.Context, arg1 *v1.Secret) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Secret", arg0, arg1)
}

// Secret indicates an expected call of Secret.
func (mr *MockScrubberMockRecorder) Secret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Secret", reflect.TypeOf((*MockScrubber)(nil).Secret), arg0, arg1)
}
