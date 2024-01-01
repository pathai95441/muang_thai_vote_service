// Code generated by MockGen. DO NOT EDIT.
// Source: ./create_new_user.go

// Package mock_commands is a generated GoMock package.
package mock_commands

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	commands "github.com/pathai95441/muang_thai_vote_service/src/services/commands"
)

// MockICreateNewUserHandler is a mock of ICreateNewUserHandler interface.
type MockICreateNewUserHandler struct {
	ctrl     *gomock.Controller
	recorder *MockICreateNewUserHandlerMockRecorder
}

// MockICreateNewUserHandlerMockRecorder is the mock recorder for MockICreateNewUserHandler.
type MockICreateNewUserHandlerMockRecorder struct {
	mock *MockICreateNewUserHandler
}

// NewMockICreateNewUserHandler creates a new mock instance.
func NewMockICreateNewUserHandler(ctrl *gomock.Controller) *MockICreateNewUserHandler {
	mock := &MockICreateNewUserHandler{ctrl: ctrl}
	mock.recorder = &MockICreateNewUserHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICreateNewUserHandler) EXPECT() *MockICreateNewUserHandlerMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockICreateNewUserHandler) Handle(ctx context.Context, cmd commands.CreateNewUserRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", ctx, cmd)
	ret0, _ := ret[0].(error)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockICreateNewUserHandlerMockRecorder) Handle(ctx, cmd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockICreateNewUserHandler)(nil).Handle), ctx, cmd)
}
