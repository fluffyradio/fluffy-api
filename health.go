package main

import "github.com/labstack/echo"
import "net/http"

// health returns an HTTP handler that provides health information
func health(c echo.Context) error {
	return c.JSON(http.StatusOK, []string{})
}
