// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tebeka/selenium (interfaces: WebElement)

// Package mock_selenium is a generated GoMock package.
package mock_selenium

import (
	gomock "github.com/golang/mock/gomock"
	selenium "github.com/tebeka/selenium"
	reflect "reflect"
)

// MockWebElement is a mock of WebElement interface
type MockWebElement struct {
	ctrl     *gomock.Controller
	recorder *MockWebElementMockRecorder
}

// MockWebElementMockRecorder is the mock recorder for MockWebElement
type MockWebElementMockRecorder struct {
	mock *MockWebElement
}

// NewMockWebElement creates a new mock instance
func NewMockWebElement(ctrl *gomock.Controller) *MockWebElement {
	mock := &MockWebElement{ctrl: ctrl}
	mock.recorder = &MockWebElementMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWebElement) EXPECT() *MockWebElementMockRecorder {
	return m.recorder
}

// CSSProperty mocks base method
func (m *MockWebElement) CSSProperty(arg0 string) (string, error) {
	ret := m.ctrl.Call(m, "CSSProperty", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CSSProperty indicates an expected call of CSSProperty
func (mr *MockWebElementMockRecorder) CSSProperty(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CSSProperty", reflect.TypeOf((*MockWebElement)(nil).CSSProperty), arg0)
}

// Clear mocks base method
func (m *MockWebElement) Clear() error {
	ret := m.ctrl.Call(m, "Clear")
	ret0, _ := ret[0].(error)
	return ret0
}

// Clear indicates an expected call of Clear
func (mr *MockWebElementMockRecorder) Clear() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockWebElement)(nil).Clear))
}

// Click mocks base method
func (m *MockWebElement) Click() error {
	ret := m.ctrl.Call(m, "Click")
	ret0, _ := ret[0].(error)
	return ret0
}

// Click indicates an expected call of Click
func (mr *MockWebElementMockRecorder) Click() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Click", reflect.TypeOf((*MockWebElement)(nil).Click))
}

// FindElement mocks base method
func (m *MockWebElement) FindElement(arg0, arg1 string) (selenium.WebElement, error) {
	ret := m.ctrl.Call(m, "FindElement", arg0, arg1)
	ret0, _ := ret[0].(selenium.WebElement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindElement indicates an expected call of FindElement
func (mr *MockWebElementMockRecorder) FindElement(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindElement", reflect.TypeOf((*MockWebElement)(nil).FindElement), arg0, arg1)
}

// FindElements mocks base method
func (m *MockWebElement) FindElements(arg0, arg1 string) ([]selenium.WebElement, error) {
	ret := m.ctrl.Call(m, "FindElements", arg0, arg1)
	ret0, _ := ret[0].([]selenium.WebElement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindElements indicates an expected call of FindElements
func (mr *MockWebElementMockRecorder) FindElements(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindElements", reflect.TypeOf((*MockWebElement)(nil).FindElements), arg0, arg1)
}

// GetAttribute mocks base method
func (m *MockWebElement) GetAttribute(arg0 string) (string, error) {
	ret := m.ctrl.Call(m, "GetAttribute", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttribute indicates an expected call of GetAttribute
func (mr *MockWebElementMockRecorder) GetAttribute(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttribute", reflect.TypeOf((*MockWebElement)(nil).GetAttribute), arg0)
}

// IsDisplayed mocks base method
func (m *MockWebElement) IsDisplayed() (bool, error) {
	ret := m.ctrl.Call(m, "IsDisplayed")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsDisplayed indicates an expected call of IsDisplayed
func (mr *MockWebElementMockRecorder) IsDisplayed() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDisplayed", reflect.TypeOf((*MockWebElement)(nil).IsDisplayed))
}

// IsEnabled mocks base method
func (m *MockWebElement) IsEnabled() (bool, error) {
	ret := m.ctrl.Call(m, "IsEnabled")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsEnabled indicates an expected call of IsEnabled
func (mr *MockWebElementMockRecorder) IsEnabled() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEnabled", reflect.TypeOf((*MockWebElement)(nil).IsEnabled))
}

// IsSelected mocks base method
func (m *MockWebElement) IsSelected() (bool, error) {
	ret := m.ctrl.Call(m, "IsSelected")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSelected indicates an expected call of IsSelected
func (mr *MockWebElementMockRecorder) IsSelected() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSelected", reflect.TypeOf((*MockWebElement)(nil).IsSelected))
}

// Location mocks base method
func (m *MockWebElement) Location() (*selenium.Point, error) {
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(*selenium.Point)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Location indicates an expected call of Location
func (mr *MockWebElementMockRecorder) Location() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockWebElement)(nil).Location))
}

// LocationInView mocks base method
func (m *MockWebElement) LocationInView() (*selenium.Point, error) {
	ret := m.ctrl.Call(m, "LocationInView")
	ret0, _ := ret[0].(*selenium.Point)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LocationInView indicates an expected call of LocationInView
func (mr *MockWebElementMockRecorder) LocationInView() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocationInView", reflect.TypeOf((*MockWebElement)(nil).LocationInView))
}

// MoveTo mocks base method
func (m *MockWebElement) MoveTo(arg0, arg1 int) error {
	ret := m.ctrl.Call(m, "MoveTo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveTo indicates an expected call of MoveTo
func (mr *MockWebElementMockRecorder) MoveTo(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveTo", reflect.TypeOf((*MockWebElement)(nil).MoveTo), arg0, arg1)
}

// SendKeys mocks base method
func (m *MockWebElement) SendKeys(arg0 string) error {
	ret := m.ctrl.Call(m, "SendKeys", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendKeys indicates an expected call of SendKeys
func (mr *MockWebElementMockRecorder) SendKeys(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendKeys", reflect.TypeOf((*MockWebElement)(nil).SendKeys), arg0)
}

// Size mocks base method
func (m *MockWebElement) Size() (*selenium.Size, error) {
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(*selenium.Size)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Size indicates an expected call of Size
func (mr *MockWebElementMockRecorder) Size() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockWebElement)(nil).Size))
}

// Submit mocks base method
func (m *MockWebElement) Submit() error {
	ret := m.ctrl.Call(m, "Submit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Submit indicates an expected call of Submit
func (mr *MockWebElementMockRecorder) Submit() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Submit", reflect.TypeOf((*MockWebElement)(nil).Submit))
}

// TagName mocks base method
func (m *MockWebElement) TagName() (string, error) {
	ret := m.ctrl.Call(m, "TagName")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagName indicates an expected call of TagName
func (mr *MockWebElementMockRecorder) TagName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagName", reflect.TypeOf((*MockWebElement)(nil).TagName))
}

// Text mocks base method
func (m *MockWebElement) Text() (string, error) {
	ret := m.ctrl.Call(m, "Text")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Text indicates an expected call of Text
func (mr *MockWebElementMockRecorder) Text() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Text", reflect.TypeOf((*MockWebElement)(nil).Text))
}
