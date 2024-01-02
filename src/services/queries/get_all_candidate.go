package queries

import (
	"context"

	candidate_domain "github.com/pathai95441/muang_thai_vote_service/src/domain/candidate"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/candidate"
)

//go:generate mockgen -source=./get_all_candidate.go -destination=./mock_queries/get_all_candidate.go -package=mock_queries
type IGetAllCandidateQuery interface {
	Handle(ctx context.Context, orderBy *string, search *string) (*[]candidate.Candidate, error)
}

type GetAllCandidateQuery struct {
	candidateDomain candidate_domain.ICandidateDomain
}

func NewGetAllCandidateQuery(candidateDomain candidate_domain.ICandidateDomain) IGetAllCandidateQuery {
	return GetAllCandidateQuery{candidateDomain: candidateDomain}
}

func (c GetAllCandidateQuery) Handle(ctx context.Context, orderBy *string, search *string) (*[]candidate.Candidate, error) {
	candidates, err := c.candidateDomain.GetAllCandidate(ctx, orderBy, search)
	if err != nil {
		return nil, err
	}

	return candidates, nil
}
