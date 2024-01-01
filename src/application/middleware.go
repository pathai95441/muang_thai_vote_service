package application

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pathai95441/muang_thai_vote_service/src/consts"
)

var (
	allPermission = []int{consts.Admin, consts.Guest}
	adminOnly     = []int{consts.Admin}
)

func authorizationAllPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := getTokenFromHeader(c)
		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}
		err = serverApp.Authorization.Authorization(context.Background(), *token, allPermission)
		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}

		err = next(c)
		return err
	}
}

func authorizationAdminPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := getTokenFromHeader(c)
		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}
		err = serverApp.Authorization.Authorization(context.Background(), *token, adminOnly)
		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}
		err = next(c)
		return err
	}
}

func getTokenFromHeader(c echo.Context) (*string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("authorization header is missing")
	}
	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		return nil, errors.New("invalid authorization header format")
	}

	token := authParts[1]
	return &token, nil
}
