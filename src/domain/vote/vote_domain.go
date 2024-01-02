package vote_domain

import (
	"context"
	"database/sql"
	"errors"

	"github.com/pathai95441/muang_thai_vote_service/src/consts"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/candidate"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/user"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/vote_history"
	"github.com/pathai95441/muang_thai_vote_service/src/utils/db_transaction"
)

//go:generate mockgen -source=./vote_domain.go -destination=./mock/mock_vote_domain.go -package=mock_vote_domain
type IVoteDomain interface {
	VoteCandidate(ctx context.Context, candidateID string, userID string) error
	UnVoteCandidate(ctx context.Context, candidateID string, userID string) error
}

type VoteDomain struct {
	candidateRepo   candidate.IRepository
	voteHistoryRepo vote_history.IRepository
	usersRepo       user.IRepository
	withTransaction db_transaction.IWithTxn
}

func NewVoteDomain(candidateRepo candidate.IRepository, voteHistoryRepo vote_history.IRepository, usersRepo user.IRepository, withTransaction db_transaction.IWithTxn) VoteDomain {
	return VoteDomain{candidateRepo, voteHistoryRepo, usersRepo, withTransaction}
}

func (d VoteDomain) VoteCandidate(ctx context.Context, candidateID string, userID string) error {
	// check existing vote
	voted, err := d.voteHistoryRepo.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}
	if voted != nil {
		return errors.New("already voted candidate")
	}

	return d.updateVoteScore(ctx, candidateID, userID, nil, false)
}

func (d VoteDomain) UnVoteCandidate(ctx context.Context, candidateID string, userID string) error {
	// check existing vote
	voted, err := d.voteHistoryRepo.GetUserVote(ctx, userID, candidateID)
	if err != nil {
		return err
	}

	if voted == nil {
		return errors.New("not found vote history")
	}

	return d.updateVoteScore(ctx, candidateID, userID, &voted.ID, true)
}

func (d VoteDomain) updateVoteScore(ctx context.Context, candidateID string, userID string, historyID *string, unVote bool) error {
	newVoteScore := 1
	if unVote {
		newVoteScore = -1
	}

	currentVoteScore, err := d.voteHistoryRepo.CountVote(ctx, candidateID)
	if err != nil {
		return err
	}

	return d.withTransaction.WithTransactionWithFunc(ctx, func(tx *sql.Tx) error {
		var (
			err error
		)
		if unVote {
			err = d.voteHistoryRepo.DeletedVoted(ctx, tx, *historyID)
			if err != nil {
				return err
			}

			err = d.usersRepo.UpdateVoteCandidate(ctx, tx, userID, nil)
			if err != nil {
				return err
			}
		} else {
			history := vote_history.NewVoteHistory(candidateID, userID)
			err = d.voteHistoryRepo.Insert(ctx, tx, history)
			if err != nil {
				return err
			}

			err = d.usersRepo.UpdateVoteCandidate(ctx, tx, userID, &candidateID)
			if err != nil {
				return err
			}
		}

		newVoteScore := int(*currentVoteScore) + newVoteScore
		err = d.candidateRepo.Update(ctx, tx, candidateID, nil, nil, &newVoteScore, consts.ServiceName)
		if err != nil {
			return err
		}

		return err
	})
}
