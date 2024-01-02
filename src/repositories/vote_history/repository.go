package vote_history

import (
	"context"
	"database/sql"
	"time"

	"github.com/pathai95441/muang_thai_vote_service/src/config"
	"github.com/pathai95441/muang_thai_vote_service/src/consts"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/db_models_gen"
	"github.com/pathai95441/muang_thai_vote_service/src/utils/retry_utils"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

//go:generate mockgen -source=./repository.go -destination=./mock/mock_repository.go -package=mock_vote_history_repo
type IRepository interface {
	GetByUserID(ctx context.Context, userID string) (*VoteHistory, error)
	GetUserVote(ctx context.Context, userID string, candidateID string) (*VoteHistory, error)
	Insert(ctx context.Context, tx *sql.Tx, history VoteHistory) error
	CountVote(ctx context.Context, candidateID string) (*int64, error)
	DeletedVoted(ctx context.Context, tx *sql.Tx, historyID string) error
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db}
}

func (r Repository) GetByUserID(ctx context.Context, userID string) (*VoteHistory, error) {
	var (
		voteHistory *db_models_gen.VoteHistory
		err         error
	)

	err = retry_utils.RetryBackOff(config.CurrentConfig.MaxRetiresDB, func() error {
		voteHistory, err = db_models_gen.VoteHistories(
			db_models_gen.VoteHistoryWhere.UserID.EQ(userID),
			db_models_gen.VoteHistoryWhere.DeletedAt.IsNull(),
		).One(ctx, r.db)
		return err
	})

	if err != nil && err.Error() == "sql: no rows in result set" {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &VoteHistory{
		ID:          voteHistory.ID,
		CandidateID: voteHistory.CandidateID,
		UserID:      voteHistory.UserID,
	}, nil
}

func (r Repository) GetUserVote(ctx context.Context, userID string, candidateID string) (*VoteHistory, error) {
	var (
		voteHistory *db_models_gen.VoteHistory
		err         error
	)

	err = retry_utils.RetryBackOff(config.CurrentConfig.MaxRetiresDB, func() error {
		voteHistory, err = db_models_gen.VoteHistories(
			db_models_gen.VoteHistoryWhere.UserID.EQ(userID),
			db_models_gen.VoteHistoryWhere.CandidateID.EQ(candidateID),
			db_models_gen.VoteHistoryWhere.DeletedAt.IsNull(),
		).One(ctx, r.db)
		return err
	})

	if err != nil && err.Error() == "sql: no rows in result set" {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &VoteHistory{
		ID:          voteHistory.ID,
		CandidateID: voteHistory.CandidateID,
		UserID:      voteHistory.UserID,
	}, nil
}

func (r Repository) Insert(ctx context.Context, tx *sql.Tx, history VoteHistory) error {
	var executor boil.ContextExecutor
	if tx == nil {
		executor = r.db
	} else {
		executor = tx
	}
	model := db_models_gen.VoteHistory{
		ID:          history.ID,
		CandidateID: history.CandidateID,
		UserID:      history.UserID,
		CreatedBy:   consts.ServiceName,
	}

	return retry_utils.RetryBackOff(config.CurrentConfig.MaxRetiresDB, func() error {
		return model.Insert(ctx, executor, boil.Infer())
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
			db_models_gen.VoteHistoryWhere.DeletedAt.IsNull(),
		).Count(ctx, r.db)
		return err
	})

	if err != nil {
		return nil, err
	}

	return &voteScore, nil
}

func (r Repository) DeletedVoted(ctx context.Context, tx *sql.Tx, historyID string) error {
	var executor boil.ContextExecutor
	if tx == nil {
		executor = r.db
	} else {
		executor = tx
	}
	model := db_models_gen.VoteHistory{
		ID:        historyID,
		DeletedAt: null.NewTime(time.Now(), true),
		UpdatedBy: null.NewString(consts.ServiceName, true),
		DeletedBy: null.NewString(consts.ServiceName, true),
	}

	return retry_utils.RetryBackOff(config.CurrentConfig.MaxRetiresDB, func() error {
		_, updateErr := model.Update(ctx, executor, boil.Whitelist(
			db_models_gen.VoteHistoryColumns.UpdatedBy,
			db_models_gen.VoteHistoryColumns.DeletedAt,
			db_models_gen.VoteHistoryColumns.DeletedBy,
		))
		return updateErr
	})
}
