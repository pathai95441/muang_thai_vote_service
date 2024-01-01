package commands

import (
	"context"

	candidate_domain "github.com/pathai95441/muang_thai_vote_service/src/domain/candidate"
)

type UpdateCandidateInfoRequest struct {
	CandidateID          string `validate:"required,uuid"`
	CandidateName        *string
	CandidateDescription *string
	UpdateBy             string `validate:"required"`
}

//go:generate mockgen -source=./update_candidate_info.go -destination=./mock_commands/update_candidate_info.go -package=mock_commands
type IUpdateCandidateInfoHandler interface {
	Handle(ctx context.Context, cmd UpdateCandidateInfoRequest) error
}

type UpdateCandidateInfoHandler struct {
	candidateDomain candidate_domain.ICandidateDomain
}

func NewUpdateCandidateInfoHandler(candidateDomain candidate_domain.ICandidateDomain) IUpdateCandidateInfoHandler {
	return UpdateCandidateInfoHandler{candidateDomain: candidateDomain}
}

func (c UpdateCandidateInfoHandler) Handle(ctx context.Context, cmd UpdateCandidateInfoRequest) error {
	err := c.candidateDomain.UpdateCandidateInfo(ctx, cmd.CandidateID, cmd.CandidateName, cmd.CandidateDescription, cmd.UpdateBy)

	return err
}
