package application

import (
	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	initApplication()

	// Candidate API
	e.GET("/candidates", GetAllCandidate, authorizationAllPermission)
	e.POST("/candidate", CreateNewCandidate, authorizationAdminPermission)
	e.PUT("/candidate", UpdateCandidateInfo, authorizationAdminPermission)
	e.DELETE("/candidate/:candidateID", DeleteCandidateByID, authorizationAdminPermission)

	// User API
	e.POST("/user", CreateNewUser)

	// Auth
	e.POST("/sign_in", SignIn)

	// Vote APi
	e.POST("/vote_candidate", VoteCandidate, authorizationAllPermission)
}
