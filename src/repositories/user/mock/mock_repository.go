// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository.go

// Package mock_user_repo is a generated GoMock package.
package mock_user_repo

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	user "github.com/pathai95441/muang_thai_vote_service/src/repositories/user"
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
func (m *MockIRepository) Get(ctx context.Context, userID string) (*user.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, userID)
	ret0, _ := ret[0].(*user.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIRepositoryMockRecorder) Get(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIRepository)(nil).Get), ctx, userID)
}

// GetByUserName mocks base method.
func (m *MockIRepository) GetByUserName(ctx context.Context, userName string) (*user.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserName", ctx, userName)
	ret0, _ := ret[0].(*user.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserName indicates an expected call of GetByUserName.
func (mr *MockIRepositoryMockRecorder) GetByUserName(ctx, userName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserName", reflect.TypeOf((*MockIRepository)(nil).GetByUserName), ctx, userName)
}

// Insert mocks base method.
func (m *MockIRepository) Insert(ctx context.Context, userInfo user.UserInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, userInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockIRepositoryMockRecorder) Insert(ctx, userInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIRepository)(nil).Insert), ctx, userInfo)
}

// UpdateVoteCandidate mocks base method.
func (m *MockIRepository) UpdateVoteCandidate(ctx context.Context, tx *sql.Tx, userID string, candidateID *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVoteCandidate", ctx, tx, userID, candidateID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateVoteCandidate indicates an expected call of UpdateVoteCandidate.
func (mr *MockIRepositoryMockRecorder) UpdateVoteCandidate(ctx, tx, userID, candidateID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVoteCandidate", reflect.TypeOf((*MockIRepository)(nil).UpdateVoteCandidate), ctx, tx, userID, candidateID)
}
