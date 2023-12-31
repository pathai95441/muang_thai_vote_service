// Code generated by MockGen. DO NOT EDIT.
// Source: ./user_domain.go

// Package mock_user_domain is a generated GoMock package.
package mock_user_domain

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUserDomain is a mock of IUserDomain interface.
type MockIUserDomain struct {
	ctrl     *gomock.Controller
	recorder *MockIUserDomainMockRecorder
}

// MockIUserDomainMockRecorder is the mock recorder for MockIUserDomain.
type MockIUserDomainMockRecorder struct {
	mock *MockIUserDomain
}

// NewMockIUserDomain creates a new mock instance.
func NewMockIUserDomain(ctrl *gomock.Controller) *MockIUserDomain {
	mock := &MockIUserDomain{ctrl: ctrl}
	mock.recorder = &MockIUserDomainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserDomain) EXPECT() *MockIUserDomainMockRecorder {
	return m.recorder
}

// CreateNewUser mocks base method.
func (m *MockIUserDomain) CreateNewUser(ctx context.Context, userName, password, email string, roleID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewUser", ctx, userName, password, email, roleID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNewUser indicates an expected call of CreateNewUser.
func (mr *MockIUserDomainMockRecorder) CreateNewUser(ctx, userName, password, email, roleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewUser", reflect.TypeOf((*MockIUserDomain)(nil).CreateNewUser), ctx, userName, password, email, roleID)
}
