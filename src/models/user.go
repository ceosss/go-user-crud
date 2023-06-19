package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID     primitive.ObjectID `json:"_id"`
	Name   string             `json:"name"`
	Age    int                `json:"age"`
	Gender string             `json:"gender"`
	DOB    time.Time          `json:"dob"`
}
