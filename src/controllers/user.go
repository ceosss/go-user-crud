package controller

import (
	"net/http"

	"github.com/ceosss/go-user-crud/src/handler"
	helper "github.com/ceosss/go-user-crud/src/helpers"
	"github.com/ceosss/go-user-crud/src/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// This function retrieves users and returns them as a JSON response using the Gin framework in Go.
func GetUsers(c *gin.Context) {
	users := handler.GetUsers()
	c.JSON(http.StatusOK, users)
}

// This function retrieves a user by their ID and returns their information in JSON format.
func GetUser(c *gin.Context) {
	userId := c.Param("id")

	id, err := primitive.ObjectIDFromHex(userId)
	helper.HandleError(err)

	user := handler.GetUser(id)

	c.JSON(http.StatusOK, user)
}

// This function creates a user by binding JSON data from a request, creating a user using a handler
// function, and returning the created user as a JSON response.
func CreateUser(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)
	helper.HandleError(err)

	user = *handler.CreateUser(&user)

	c.JSON(http.StatusOK, user)
}

// This function updates a user's information in a database using their ID.
func UpdateUser(c *gin.Context) {
	userId := c.Param("id")

	id, err := primitive.ObjectIDFromHex(userId)
	helper.HandleError(err)

	var user models.User

	err = c.BindJSON(&user)
	helper.HandleError(err)

	handler.UpdateUser(id, &user)

	c.JSON(http.StatusOK, user)
}

// The function deletes a user with the specified ID and returns a JSON response with a status code of
// 200.
func DeleteUser(c *gin.Context) {
	userId := c.Param("id")

	id, err := primitive.ObjectIDFromHex(userId)
	helper.HandleError(err)

	handler.DeleteUser(id)

	c.JSON(http.StatusOK, nil)
}
