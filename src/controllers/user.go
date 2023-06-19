package controller

import (
	"net/http"

	"github.com/ceosss/go-user-crud/handler"
	helper "github.com/ceosss/go-user-crud/src/helpers"
	"github.com/ceosss/go-user-crud/src/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUsers(c *gin.Context) {
	users := handler.GetUsers()
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	userId := c.Param("id")

	id, err := primitive.ObjectIDFromHex(userId)
	helper.HandleError(err)

	user := handler.GetUser(id)

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)
	helper.HandleError(err)

	user = *handler.CreateUser(&user)

	c.JSON(http.StatusOK, user)
}

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

func DeleteUser(c *gin.Context) {
	userId := c.Param("id")

	id, err := primitive.ObjectIDFromHex(userId)
	helper.HandleError(err)

	handler.DeleteUser(id)

	c.JSON(http.StatusOK, nil)
}
