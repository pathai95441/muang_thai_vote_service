package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pathai95441/muang_thai_vote_service/src/application"
)

func main() {
	// Echo instance
	e := echo.New()

	// Use the CORS middleware with Echo
	e.Use(middleware.CORS())

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	application.InitRouter(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
