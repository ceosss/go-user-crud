package main

import (
	"fmt"

	"github.com/ceosss/go-user-crud/src/connections"
	helper "github.com/ceosss/go-user-crud/src/helpers"
	router "github.com/ceosss/go-user-crud/src/routers"
	"github.com/go-playground/validator/v10"
)

func main() {
	fmt.Println("User Crud")

	connections.InitializeMongo()

	router.InitializeRouter()

	helper.Validator = validator.New()
}
