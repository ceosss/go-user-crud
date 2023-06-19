# Go User CRUD Project

This project is a simple User CRUD (Create, Read, Update, Delete) application built with Go and MongoDB.

## Prerequisites

Before running this project, make sure you have the following prerequisites installed:

- Go: [Installation Guide](https://golang.org/doc/install)
- MongoDB: [Installation Guide](https://docs.mongodb.com/manual/installation)
- Gin: [Installation Guide](https://github.com/gin-gonic/gin)

## Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/ceosss/go-user-crud.git
   cd go-user-crud

2. Install the project dependencies:

    ```shell
    go mod download
 
3. Update the MongoDB connection URI:

    Open the src/connections/index.go file and update the MongoDB connection URI based on your MongoDB server configuration.

4. Run the application:

      ```shell
      go run main.go
      
## Usage
The application will start running on the specified port (default: 8080).
Use a tool like cURL or Postman to interact with the API endpoints.

The available endpoints are:

- `GET /users`: Get all users.
- `GET /users/{id}`: Get a user by ID.
- `POST /users`: Create a new user.
- `PUT /users/{id}`: Update a user by ID.
- `DELETE /users/{id}`: Delete a user by ID.

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please create a new issue or submit a pull request.

When contributing to this project, please follow the existing code style and conventions. Additionally, ensure that your changes are well-tested.
