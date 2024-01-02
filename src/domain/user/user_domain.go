package user_domain

import (
	"context"

	"github.com/pathai95441/muang_thai_vote_service/src/repositories/user"
	"golang.org/x/crypto/bcrypt"
)

//go:generate mockgen -source=./user_domain.go -destination=./mock/mock_user_domain.go -package=mock_user_domain
type IUserDomain interface {
	CreateNewUser(ctx context.Context, userName string, password string, email string, roleID int) error
}

type UserDomain struct {
	userRepo user.IRepository
}

func NewUserDomain(userRepo user.IRepository) UserDomain {
	return UserDomain{userRepo}
}

func (d UserDomain) CreateNewUser(ctx context.Context, userName string, password string, email string, roleID int) error {
	hash, err := hashPassword(password)
	if err != nil {
		return err
	}
	newUser := user.NewUserInfo(userName, hash, email, roleID, nil)

	return d.userRepo.Insert(ctx, newUser)
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
