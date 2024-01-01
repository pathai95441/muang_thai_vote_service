package vote_history

import (
	"context"
	"database/sql"

	"github.com/pathai95441/muang_thai_vote_service/src/config"
	"github.com/pathai95441/muang_thai_vote_service/src/consts"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/db_models_gen"
	"github.com/pathai95441/muang_thai_vote_service/src/utils/retry_utils"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

//go:generate mockgen -source=./repository.go -destination=./mock/mock_repository.go -package=mock_user
type IRepository interface {
	Insert(ctx context.Context, history VoteHistory) error
	CountVote(ctx context.Context, candidateID string) (*int64, error)
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db}
}

func (r Repository) Insert(ctx context.Context, history VoteHistory) error {
	model := db_models_gen.VoteHistory{
		ID:          history.ID,
		CandidateID: history.CandidateID,
		UserID:      history.UserID,
		CreatedBy:   consts.ServiceName,
	}

	return retry_utils.RetryBackOff(config.CurrentConfig.MaxRetiresDB, func() error {
		return model.Insert(ctx, r.db, boil.Infer())
	})
}

func (r Repository) CountVote(ctx context.Context, candidateID string) (*int64, error) {
	var (
		voteScore int64
		err       error
	)

	err = retry_utils.RetryBackOff(config.CurrentConfig.MaxRetiresDB, func() error {
		voteScore, err = db_models_gen.VoteHistories(
			db_models_gen.VoteHistoryWhere.CandidateID.EQ(candidateID),
		).Count(ctx, r.db)
		return err
	})

	if err != nil {
		return nil, err
	}

	return &voteScore, nil
}
