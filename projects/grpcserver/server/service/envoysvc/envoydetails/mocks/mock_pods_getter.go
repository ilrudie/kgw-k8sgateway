// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/core/v1 (interfaces: PodsGetter,PodInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v1beta1 "k8s.io/api/policy/v1beta1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/kubernetes/typed/core/v1"
	rest "k8s.io/client-go/rest"
)

// MockPodsGetter is a mock of PodsGetter interface.
type MockPodsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockPodsGetterMockRecorder
}

// MockPodsGetterMockRecorder is the mock recorder for MockPodsGetter.
type MockPodsGetterMockRecorder struct {
	mock *MockPodsGetter
}

// NewMockPodsGetter creates a new mock instance.
func NewMockPodsGetter(ctrl *gomock.Controller) *MockPodsGetter {
	mock := &MockPodsGetter{ctrl: ctrl}
	mock.recorder = &MockPodsGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodsGetter) EXPECT() *MockPodsGetterMockRecorder {
	return m.recorder
}

// Pods mocks base method.
func (m *MockPodsGetter) Pods(arg0 string) v11.PodInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pods", arg0)
	ret0, _ := ret[0].(v11.PodInterface)
	return ret0
}

// Pods indicates an expected call of Pods.
func (mr *MockPodsGetterMockRecorder) Pods(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pods", reflect.TypeOf((*MockPodsGetter)(nil).Pods), arg0)
}

// MockPodInterface is a mock of PodInterface interface.
type MockPodInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPodInterfaceMockRecorder
}

// MockPodInterfaceMockRecorder is the mock recorder for MockPodInterface.
type MockPodInterfaceMockRecorder struct {
	mock *MockPodInterface
}

// NewMockPodInterface creates a new mock instance.
func NewMockPodInterface(ctrl *gomock.Controller) *MockPodInterface {
	mock := &MockPodInterface{ctrl: ctrl}
	mock.recorder = &MockPodInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodInterface) EXPECT() *MockPodInterfaceMockRecorder {
	return m.recorder
}

// Bind mocks base method.
func (m *MockPodInterface) Bind(arg0 *v1.Binding) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bind", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bind indicates an expected call of Bind.
func (mr *MockPodInterfaceMockRecorder) Bind(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bind", reflect.TypeOf((*MockPodInterface)(nil).Bind), arg0)
}

// Create mocks base method.
func (m *MockPodInterface) Create(arg0 *v1.Pod) (*v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPodInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPodInterface)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockPodInterface) Delete(arg0 string, arg1 *v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPodInterfaceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPodInterface)(nil).Delete), arg0, arg1)
}

// DeleteCollection mocks base method.
func (m *MockPodInterface) DeleteCollection(arg0 *v10.DeleteOptions, arg1 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockPodInterfaceMockRecorder) DeleteCollection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockPodInterface)(nil).DeleteCollection), arg0, arg1)
}

// Evict mocks base method.
func (m *MockPodInterface) Evict(arg0 *v1beta1.Eviction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Evict", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Evict indicates an expected call of Evict.
func (mr *MockPodInterfaceMockRecorder) Evict(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Evict", reflect.TypeOf((*MockPodInterface)(nil).Evict), arg0)
}

// Get mocks base method.
func (m *MockPodInterface) Get(arg0 string, arg1 v10.GetOptions) (*v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPodInterfaceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPodInterface)(nil).Get), arg0, arg1)
}

// GetEphemeralContainers mocks base method.
func (m *MockPodInterface) GetEphemeralContainers(arg0 string, arg1 v10.GetOptions) (*v1.EphemeralContainers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEphemeralContainers", arg0, arg1)
	ret0, _ := ret[0].(*v1.EphemeralContainers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEphemeralContainers indicates an expected call of GetEphemeralContainers.
func (mr *MockPodInterfaceMockRecorder) GetEphemeralContainers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEphemeralContainers", reflect.TypeOf((*MockPodInterface)(nil).GetEphemeralContainers), arg0, arg1)
}

// GetLogs mocks base method.
func (m *MockPodInterface) GetLogs(arg0 string, arg1 *v1.PodLogOptions) *rest.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogs", arg0, arg1)
	ret0, _ := ret[0].(*rest.Request)
	return ret0
}

// GetLogs indicates an expected call of GetLogs.
func (mr *MockPodInterfaceMockRecorder) GetLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogs", reflect.TypeOf((*MockPodInterface)(nil).GetLogs), arg0, arg1)
}

// List mocks base method.
func (m *MockPodInterface) List(arg0 v10.ListOptions) (*v1.PodList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(*v1.PodList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPodInterfaceMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPodInterface)(nil).List), arg0)
}

// Patch mocks base method.
func (m *MockPodInterface) Patch(arg0 string, arg1 types.PatchType, arg2 []byte, arg3 ...string) (*v1.Pod, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockPodInterfaceMockRecorder) Patch(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockPodInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockPodInterface) Update(arg0 *v1.Pod) (*v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPodInterfaceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPodInterface)(nil).Update), arg0)
}

// UpdateEphemeralContainers mocks base method.
func (m *MockPodInterface) UpdateEphemeralContainers(arg0 string, arg1 *v1.EphemeralContainers) (*v1.EphemeralContainers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEphemeralContainers", arg0, arg1)
	ret0, _ := ret[0].(*v1.EphemeralContainers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEphemeralContainers indicates an expected call of UpdateEphemeralContainers.
func (mr *MockPodInterfaceMockRecorder) UpdateEphemeralContainers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEphemeralContainers", reflect.TypeOf((*MockPodInterface)(nil).UpdateEphemeralContainers), arg0, arg1)
}

// UpdateStatus mocks base method.
func (m *MockPodInterface) UpdateStatus(arg0 *v1.Pod) (*v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockPodInterfaceMockRecorder) UpdateStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockPodInterface)(nil).UpdateStatus), arg0)
}

// Watch mocks base method.
func (m *MockPodInterface) Watch(arg0 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockPodInterfaceMockRecorder) Watch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockPodInterface)(nil).Watch), arg0)
}
