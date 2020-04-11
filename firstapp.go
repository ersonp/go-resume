// main.go

package main

import (
	"github.com/ersonp/firstapp/server"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	server.Router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	server.Router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	server.InitializeRoutes()

	// Start serving the application
	server.Router.Run()
}
