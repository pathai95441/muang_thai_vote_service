package commands

import (
	"context"

	candidate_domain "github.com/pathai95441/muang_thai_vote_service/src/domain/candidate"
)

type DeleteCandidateRequest struct {
	CandidateID string `validate:"required"`
	DeletedBy   string `validate:"required"`
}

//go:generate mockgen -source=./delete_candidate.go -destination=./mock_commands/mock_delete_candidate.go -package=mock_commands
type IDeleteCandidateHandler interface {
	Handle(ctx context.Context, cmd DeleteCandidateRequest) error
}

type DeleteCandidateHandler struct {
	candidateDomain candidate_domain.ICandidateDomain
}

func NewDeleteCandidateHandler(candidateDomain candidate_domain.ICandidateDomain) IDeleteCandidateHandler {
	return DeleteCandidateHandler{candidateDomain}
}

func (c DeleteCandidateHandler) Handle(ctx context.Context, cmd DeleteCandidateRequest) error {
	err := c.candidateDomain.DeleteCandidate(ctx, cmd.CandidateID, cmd.DeletedBy)

	return err
}
