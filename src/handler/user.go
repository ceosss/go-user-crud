package handler

import (
	"context"

	"github.com/ceosss/go-user-crud/src/connections"
	helper "github.com/ceosss/go-user-crud/src/helpers"
	"github.com/ceosss/go-user-crud/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// The function retrieves all users from a MongoDB collection and returns them as a slice of User
// models.
func GetUsers() []models.User {
	var users []models.User

	cursor, err := connections.UsersCollection.Find(context.Background(), bson.M{})
	helper.HandleError(err)

	err = cursor.All(context.Background(), &users)
	helper.HandleError(err)

	return users
}

// This function retrieves a user from a MongoDB database by their ID.
func GetUser(id primitive.ObjectID) models.User {
	var user models.User

	err := connections.UsersCollection.FindOne(context.Background(), bson.M{"_id:": id}).Decode(&user)
	helper.HandleError(err)

	return user
}

// The function deletes a user from a MongoDB collection based on their ID.
func DeleteUser(id primitive.ObjectID) {
	connections.UsersCollection.DeleteOne(context.Background(), bson.M{"_id": id})
}

// The function updates a user in a MongoDB collection and returns the updated user.
func UpdateUser(id primitive.ObjectID, user *models.User) *models.User {
	_, err := connections.UsersCollection.ReplaceOne(context.Background(), bson.M{"_id": id}, user)
	helper.HandleError(err)

	return user
}

// The function creates a new user in a MongoDB collection and returns the created user.
func CreateUser(user *models.User) *models.User {
	connections.UsersCollection.InsertOne(context.Background(), user)

	return user
}
