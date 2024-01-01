package application

import (
	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	initApplication()

	e.GET("/candidates", GetAllCandidate)
	e.POST("/candidate", CreateNewCandidate)
	e.PUT("/candidate", UpdateCandidateInfo)
}
