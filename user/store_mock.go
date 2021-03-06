// Code generated by MockGen. DO NOT EDIT.
// Source: ../user/store.go

// Package user is a generated GoMock package.
package user

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

// AddUpdateUser mocks base method
func (m *MockStore) AddUpdateUser(employeeNumber, CalloutName, JiraName string) {
	m.ctrl.Call(m, "AddUpdateUser", employeeNumber, CalloutName, JiraName)
}

// AddUpdateUser indicates an expected call of AddUpdateUser
func (mr *MockStoreMockRecorder) AddUpdateUser(employeeNumber, CalloutName, JiraName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUpdateUser", reflect.TypeOf((*MockStore)(nil).AddUpdateUser), employeeNumber, CalloutName, JiraName)
}

// FindUserByCalloutName mocks base method
func (m *MockStore) FindUserByCalloutName(name string) User {
	ret := m.ctrl.Call(m, "FindUserByCalloutName", name)
	ret0, _ := ret[0].(User)
	return ret0
}

// FindUserByCalloutName indicates an expected call of FindUserByCalloutName
func (mr *MockStoreMockRecorder) FindUserByCalloutName(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByCalloutName", reflect.TypeOf((*MockStore)(nil).FindUserByCalloutName), name)
}
