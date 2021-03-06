// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package appdynamics is a generated GoMock package.
package appdynamics

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// GetAppDynamics mocks base method
func (m *MockStore) GetAppDynamics(chat uint32) (*AppDynamics, error) {
	ret := m.ctrl.Call(m, "GetAppDynamics", chat)
	ret0, _ := ret[0].(*AppDynamics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppDynamics indicates an expected call of GetAppDynamics
func (mr *MockStoreMockRecorder) GetAppDynamics(chat interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppDynamics", reflect.TypeOf((*MockStore)(nil).GetAppDynamics), chat)
}

// addAppDynamicsEndpoint mocks base method
func (m *MockStore) addAppDynamicsEndpoint(chat uint32, endpoint string) error {
	ret := m.ctrl.Call(m, "addAppDynamicsEndpoint", chat, endpoint)
	ret0, _ := ret[0].(error)
	return ret0
}

// addAppDynamicsEndpoint indicates an expected call of addAppDynamicsEndpoint
func (mr *MockStoreMockRecorder) addAppDynamicsEndpoint(chat, endpoint interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addAppDynamicsEndpoint", reflect.TypeOf((*MockStore)(nil).addAppDynamicsEndpoint), chat, endpoint)
}

// addMqEndpoint mocks base method
func (m *MockStore) addMqEndpoint(name, application, metricPath string, chat uint32, ignorePrefix []string) error {
	ret := m.ctrl.Call(m, "addMqEndpoint", name, application, metricPath, chat, ignorePrefix)
	ret0, _ := ret[0].(error)
	return ret0
}

// addMqEndpoint indicates an expected call of addMqEndpoint
func (mr *MockStoreMockRecorder) addMqEndpoint(name, application, metricPath, chat, ignorePrefix interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addMqEndpoint", reflect.TypeOf((*MockStore)(nil).addMqEndpoint), name, application, metricPath, chat, ignorePrefix)
}

// updateAppDynamics mocks base method
func (m *MockStore) updateAppDynamics(appd AppDynamics) error {
	ret := m.ctrl.Call(m, "updateAppDynamics", appd)
	ret0, _ := ret[0].(error)
	return ret0
}

// updateAppDynamics indicates an expected call of updateAppDynamics
func (mr *MockStoreMockRecorder) updateAppDynamics(appd interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "updateAppDynamics", reflect.TypeOf((*MockStore)(nil).updateAppDynamics), appd)
}

// getAllEndpoints mocks base method
func (m *MockStore) getAllEndpoints() ([]AppDynamics, error) {
	ret := m.ctrl.Call(m, "getAllEndpoints")
	ret0, _ := ret[0].([]AppDynamics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getAllEndpoints indicates an expected call of getAllEndpoints
func (mr *MockStoreMockRecorder) getAllEndpoints() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getAllEndpoints", reflect.TypeOf((*MockStore)(nil).getAllEndpoints))
}
