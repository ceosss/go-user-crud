package route

import (
	controller "github.com/ceosss/go-user-crud/src/controllers"
	"github.com/gin-gonic/gin"
)

// The function initializes routes for CRUD operations on user data using the Gin framework in Go.
func InitializeRoutes(r *gin.Engine) {
	r.GET("/users", controller.GetUsers)
	r.GET("/users/:id", controller.GetUser)
	r.POST("/users", controller.CreateUser)
	r.PATCH("/users/:id", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)
}
