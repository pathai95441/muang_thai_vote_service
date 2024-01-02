// Code generated by MockGen. DO NOT EDIT.
// Source: ./vote_config_domain.go

// Package mock_vote_config_domain is a generated GoMock package.
package mock_vote_config_domain

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIVoteConfigDomain is a mock of IVoteConfigDomain interface.
type MockIVoteConfigDomain struct {
	ctrl     *gomock.Controller
	recorder *MockIVoteConfigDomainMockRecorder
}

// MockIVoteConfigDomainMockRecorder is the mock recorder for MockIVoteConfigDomain.
type MockIVoteConfigDomainMockRecorder struct {
	mock *MockIVoteConfigDomain
}

// NewMockIVoteConfigDomain creates a new mock instance.
func NewMockIVoteConfigDomain(ctrl *gomock.Controller) *MockIVoteConfigDomain {
	mock := &MockIVoteConfigDomain{ctrl: ctrl}
	mock.recorder = &MockIVoteConfigDomainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIVoteConfigDomain) EXPECT() *MockIVoteConfigDomainMockRecorder {
	return m.recorder
}

// GetClosedConfig mocks base method.
func (m *MockIVoteConfigDomain) GetClosedConfig(ctx context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClosedConfig", ctx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClosedConfig indicates an expected call of GetClosedConfig.
func (mr *MockIVoteConfigDomainMockRecorder) GetClosedConfig(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClosedConfig", reflect.TypeOf((*MockIVoteConfigDomain)(nil).GetClosedConfig), ctx)
}