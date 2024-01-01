package candidate_domain

import (
	"context"

	"github.com/pathai95441/muang_thai_vote_service/src/repositories/candidate"
)

//go:generate mockgen -source=./candidate_domain.go -destination=./mock/mock_candidate_domain.go -package=mock_candidate_domain
type ICandidateDomain interface {
	AddNewCandidate(ctx context.Context, candidateName string, candidateDescription string, createBy string) error
	GetAllCandidate(ctx context.Context) (*[]candidate.Candidate, error)
	UpdateCandidateInfo(ctx context.Context, candidateID string, candidateName *string, candidateDescription *string, updateBy string) error
}

type CandidateDomain struct {
	candidateRepo candidate.IRepository
}

func NewCandidateDomain(candidateRepo candidate.IRepository) CandidateDomain {
	return CandidateDomain{candidateRepo}
}

func (d CandidateDomain) AddNewCandidate(ctx context.Context, candidateName string, candidateDescription string, createBy string) error {
	newCandidate := candidate.NewCandidate(candidateName, candidateDescription)

	return d.candidateRepo.Insert(ctx, newCandidate, createBy)
}

func (d CandidateDomain) GetAllCandidate(ctx context.Context) (*[]candidate.Candidate, error) {
	candidates, err := d.candidateRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return candidates, nil
}

func (d CandidateDomain) UpdateCandidateInfo(ctx context.Context, candidateID string, candidateName *string, candidateDescription *string, updateBy string) error {
	return d.candidateRepo.Update(ctx, candidateID, candidateName, candidateDescription, nil, updateBy)
}
