package handler

import (
	"context"

	"github.com/ceosss/go-user-crud/src/connections"
	helper "github.com/ceosss/go-user-crud/src/helpers"
	"github.com/ceosss/go-user-crud/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUsers() []models.User {
	var users []models.User

	cursor, err := connections.UsersCollection.Find(context.Background(), bson.M{})
	helper.HandleError(err)

	err = cursor.All(context.Background(), &users)
	helper.HandleError(err)

	return users
}

func GetUser(id primitive.ObjectID) models.User {
	var user models.User

	err := connections.UsersCollection.FindOne(context.Background(), bson.M{"_id:": id}).Decode(&user)
	helper.HandleError(err)

	return user
}

func DeleteUser(id primitive.ObjectID) {
	connections.UsersCollection.DeleteOne(context.Background(), bson.M{"_id": id})
}

func UpdateUser(id primitive.ObjectID, user *models.User) *models.User {
	_, err := connections.UsersCollection.ReplaceOne(context.Background(), bson.M{"_id": id}, user)
	helper.HandleError(err)

	return user
}

func CreateUser(user *models.User) *models.User {
	connections.UsersCollection.InsertOne(context.Background(), user)

	return user
}
