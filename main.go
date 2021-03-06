package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance is just a comment B8
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", Hello)
	e.GET("/test", Hello1)

	// Start server
	e.Logger.Fatal(e.Start(":5000"))
}

// Hello ...
func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "OK",
	})
}

// Hello1 ...
func Hello1(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "OK1111",
	})
}
