// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/vote_service/vote_service_api_client.gen.go

// Package mock_vote_service is a generated GoMock package.
package mock_vote_service

import (
	context "context"
	io "io"
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

// CreateNewCandidate mocks base method.
func (m *MockClientInterface) CreateNewCandidate(ctx context.Context, body vote_service.CreateNewCandidateJSONRequestBody, reqEditors ...vote_service.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateNewCandidate", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewCandidate indicates an expected call of CreateNewCandidate.
func (mr *MockClientInterfaceMockRecorder) CreateNewCandidate(ctx, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewCandidate", reflect.TypeOf((*MockClientInterface)(nil).CreateNewCandidate), varargs...)
}

// CreateNewCandidateWithBody mocks base method.
func (m *MockClientInterface) CreateNewCandidateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...vote_service.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, contentType, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateNewCandidateWithBody", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewCandidateWithBody indicates an expected call of CreateNewCandidateWithBody.
func (mr *MockClientInterfaceMockRecorder) CreateNewCandidateWithBody(ctx, contentType, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, contentType, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewCandidateWithBody", reflect.TypeOf((*MockClientInterface)(nil).CreateNewCandidateWithBody), varargs...)
}

// CreateNewUser mocks base method.
func (m *MockClientInterface) CreateNewUser(ctx context.Context, body vote_service.CreateNewUserJSONRequestBody, reqEditors ...vote_service.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateNewUser", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewUser indicates an expected call of CreateNewUser.
func (mr *MockClientInterfaceMockRecorder) CreateNewUser(ctx, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewUser", reflect.TypeOf((*MockClientInterface)(nil).CreateNewUser), varargs...)
}

// CreateNewUserWithBody mocks base method.
func (m *MockClientInterface) CreateNewUserWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...vote_service.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, contentType, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateNewUserWithBody", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewUserWithBody indicates an expected call of CreateNewUserWithBody.
func (mr *MockClientInterfaceMockRecorder) CreateNewUserWithBody(ctx, contentType, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, contentType, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewUserWithBody", reflect.TypeOf((*MockClientInterface)(nil).CreateNewUserWithBody), varargs...)
}

// GetAllCandidate mocks base method.
func (m *MockClientInterface) GetAllCandidate(ctx context.Context, reqEditors ...vote_service.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllCandidate", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCandidate indicates an expected call of GetAllCandidate.
func (mr *MockClientInterfaceMockRecorder) GetAllCandidate(ctx interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCandidate", reflect.TypeOf((*MockClientInterface)(nil).GetAllCandidate), varargs...)
}

// SignIn mocks base method.
func (m *MockClientInterface) SignIn(ctx context.Context, body vote_service.SignInJSONRequestBody, reqEditors ...vote_service.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SignIn", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockClientInterfaceMockRecorder) SignIn(ctx, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockClientInterface)(nil).SignIn), varargs...)
}

// SignInWithBody mocks base method.
func (m *MockClientInterface) SignInWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...vote_service.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, contentType, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SignInWithBody", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignInWithBody indicates an expected call of SignInWithBody.
func (mr *MockClientInterfaceMockRecorder) SignInWithBody(ctx, contentType, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, contentType, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignInWithBody", reflect.TypeOf((*MockClientInterface)(nil).SignInWithBody), varargs...)
}

// UpdateCandidateInfo mocks base method.
func (m *MockClientInterface) UpdateCandidateInfo(ctx context.Context, body vote_service.UpdateCandidateInfoJSONRequestBody, reqEditors ...vote_service.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCandidateInfo", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCandidateInfo indicates an expected call of UpdateCandidateInfo.
func (mr *MockClientInterfaceMockRecorder) UpdateCandidateInfo(ctx, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCandidateInfo", reflect.TypeOf((*MockClientInterface)(nil).UpdateCandidateInfo), varargs...)
}

// UpdateCandidateInfoWithBody mocks base method.
func (m *MockClientInterface) UpdateCandidateInfoWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...vote_service.RequestEditorFn) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, contentType, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCandidateInfoWithBody", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCandidateInfoWithBody indicates an expected call of UpdateCandidateInfoWithBody.
func (mr *MockClientInterfaceMockRecorder) UpdateCandidateInfoWithBody(ctx, contentType, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, contentType, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCandidateInfoWithBody", reflect.TypeOf((*MockClientInterface)(nil).UpdateCandidateInfoWithBody), varargs...)
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

// CreateNewCandidateWithBodyWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) CreateNewCandidateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...vote_service.RequestEditorFn) (*vote_service.CreateNewCandidateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, contentType, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateNewCandidateWithBodyWithResponse", varargs...)
	ret0, _ := ret[0].(*vote_service.CreateNewCandidateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewCandidateWithBodyWithResponse indicates an expected call of CreateNewCandidateWithBodyWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) CreateNewCandidateWithBodyWithResponse(ctx, contentType, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, contentType, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewCandidateWithBodyWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).CreateNewCandidateWithBodyWithResponse), varargs...)
}

// CreateNewCandidateWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) CreateNewCandidateWithResponse(ctx context.Context, body vote_service.CreateNewCandidateJSONRequestBody, reqEditors ...vote_service.RequestEditorFn) (*vote_service.CreateNewCandidateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateNewCandidateWithResponse", varargs...)
	ret0, _ := ret[0].(*vote_service.CreateNewCandidateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewCandidateWithResponse indicates an expected call of CreateNewCandidateWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) CreateNewCandidateWithResponse(ctx, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewCandidateWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).CreateNewCandidateWithResponse), varargs...)
}

// CreateNewUserWithBodyWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) CreateNewUserWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...vote_service.RequestEditorFn) (*vote_service.CreateNewUserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, contentType, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateNewUserWithBodyWithResponse", varargs...)
	ret0, _ := ret[0].(*vote_service.CreateNewUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewUserWithBodyWithResponse indicates an expected call of CreateNewUserWithBodyWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) CreateNewUserWithBodyWithResponse(ctx, contentType, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, contentType, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewUserWithBodyWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).CreateNewUserWithBodyWithResponse), varargs...)
}

// CreateNewUserWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) CreateNewUserWithResponse(ctx context.Context, body vote_service.CreateNewUserJSONRequestBody, reqEditors ...vote_service.RequestEditorFn) (*vote_service.CreateNewUserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateNewUserWithResponse", varargs...)
	ret0, _ := ret[0].(*vote_service.CreateNewUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewUserWithResponse indicates an expected call of CreateNewUserWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) CreateNewUserWithResponse(ctx, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewUserWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).CreateNewUserWithResponse), varargs...)
}

// GetAllCandidateWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) GetAllCandidateWithResponse(ctx context.Context, reqEditors ...vote_service.RequestEditorFn) (*vote_service.GetAllCandidateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllCandidateWithResponse", varargs...)
	ret0, _ := ret[0].(*vote_service.GetAllCandidateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCandidateWithResponse indicates an expected call of GetAllCandidateWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) GetAllCandidateWithResponse(ctx interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCandidateWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).GetAllCandidateWithResponse), varargs...)
}

// SignInWithBodyWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) SignInWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...vote_service.RequestEditorFn) (*vote_service.SignInResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, contentType, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SignInWithBodyWithResponse", varargs...)
	ret0, _ := ret[0].(*vote_service.SignInResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignInWithBodyWithResponse indicates an expected call of SignInWithBodyWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) SignInWithBodyWithResponse(ctx, contentType, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, contentType, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignInWithBodyWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).SignInWithBodyWithResponse), varargs...)
}

// SignInWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) SignInWithResponse(ctx context.Context, body vote_service.SignInJSONRequestBody, reqEditors ...vote_service.RequestEditorFn) (*vote_service.SignInResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SignInWithResponse", varargs...)
	ret0, _ := ret[0].(*vote_service.SignInResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignInWithResponse indicates an expected call of SignInWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) SignInWithResponse(ctx, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignInWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).SignInWithResponse), varargs...)
}

// UpdateCandidateInfoWithBodyWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) UpdateCandidateInfoWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...vote_service.RequestEditorFn) (*vote_service.UpdateCandidateInfoResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, contentType, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCandidateInfoWithBodyWithResponse", varargs...)
	ret0, _ := ret[0].(*vote_service.UpdateCandidateInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCandidateInfoWithBodyWithResponse indicates an expected call of UpdateCandidateInfoWithBodyWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) UpdateCandidateInfoWithBodyWithResponse(ctx, contentType, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, contentType, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCandidateInfoWithBodyWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).UpdateCandidateInfoWithBodyWithResponse), varargs...)
}

// UpdateCandidateInfoWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) UpdateCandidateInfoWithResponse(ctx context.Context, body vote_service.UpdateCandidateInfoJSONRequestBody, reqEditors ...vote_service.RequestEditorFn) (*vote_service.UpdateCandidateInfoResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, body}
	for _, a := range reqEditors {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCandidateInfoWithResponse", varargs...)
	ret0, _ := ret[0].(*vote_service.UpdateCandidateInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCandidateInfoWithResponse indicates an expected call of UpdateCandidateInfoWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) UpdateCandidateInfoWithResponse(ctx, body interface{}, reqEditors ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, body}, reqEditors...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCandidateInfoWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).UpdateCandidateInfoWithResponse), varargs...)
}
