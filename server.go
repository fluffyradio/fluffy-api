// The Fluffy Radio API
package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/namsral/flag"
)

var productionMode bool
var spacialToken string

func main() {
	fmt.Println("Initializing server...")
	e := echo.New()

	flag.BoolVar(&productionMode, "productionMode", false, "False for Debug mode, otherwise True")
	flag.StringVar(&spacialToken, "spacialToken", "", "Authentication token for Spacial Audio")
	flag.Parse()

	if productionMode == false {
		fmt.Println("Running in Debug Mode!")
		e.Debug = true
	}

	if spacialToken == "" {
		e.Logger.Panic("Spacial Token not provided, exiting.")
	}

	fmt.Println("Loading middleware...")
	registerMiddleware(e)

	fmt.Println("Registering routes...")
	registerHandlers(e)

	fmt.Println("Starting server...")
	e.Logger.Fatal(e.Start(":8080"))
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
	e.GET("/", health)
	e.GET("/songs", songs)
	e.GET("/songs/current", currentSong)
}
