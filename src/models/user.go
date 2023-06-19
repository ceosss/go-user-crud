package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// The above code defines a Go struct type called User with fields for ID, name, age, gender, and date
// of birth.
// @property ID - ID is a unique identifier for the user, represented as a primitive.ObjectID. It is
// used to uniquely identify each user in the database.
// @property {string} Name - Name is a string property of the User struct, which represents the name of
// the user.
// @property {int} Age - The age of the user, represented as an integer.
// @property {string} Gender - Gender is a property of the User struct which represents the gender of
// the user. It is of type string and is represented in the JSON output as "gender".
// @property DOB - DOB stands for Date of Birth. It is a property of the User struct and is of type
// time.Time. It represents the date of birth of a user.
type User struct {
	ID     primitive.ObjectID `json:"_id"`
	Name   string             `json:"name"`
	Age    int                `json:"age"`
	Gender string             `json:"gender"`
	DOB    time.Time          `json:"dob"`
}
