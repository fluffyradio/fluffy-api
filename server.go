// The Fluffy Radio API
package main

import (
	"fmt"

	"github.com/alecthomas/kingpin"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	productionMode = kingpin.Flag("prod-mode", "Run the server in production mode.").Envar("PRODUCTION_MODE").Default("false").Bool()
	port           = kingpin.Flag("port", "HTTP port for the server to run on.").Envar("PORT").Default("8080").String()
	spacialID      = kingpin.Flag("spacial-id", "The Spacial Audio station ID.").Envar("SPACIAL_ID").Required().String()
	spacialToken   = kingpin.Flag("spacial-token", "The Spacial Audio API token.").Envar("SPACIAL_TOKEN").Required().String()
)

func main() {
	fmt.Println("Initializing server...")
	e := echo.New()

	kingpin.Parse()

	if *productionMode == false {
		fmt.Println("Running in Debug Mode!")
		e.Debug = true
	}

	fmt.Println("Loading middleware...")
	registerMiddleware(e)

	fmt.Println("Registering routes...")
	registerHandlers(e)

	fmt.Println("Starting server...")
	e.Logger.Fatal(e.Start(":" + *port))
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
	e.POST("/requests/:id", requestSong)
}
