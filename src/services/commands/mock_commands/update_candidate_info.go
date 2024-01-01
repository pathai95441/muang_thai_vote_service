// Code generated by MockGen. DO NOT EDIT.
// Source: ./update_candidate_info.go

// Package mock_commands is a generated GoMock package.
package mock_commands

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	commands "github.com/pathai95441/muang_thai_vote_service/src/services/commands"
)

// MockIUpdateCandidateInfoHandler is a mock of IUpdateCandidateInfoHandler interface.
type MockIUpdateCandidateInfoHandler struct {
	ctrl     *gomock.Controller
	recorder *MockIUpdateCandidateInfoHandlerMockRecorder
}

// MockIUpdateCandidateInfoHandlerMockRecorder is the mock recorder for MockIUpdateCandidateInfoHandler.
type MockIUpdateCandidateInfoHandlerMockRecorder struct {
	mock *MockIUpdateCandidateInfoHandler
}

// NewMockIUpdateCandidateInfoHandler creates a new mock instance.
func NewMockIUpdateCandidateInfoHandler(ctrl *gomock.Controller) *MockIUpdateCandidateInfoHandler {
	mock := &MockIUpdateCandidateInfoHandler{ctrl: ctrl}
	mock.recorder = &MockIUpdateCandidateInfoHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUpdateCandidateInfoHandler) EXPECT() *MockIUpdateCandidateInfoHandlerMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockIUpdateCandidateInfoHandler) Handle(ctx context.Context, cmd commands.UpdateCandidateInfoRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", ctx, cmd)
	ret0, _ := ret[0].(error)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockIUpdateCandidateInfoHandlerMockRecorder) Handle(ctx, cmd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockIUpdateCandidateInfoHandler)(nil).Handle), ctx, cmd)
}
