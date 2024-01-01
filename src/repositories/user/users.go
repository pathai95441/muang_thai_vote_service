package user

import "github.com/google/uuid"

type UserInfo struct {
	ID       string
	UserName string
	Password string
	Email    string
	RoleID   int
}

func NewUserInfo(
	UserName string,
	Password string,
	Email string,
	RoleID int,
) UserInfo {
	return UserInfo{
		ID:       uuid.NewString(),
		UserName: UserName,
		Password: Password,
		Email:    Email,
		RoleID:   RoleID,
	}
}

func UnmarshallUserInfoFromDB(
	ID string,
	UserName string,
	Password string,
	Email string,
	RoleID int,
) UserInfo {
	return UserInfo{
		ID,
		UserName,
		Password,
		Email,
		RoleID,
	}
}
