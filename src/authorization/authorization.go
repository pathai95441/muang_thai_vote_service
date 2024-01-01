package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pathai95441/muang_thai_vote_service/src/config"
	"github.com/pathai95441/muang_thai_vote_service/src/repositories/user"
	"github.com/pathai95441/muang_thai_vote_service/src/utils/contains"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	UserID string `json:"username"`
	jwt.StandardClaims
}

//go:generate mockgen -source=./authorization.go -destination=./mock/authorization.go -package=mock_auth
type IAuthHandler interface {
	Authorization(ctx context.Context, tokenString string, Permission []int) error
	SignIn(ctx context.Context, userName string, password string) (*string, error)
}

type AuthHandler struct {
	userRepo user.IRepository
}

func NewAuthHandler(userRepo user.IRepository) IAuthHandler {
	return AuthHandler{userRepo}
}

func (c AuthHandler) Authorization(ctx context.Context, tokenString string, permissions []int) error {
	tokenData, err := validateToken(tokenString)
	if err != nil {
		return err
	}

	userInfo, err := c.userRepo.Get(ctx, tokenData.UserID)
	if err != nil {
		return err
	}
	if !contains.ContainsElement(permissions, userInfo.RoleID) {
		return errors.New("not has permission")
	}

	return nil
}

func (c AuthHandler) SignIn(ctx context.Context, userName string, password string) (*string, error) {
	userInfo, err := c.userRepo.GetByUserName(ctx, userName)
	if err != nil {
		return nil, err
	}

	passwordIsMatch := c.checkPassword(password, userInfo.Password)
	if !passwordIsMatch {
		return nil, fmt.Errorf("invalid password")
	}

	claims := Token{
		UserID: userInfo.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.CurrentConfig.SecretKey)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func (c AuthHandler) checkPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func validateToken(tokenString string) (*Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return config.CurrentConfig.SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Token); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Token is invalid")
}
