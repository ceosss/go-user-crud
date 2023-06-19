package router

import (
	route "github.com/ceosss/go-user-crud/src/routes"
	"github.com/gin-gonic/gin"
)

// This function initializes a new router in Go using the Gin framework and sets up logging, error
// recovery, and routes.
func InitializeRouter() {
	// `r := gin.New()` is initializing a new instance of the Gin router. This creates a new router that
	// can be used to define routes and middleware for the application.
	r := gin.New()

	// `r.Use(gin.Logger())` is setting up logging middleware for the Gin router. This middleware logs
	// information about incoming requests and outgoing responses, including the HTTP method, URL, status
	// code, and response time. This can be useful for debugging and monitoring purposes.
	r.Use(gin.Logger())

	// `r.Use(gin.Recovery())` is setting up error recovery middleware for the Gin router. This middleware
	// catches any panics that occur during request processing and returns an error response with a 500
	// status code. This can help prevent the server from crashing due to unexpected errors and provide a
	// more user-friendly error message to clients.
	r.Use(gin.Recovery())

	// `route.InitializeRoutes(r)` is calling a function named `InitializeRoutes` from the `route` package
	// and passing in the `r` router instance as an argument. This function sets up all the necessary
	// routes for the application using the Gin framework.
	route.InitializeRoutes(r)

	r.Run(":8080")
}
