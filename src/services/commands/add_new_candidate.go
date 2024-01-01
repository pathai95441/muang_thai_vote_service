package commands

import (
	"context"

	candidate_domain "github.com/pathai95441/muang_thai_vote_service/src/domain/candidate"
)

type AddNewCandidateRequest struct {
	CandidateName        string `validate:"required"`
	CandidateDescription string `validate:"required"`
	CreateBy             string `validate:"required"`
}

//go:generate mockgen -source=./add_new_candidate.go -destination=./mock_commands/add_new_candidate.go -package=mock_commands
type IAddNewCandidateHandler interface {
	Handle(ctx context.Context, cmd AddNewCandidateRequest) error
}

type AddNewCandidateHandler struct {
	candidateDomain candidate_domain.ICandidateDomain
}

func NewAddNewCandidateHandler(candidateDomain candidate_domain.ICandidateDomain) IAddNewCandidateHandler {
	return AddNewCandidateHandler{candidateDomain: candidateDomain}
}

func (c AddNewCandidateHandler) Handle(ctx context.Context, cmd AddNewCandidateRequest) error {
	err := c.candidateDomain.AddNewCandidate(ctx, cmd.CandidateName, cmd.CandidateDescription, cmd.CreateBy)

	return err
}
