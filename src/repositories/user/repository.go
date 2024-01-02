package user

import (
	"context"
	"database/sql"

	"github.com/pathai95441/muang_thai_vote_service/src/config"
	"github.com/pathai95441/muang_thai_vote_service/src/consts"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/db_models_gen"
	"github.com/pathai95441/muang_thai_vote_service/src/utils/retry_utils"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

//go:generate mockgen -source=./repository.go -destination=./mock/mock_repository.go -package=mock_user_repo
type IRepository interface {
	Get(ctx context.Context, userID string) (*UserInfo, error)
	GetByUserName(ctx context.Context, userName string) (*UserInfo, error)
	Insert(ctx context.Context, userInfo UserInfo) error
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db}
}

func (r Repository) Get(ctx context.Context, userID string) (*UserInfo, error) {
	var (
		userInfo *db_models_gen.User
		err      error
	)

	err = retry_utils.RetryBackOff(config.CurrentConfig.MaxRetiresDB, func() error {
		userInfo, err = db_models_gen.Users(
			db_models_gen.UserWhere.ID.EQ(userID),
			db_models_gen.UserWhere.DeletedAt.IsNull(),
		).One(ctx, r.db)
		return err
	})

	if err != nil {
		return nil, err
	}

	return &UserInfo{
		ID:       userInfo.ID,
		UserName: userInfo.UserName,
		Password: userInfo.Password,
		Email:    userInfo.Email,
		RoleID:   userInfo.RoleID,
	}, nil
}

func (r Repository) GetByUserName(ctx context.Context, userName string) (*UserInfo, error) {
	var (
		userInfo *db_models_gen.User
		err      error
	)

	err = retry_utils.RetryBackOff(config.CurrentConfig.MaxRetiresDB, func() error {
		userInfo, err = db_models_gen.Users(
			db_models_gen.UserWhere.UserName.EQ(userName),
			db_models_gen.UserWhere.DeletedAt.IsNull(),
		).One(ctx, r.db)
		return err
	})

	if err != nil {
		return nil, err
	}

	return &UserInfo{
		ID:       userInfo.ID,
		UserName: userInfo.UserName,
		Password: userInfo.Password,
		Email:    userInfo.Email,
		RoleID:   userInfo.RoleID,
	}, nil
}

func (r Repository) Insert(ctx context.Context, userInfo UserInfo) error {
	model := db_models_gen.User{
		ID:        userInfo.ID,
		UserName:  userInfo.UserName,
		Email:     userInfo.Email,
		Password:  userInfo.Password,
		RoleID:    userInfo.RoleID,
		CreatedBy: consts.ServiceName,
	}

	return retry_utils.RetryBackOff(config.CurrentConfig.MaxRetiresDB, func() error {
		return model.Insert(ctx, r.db, boil.Infer())
	})
}
