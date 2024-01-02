package commands

import (
	"context"
	"errors"

	vote_domain "github.com/pathai95441/muang_thai_vote_service/src/domain/vote"
	vote_config_domain "github.com/pathai95441/muang_thai_vote_service/src/domain/vote_config"
)

type VoteCandidateRequest struct {
	CandidateID string `validate:"required"`
	UserID      string `validate:"required"`
	UnVote      bool
}

//go:generate mockgen -source=./vote_candidate.go -destination=./mock_commands/vote_candidate.go -package=mock_commands
type IVoteCandidateHandler interface {
	Handle(ctx context.Context, cmd VoteCandidateRequest) error
}

type VoteCandidateHandler struct {
	VoteConfigDomain vote_config_domain.IVoteConfigDomain
	VoteDomain       vote_domain.IVoteDomain
}

func NewVoteCandidateHandler(VoteConfigDomain vote_config_domain.IVoteConfigDomain, VoteDomain vote_domain.IVoteDomain) IVoteCandidateHandler {
	return VoteCandidateHandler{VoteConfigDomain, VoteDomain}
}

func (c VoteCandidateHandler) Handle(ctx context.Context, cmd VoteCandidateRequest) error {
	isClosed, err := c.VoteConfigDomain.GetClosedConfig(ctx)
	if err != nil {
		return err
	}

	if isClosed {
		return errors.New("closed voted")
	}
	if cmd.UnVote {
		return c.VoteDomain.UnVoteCandidate(ctx, cmd.CandidateID, cmd.UserID)
	}

	return c.VoteDomain.VoteCandidate(ctx, cmd.CandidateID, cmd.UserID)
}
