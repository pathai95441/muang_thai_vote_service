// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/vote_service/vote_service_api_client.gen.go

// Package mock_vote_service is a generated GoMock package.
package mock_vote_service

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	vote_service "github.com/pathai95441/muang_thai_vote_service/pkg/vote_service"
)

// MockHttpRequestDoer is a mock of HttpRequestDoer interface.
type MockHttpRequestDoer struct {
	ctrl     *gomock.Controller
	recorder *MockHttpRequestDoerMockRecorder
}

// MockHttpRequestDoerMockRecorder is the mock recorder for MockHttpRequestDoer.
type MockHttpRequestDoerMockRecorder struct {
	mock *MockHttpRequestDoer
}

// NewMockHttpRequestDoer creates a new mock instance.
func NewMockHttpRequestDoer(ctrl *gomock.Controller) *MockHttpRequestDoer {
	mock := &MockHttpRequestDoer{ctrl: ctrl}
	mock.recorder = &MockHttpRequestDoerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttpRequestDoer) EXPECT() *MockHttpRequestDoerMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockHttpRequestDoer) Do(req *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockHttpRequestDoerMockRecorder) Do(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockHttpRequestDoer)(nil).Do), req)
}

// MockClientInterface is a mock of ClientInterface interface.
type MockClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClientInterfaceMockRecorder
}

// MockClientInterfaceMockRecorder is the mock recorder for MockClientInterface.
type MockClientInterfaceMockRecorder struct {
	mock *MockClientInterface
}

// NewMockClientInterface creates a new mock instance.
func NewMockClientInterface(ctrl *gomock.Controller) *MockClientInterface {
	mock := &MockClientInterface{ctrl: ctrl}
	mock.recorder = &MockClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientInterface) EXPECT() *MockClientInterfaceMockRecorder {
	return m.recorder
}

// GetConfig mocks base method.
func (m *MockClientInterface) GetConfig(ctx context.Context, reqEditors ...vote_service.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConfig", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockClientInterfaceMockRecorder) GetConfig(ctx interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockClientInterface)(nil).GetConfig), varargs...)
}

// MockClientWithResponsesInterface is a mock of ClientWithResponsesInterface interface.
type MockClientWithResponsesInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClientWithResponsesInterfaceMockRecorder
}

// MockClientWithResponsesInterfaceMockRecorder is the mock recorder for MockClientWithResponsesInterface.
type MockClientWithResponsesInterfaceMockRecorder struct {
	mock *MockClientWithResponsesInterface
}

// NewMockClientWithResponsesInterface creates a new mock instance.
func NewMockClientWithResponsesInterface(ctrl *gomock.Controller) *MockClientWithResponsesInterface {
	mock := &MockClientWithResponsesInterface{ctrl: ctrl}
	mock.recorder = &MockClientWithResponsesInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientWithResponsesInterface) EXPECT() *MockClientWithResponsesInterfaceMockRecorder {
	return m.recorder
}

// GetConfigWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) GetConfigWithResponse(ctx context.Context, reqEditors ...vote_service.RequestEditorFn) (*vote_service.GetConfigResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConfigWithResponse", varargs...)
	ret0, _ := ret[0].(*vote_service.GetConfigResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigWithResponse indicates an expected call of GetConfigWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) GetConfigWithResponse(ctx interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).GetConfigWithResponse), varargs...)
}
