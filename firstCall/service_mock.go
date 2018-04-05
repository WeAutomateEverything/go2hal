// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_firstCall is a generated GoMock package.
package firstCall

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetFirstCall mocks base method
func (m *MockService) GetFirstCall(ctx context.Context) (string, string, error) {
	ret := m.ctrl.Call(m, "GetFirstCall", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetFirstCall indicates an expected call of GetFirstCall
func (mr *MockServiceMockRecorder) GetFirstCall(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFirstCall", reflect.TypeOf((*MockService)(nil).GetFirstCall), ctx)
}
