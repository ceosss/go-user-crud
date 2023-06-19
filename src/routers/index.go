package router

import (
	route "github.com/ceosss/go-user-crud/src/routes"
	"github.com/gin-gonic/gin"
)

func InitializeRouter() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	route.InitializeRoutes(r)

	r.Run(":8080")
}
