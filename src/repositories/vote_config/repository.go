package vote_config

import (
	"context"
	"database/sql"

	"github.com/pathai95441/muang_thai_vote_service/src/config"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/db_models_gen"
	"github.com/pathai95441/muang_thai_vote_service/src/utils/retry_utils"
)

//go:generate mockgen -source=./repository.go -destination=./mock/mock_repository.go -package=mock_vote_config_repo
type IRepository interface {
	Get(ctx context.Context) (bool, error)
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db}
}

func (r Repository) Get(ctx context.Context) (bool, error) {
	var (
		voteConfig *db_models_gen.VoteConfig
		err        error
	)

	err = retry_utils.RetryBackOff(config.CurrentConfig.MaxRetiresDB, func() error {
		voteConfig, err = db_models_gen.VoteConfigs(
			db_models_gen.VoteConfigWhere.DeletedAt.IsNull(),
		).One(ctx, r.db)
		return err
	})

	if err != nil {
		return false, err
	}

	return voteConfig.VoteClose != 0, nil
}
