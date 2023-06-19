package main

import (
	"fmt"

	"github.com/ceosss/go-user-crud/src/connections"
	helper "github.com/ceosss/go-user-crud/src/helpers"
	router "github.com/ceosss/go-user-crud/src/routers"
	"github.com/go-playground/validator/v10"
)

// The main function initializes a MongoDB connection, a router, and a validator for a user CRUD
// application in Go.
func main() {
	fmt.Println("User Crud")

	// `connections.InitializeMongo()` is initializing a connection to a MongoDB database.
	connections.InitializeMongo()

	// `router.InitializeRouter()` is initializing the router for the user CRUD application in Go. This
	// function sets up the routes and handlers for the application, allowing it to handle incoming
	// requests and respond with the appropriate data.
	router.InitializeRouter()

	helper.Validator = validator.New()
}
