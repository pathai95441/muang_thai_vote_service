package candidate_domain

import (
	"context"
	"errors"

	"github.com/pathai95441/muang_thai_vote_service/src/repositories/candidate"
)

//go:generate mockgen -source=./candidate_domain.go -destination=./mock/mock_candidate_domain.go -package=mock_candidate_domain
type ICandidateDomain interface {
	AddNewCandidate(ctx context.Context, candidateName string, candidateDescription string, createBy string) error
	GetAllCandidate(ctx context.Context, orderBy *string, search *string) (*[]candidate.Candidate, error)
	UpdateCandidateInfo(ctx context.Context, candidateID string, candidateName *string, candidateDescription *string, updateBy string) error
	DeleteCandidate(ctx context.Context, candidateID string, deletedBy string) error
}

type CandidateDomain struct {
	// i
	candidateRepo candidate.IRepository
}

func NewCandidateDomain(candidateRepo candidate.IRepository) CandidateDomain {
	return CandidateDomain{candidateRepo}
}

func (d CandidateDomain) AddNewCandidate(ctx context.Context, candidateName string, candidateDescription string, createBy string) error {
	newCandidate := candidate.NewCandidate(candidateName, candidateDescription)

	return d.candidateRepo.Insert(ctx, newCandidate, createBy)
}

func (d CandidateDomain) GetAllCandidate(ctx context.Context, orderBy *string, search *string) (*[]candidate.Candidate, error) {
	candidates, err := d.candidateRepo.GetAll(ctx, orderBy, search)
	if err != nil {
		return nil, err
	}

	return candidates, nil
}

func (d CandidateDomain) UpdateCandidateInfo(ctx context.Context, candidateID string, candidateName *string, candidateDescription *string, updateBy string) error {
	return d.candidateRepo.Update(ctx, nil, candidateID, candidateName, candidateDescription, nil, updateBy)
}

func (d CandidateDomain) DeleteCandidate(ctx context.Context, candidateID string, deletedBy string) error {
	candidateInfo, err := d.candidateRepo.Get(ctx, candidateID)
	if err != nil {
		return err
	}

	if candidateInfo.VoteScore != 0 {
		return errors.New("can't deleted candidate vote score is not zero")
	}

	return d.candidateRepo.Delete(ctx, nil, candidateID, deletedBy)
}
