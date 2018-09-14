package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// SetCustomNotFoundHandler is a function to make custom handler for not found exception
func SetCustomNotFoundHandler() {
	echo.NotFoundHandler = func(c echo.Context) error {
		// render your 404 page
		return c.Render(http.StatusNotFound, "404_error", nil)
	}
}
