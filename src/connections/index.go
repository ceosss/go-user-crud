package connections

import (
	"context"
	"fmt"

	helper "github.com/ceosss/go-user-crud/src/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var UsersCollection *mongo.Collection

func InitializeMongo() {
	var err error

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	Client, err = mongo.NewClient(clientOptions)
	helper.HandleError(err)

	err = Client.Connect(context.Background())
	helper.HandleError(err)

	err = Client.Ping(context.Background(), nil)
	helper.HandleError(err)

	fmt.Println("Connected to MongoDB")

	UsersCollection = Client.Database("go-user-crud").Collection("users")
}
