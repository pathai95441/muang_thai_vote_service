package commands

import (
	"context"

	user_domain "github.com/pathai95441/muang_thai_vote_service/src/domain/user"
)

type CreateNewUserRequest struct {
	UserName string `validate:"required"`
	Password string `validate:"required"`
	Email    string `validate:"required"`
	RoleID   int    `validate:"required"`
}

//go:generate mockgen -source=./create_new_user.go -destination=./mock_commands/create_new_user.go -package=mock_commands
type ICreateNewUserHandler interface {
	Handle(ctx context.Context, cmd CreateNewUserRequest) error
}

type CreateNewUserHandler struct {
	userDomain user_domain.IUserDomain
}

func NewCreateNewUserHandler(userDomain user_domain.IUserDomain) ICreateNewUserHandler {
	return CreateNewUserHandler{userDomain}
}

func (c CreateNewUserHandler) Handle(ctx context.Context, cmd CreateNewUserRequest) error {
	err := c.userDomain.CreateNewUser(ctx, cmd.UserName, cmd.Password, cmd.Email, cmd.RoleID)

	return err
}
