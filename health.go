package main

import "github.com/labstack/echo"
import "net/http"

// getHealth returns an HTTP handler that provides health information
func getHealth(c echo.Context) error {
	h := &Health{"Fluffy Radio Api", "1.0.0", "Just Keep Fluffing!"}
	return c.JSON(http.StatusOK, h)
}

type (
	// Health provides basic information about the API used for health monitoring
	Health struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Message string `json:"message"`
	}
)
