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

// The function initializes a connection to a MongoDB database and sets up a collection for users.
func InitializeMongo() {
	var err error

	// `clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")` is creating a new
	// `options.Client` object and setting the URI to connect to the MongoDB database running on the local
	// machine at port 27017. This object is then used to configure the `mongo.NewClient` instance that will
	// be used to interact with the database.
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// `Client, err = mongo.NewClient(clientOptions)` is creating a new `mongo.Client` instance with the
	// provided client options. It returns the client instance and an error (if any). The client instance
	// is used to interact with the MongoDB database.
	Client, err = mongo.NewClient(clientOptions)
	helper.HandleError(err)

	// `err = Client.Connect(context.Background())` is establishing a connection to the MongoDB database
	// using the `mongo.Client` instance `Client`. The `context.Background()` function is used to create a
	// new context with no values and is passed as an argument to the `Connect` method. The `Connect`
	// method returns an error (if any) which is assigned to the `err` variable.
	err = Client.Connect(context.Background())
	helper.HandleError(err)

	// `err = Client.Ping(context.Background(), nil)` is pinging the MongoDB server to check if the
	// connection is established and working properly. The `Ping` method takes a context and an options
	// parameter, which is set to `nil` in this case. If the ping is successful, no error is returned. If
	// there is an error, it is assigned to the `err` variable.
	err = Client.Ping(context.Background(), nil)
	helper.HandleError(err)

	fmt.Println("Connected to MongoDB")

	// `UsersCollection = Client.Database("go-user-crud").Collection("users")` is setting up a collection
	// named "users" in the "go-user-crud" database. It uses the `Client` instance to access the
	// "go-user-crud" database and then creates a collection named "users" within that database. The
	// `UsersCollection` variable is then assigned to this collection, allowing the program to interact
	// with the "users" collection in the "go-user-crud" database.
	UsersCollection = Client.Database("go-user-crud").Collection("users")
}
