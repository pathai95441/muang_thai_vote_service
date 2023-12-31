// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository.go

// Package mock_vote_config_repo is a generated GoMock package.
package mock_vote_config_repo

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIRepository is a mock of IRepository interface.
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository.
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance.
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockIRepository) Get(ctx context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIRepositoryMockRecorder) Get(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIRepository)(nil).Get), ctx)
}
