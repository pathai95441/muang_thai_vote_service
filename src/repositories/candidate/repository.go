package candidate

import (
	"context"
	"database/sql"
	"time"

	"github.com/pathai95441/muang_thai_vote_service/src/config"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/db_models_gen"
	"github.com/pathai95441/muang_thai_vote_service/src/utils/retry_utils"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

//go:generate mockgen -source=./repository.go -destination=./mock/mock_repository.go -package=mock_candidate
type IRepository interface {
	Get(ctx context.Context, candidateID string) (*Candidate, error)
	GetAll(ctx context.Context, orderBy *string, search *string) (*[]Candidate, error)
	Insert(ctx context.Context, candidate Candidate, createBy string) error
	Update(ctx context.Context, tx *sql.Tx, candidateID string, candidateName *string, candidateDescription *string, voteScore *int, updateBy string) error
	Delete(ctx context.Context, tx *sql.Tx, candidateID string, deleteBy string) error
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db}
}

func (r Repository) Get(ctx context.Context, candidateID string) (*Candidate, error) {
	var (
		candidateInfo *db_models_gen.Candidate
		err           error
	)

	err = retry_utils.RetryBackOff(config.CurrentConfig.MaxRetiresDB, func() error {
		candidateInfo, err = db_models_gen.Candidates(
			db_models_gen.CandidateWhere.ID.EQ(candidateID),
			db_models_gen.CandidateWhere.DeletedAt.IsNull(),
		).One(ctx, r.db)
		return err
	})

	if err != nil {
		return nil, err
	}

	return &Candidate{
		ID:                   candidateInfo.ID,
		CandidateName:        candidateInfo.CandidateName,
		CandidateDescription: candidateInfo.CandidateDescription,
		VoteScore:            int64(candidateInfo.VoteScore),
	}, nil
}

func (r Repository) GetAll(ctx context.Context, orderBy *string, search *string) (*[]Candidate, error) {
	var (
		candidates        db_models_gen.CandidateSlice
		mapCandidatesData []Candidate
		err               error
	)

	var conditions []qm.QueryMod
	conditions = append(conditions,
		db_models_gen.CandidateWhere.DeletedAt.IsNull(),
	)
	if orderBy != nil {
		conditions = append(conditions,
			qm.OrderBy(db_models_gen.CandidateColumns.VoteScore+" "+*orderBy),
		)

	}
	if search != nil {
		likeSubstring := "%" + *search + "%"
		conditions = append(conditions,
			qm.Where(db_models_gen.CandidateColumns.CandidateName+" LIKE ?", likeSubstring),
		)
	}

	err = retry_utils.RetryBackOff(config.CurrentConfig.MaxRetiresDB, func() error {
		candidates, err = db_models_gen.Candidates(
			conditions...,
		).All(ctx, r.db)
		return err
	})

	if err != nil {
		return nil, err
	}

	for _, candidate := range candidates {
		mapCandidatesData = append(mapCandidatesData, UnmarshallCandidateFromDB(candidate.ID, candidate.CandidateName, candidate.CandidateDescription, int64(candidate.VoteScore)))
	}

	return &mapCandidatesData, nil
}

func (r Repository) Insert(ctx context.Context, candidate Candidate, createBy string) error {
	model := db_models_gen.Candidate{
		ID:                   candidate.ID,
		CandidateName:        candidate.CandidateName,
		CandidateDescription: candidate.CandidateDescription,
		VoteScore:            0,
		CreatedBy:            createBy,
	}

	return retry_utils.RetryBackOff(config.CurrentConfig.MaxRetiresDB, func() error {
		return model.Insert(ctx, r.db, boil.Infer())
	})
}

func (r Repository) Update(ctx context.Context, tx *sql.Tx, candidateID string, candidateName *string, candidateDescription *string, voteScore *int, updateBy string) error {
	var executor boil.ContextExecutor
	var updateColumn []string
	if tx == nil {
		executor = r.db
	} else {
		executor = tx
	}
	model := db_models_gen.Candidate{
		ID:        candidateID,
		UpdatedBy: null.StringFrom(updateBy),
	}
	updateColumn = append(updateColumn, db_models_gen.CandidateColumns.UpdatedBy)

	if candidateName != nil {
		model.CandidateName = *candidateName
		updateColumn = append(updateColumn, db_models_gen.CandidateColumns.CandidateName)
	}

	if candidateDescription != nil {
		model.CandidateDescription = *candidateDescription
		updateColumn = append(updateColumn, db_models_gen.CandidateColumns.CandidateDescription)
	}

	if voteScore != nil {
		model.VoteScore = *voteScore
		updateColumn = append(updateColumn, db_models_gen.CandidateColumns.VoteScore)
	}

	return retry_utils.RetryBackOff(config.CurrentConfig.MaxRetiresDB, func() error {
		_, updateErr := model.Update(ctx, executor, boil.Whitelist(
			updateColumn...,
		))
		return updateErr
	})
}

func (r Repository) Delete(ctx context.Context, tx *sql.Tx, candidateID string, deleteBy string) error {
	var executor boil.ContextExecutor
	if tx == nil {
		executor = r.db
	} else {
		executor = tx
	}
	model := db_models_gen.Candidate{
		ID:        candidateID,
		DeletedAt: null.NewTime(time.Now(), true),
		DeletedBy: null.StringFrom(deleteBy),
	}

	return retry_utils.RetryBackOff(config.CurrentConfig.MaxRetiresDB, func() error {
		_, updateErr := model.Update(ctx, executor, boil.Whitelist(
			db_models_gen.CandidateColumns.DeletedAt,
			db_models_gen.CandidateColumns.DeletedBy,
		))
		return updateErr
	})
}
