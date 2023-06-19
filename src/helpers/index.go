package helper

import "github.com/go-playground/validator/v10"

var Validator *validator.Validate

// The function handles errors by panicking if an error is not nil.
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
