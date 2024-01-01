package candidate

import (
	"context"
	"database/sql"

	"github.com/pathai95441/muang_thai_vote_service/src/config"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/db_models_gen"
	"github.com/pathai95441/muang_thai_vote_service/src/utils/retry_utils"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

//go:generate mockgen -source=./repository.go -destination=./mock/mock_repository.go -package=mock_candidate
type IRepository interface {
	GetAll(ctx context.Context) (*[]Candidate, error)
	Insert(ctx context.Context, candidate Candidate, createBy string) error
	Update(ctx context.Context, candidateID string, candidateName *string, candidateDescription *string, voteScore *int, updateBy string) error
	// Deleted(ctx context.Context, tx *sql.Tx, referenceID string, status int8) error
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db}
}

func (r Repository) GetAll(ctx context.Context) (*[]Candidate, error) {
	var (
		candidates        db_models_gen.CandidateSlice
		mapCandidatesData []Candidate
		err               error
	)

	err = retry_utils.RetryBackOff(config.CurrentConfig.MaxRetiresDB, func() error {
		candidates, err = db_models_gen.Candidates(
			db_models_gen.CandidateWhere.DeletedAt.IsNull(),
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

func (r Repository) Update(ctx context.Context, candidateID string, candidateName *string, candidateDescription *string, voteScore *int, updateBy string) error {
	model := db_models_gen.Candidate{
		ID:        candidateID,
		UpdatedBy: null.StringFrom(updateBy),
	}

	if candidateName != nil {
		model.CandidateName = *candidateName
	}

	if candidateDescription != nil {
		model.CandidateDescription = *candidateDescription
	}

	if voteScore != nil {
		model.VoteScore = *voteScore
	}

	return retry_utils.RetryBackOff(config.CurrentConfig.MaxRetiresDB, func() error {
		_, updateErr := model.Update(ctx, r.db, boil.Whitelist(
			db_models_gen.CandidateColumns.CandidateName,
			db_models_gen.CandidateColumns.CandidateDescription,
			db_models_gen.CandidateColumns.VoteScore,
			db_models_gen.CandidateColumns.UpdatedBy,
		))
		return updateErr
	})
}
