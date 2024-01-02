// Code generated by MockGen. DO NOT EDIT.
// Source: ./authorization.go

// Package mock_auth is a generated GoMock package.
package mock_auth

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	user "github.com/pathai95441/muang_thai_vote_service/src/repositories/user"
)

// MockIAuthHandler is a mock of IAuthHandler interface.
type MockIAuthHandler struct {
	ctrl     *gomock.Controller
	recorder *MockIAuthHandlerMockRecorder
}

// MockIAuthHandlerMockRecorder is the mock recorder for MockIAuthHandler.
type MockIAuthHandlerMockRecorder struct {
	mock *MockIAuthHandler
}

// NewMockIAuthHandler creates a new mock instance.
func NewMockIAuthHandler(ctrl *gomock.Controller) *MockIAuthHandler {
	mock := &MockIAuthHandler{ctrl: ctrl}
	mock.recorder = &MockIAuthHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAuthHandler) EXPECT() *MockIAuthHandlerMockRecorder {
	return m.recorder
}

// Authorization mocks base method.
func (m *MockIAuthHandler) Authorization(ctx context.Context, tokenString string, Permission []int) (*user.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorization", ctx, tokenString, Permission)
	ret0, _ := ret[0].(*user.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authorization indicates an expected call of Authorization.
func (mr *MockIAuthHandlerMockRecorder) Authorization(ctx, tokenString, Permission interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorization", reflect.TypeOf((*MockIAuthHandler)(nil).Authorization), ctx, tokenString, Permission)
}

// SignIn mocks base method.
func (m *MockIAuthHandler) SignIn(ctx context.Context, userName, password string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", ctx, userName, password)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockIAuthHandlerMockRecorder) SignIn(ctx, userName, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockIAuthHandler)(nil).SignIn), ctx, userName, password)
}
