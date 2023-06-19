package helper

import "github.com/go-playground/validator/v10"

var Validator *validator.Validate

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
