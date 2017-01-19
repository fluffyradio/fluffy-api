package main

import (
	"net/http"

	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/namsral/flag"
)

var productionMode bool

func main() {
	fmt.Println("Initializing server...")
	e := echo.New()

	flag.BoolVar(&productionMode, "productionMode", false, "False for Debug mode, otherwise True")
	flag.Parse()

	if productionMode == false {
		fmt.Println("Running in Debug Mode!")
		e.Debug = true
	}

	fmt.Println("Loading middleware...")
	registerMiddleware(e)

	fmt.Println("Registering routes...")
	registerHandlers(e)

	fmt.Println("Starting server...")
	e.Logger.Fatal(e.Start(":1323"))
}

func registerMiddleware(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())
}

func registerHandlers(e *echo.Echo) {
	// Route => handler
	e.GET("/", func(c echo.Context) error {
		h := &Health{"Fluffy Radio Api", "1.0.0", "Just Keep Fluffing!"}
		return c.JSON(http.StatusOK, h)
	})
}

type Health struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Message string `json:"message"`
}
