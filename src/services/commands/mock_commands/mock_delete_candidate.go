// Code generated by MockGen. DO NOT EDIT.
// Source: ./delete_candidate.go

// Package mock_commands is a generated GoMock package.
package mock_commands

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	commands "github.com/pathai95441/muang_thai_vote_service/src/services/commands"
)

// MockIDeleteCandidateHandler is a mock of IDeleteCandidateHandler interface.
type MockIDeleteCandidateHandler struct {
	ctrl     *gomock.Controller
	recorder *MockIDeleteCandidateHandlerMockRecorder
}

// MockIDeleteCandidateHandlerMockRecorder is the mock recorder for MockIDeleteCandidateHandler.
type MockIDeleteCandidateHandlerMockRecorder struct {
	mock *MockIDeleteCandidateHandler
}

// NewMockIDeleteCandidateHandler creates a new mock instance.
func NewMockIDeleteCandidateHandler(ctrl *gomock.Controller) *MockIDeleteCandidateHandler {
	mock := &MockIDeleteCandidateHandler{ctrl: ctrl}
	mock.recorder = &MockIDeleteCandidateHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDeleteCandidateHandler) EXPECT() *MockIDeleteCandidateHandlerMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockIDeleteCandidateHandler) Handle(ctx context.Context, cmd commands.DeleteCandidateRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", ctx, cmd)
	ret0, _ := ret[0].(error)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockIDeleteCandidateHandlerMockRecorder) Handle(ctx, cmd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockIDeleteCandidateHandler)(nil).Handle), ctx, cmd)
}
