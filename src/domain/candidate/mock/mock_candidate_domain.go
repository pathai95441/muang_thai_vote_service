// Code generated by MockGen. DO NOT EDIT.
// Source: ./candidate_domain.go

// Package mock_candidate_domain is a generated GoMock package.
package mock_candidate_domain

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	candidate "github.com/pathai95441/muang_thai_vote_service/src/repositories/candidate"
)

// MockICandidateDomain is a mock of ICandidateDomain interface.
type MockICandidateDomain struct {
	ctrl     *gomock.Controller
	recorder *MockICandidateDomainMockRecorder
}

// MockICandidateDomainMockRecorder is the mock recorder for MockICandidateDomain.
type MockICandidateDomainMockRecorder struct {
	mock *MockICandidateDomain
}

// NewMockICandidateDomain creates a new mock instance.
func NewMockICandidateDomain(ctrl *gomock.Controller) *MockICandidateDomain {
	mock := &MockICandidateDomain{ctrl: ctrl}
	mock.recorder = &MockICandidateDomainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICandidateDomain) EXPECT() *MockICandidateDomainMockRecorder {
	return m.recorder
}

// AddNewCandidate mocks base method.
func (m *MockICandidateDomain) AddNewCandidate(ctx context.Context, candidateName, candidateDescription, createBy string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNewCandidate", ctx, candidateName, candidateDescription, createBy)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNewCandidate indicates an expected call of AddNewCandidate.
func (mr *MockICandidateDomainMockRecorder) AddNewCandidate(ctx, candidateName, candidateDescription, createBy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewCandidate", reflect.TypeOf((*MockICandidateDomain)(nil).AddNewCandidate), ctx, candidateName, candidateDescription, createBy)
}

// DeleteCandidate mocks base method.
func (m *MockICandidateDomain) DeleteCandidate(ctx context.Context, candidateID, deletedBy string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCandidate", ctx, candidateID, deletedBy)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCandidate indicates an expected call of DeleteCandidate.
func (mr *MockICandidateDomainMockRecorder) DeleteCandidate(ctx, candidateID, deletedBy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCandidate", reflect.TypeOf((*MockICandidateDomain)(nil).DeleteCandidate), ctx, candidateID, deletedBy)
}

// GetAllCandidate mocks base method.
func (m *MockICandidateDomain) GetAllCandidate(ctx context.Context) (*[]candidate.Candidate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCandidate", ctx)
	ret0, _ := ret[0].(*[]candidate.Candidate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCandidate indicates an expected call of GetAllCandidate.
func (mr *MockICandidateDomainMockRecorder) GetAllCandidate(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCandidate", reflect.TypeOf((*MockICandidateDomain)(nil).GetAllCandidate), ctx)
}

// UpdateCandidateInfo mocks base method.
func (m *MockICandidateDomain) UpdateCandidateInfo(ctx context.Context, candidateID string, candidateName, candidateDescription *string, updateBy string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCandidateInfo", ctx, candidateID, candidateName, candidateDescription, updateBy)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCandidateInfo indicates an expected call of UpdateCandidateInfo.
func (mr *MockICandidateDomainMockRecorder) UpdateCandidateInfo(ctx, candidateID, candidateName, candidateDescription, updateBy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCandidateInfo", reflect.TypeOf((*MockICandidateDomain)(nil).UpdateCandidateInfo), ctx, candidateID, candidateName, candidateDescription, updateBy)
}
