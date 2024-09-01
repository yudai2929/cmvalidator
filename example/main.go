package main

import (
	"errors"
	"github.com/yudai2929/cmvalidator"
)

type User struct {
	FirstName string `validate:"required" customMessage:"First Name is required!"`
	LastName  string `validate:"required" customMessage:"Last Name is required!"`
	Age       int    `validate:"gte=0,lte=130" customMessage:"Age must be between 0 and 130!"`
	Email     string `validate:"required,email" customMessage:"Email is invalid!"`
}

func main() {

	// Create a new instance of the validator
	validate := cmvalidator.New()

	// Create a new instance of the struct to validate
	newUser := User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       200, // invalid age
		Email:     "john@example.com",
	}

	// Validate the struct
	if err := validate.Struct(newUser); err != nil {
		var cmvalidateErrors cmvalidator.CFValidateErrors
		if errors.As(err, &cmvalidateErrors) {
			// Handle the error
			for _, cmvalidateError := range cmvalidateErrors {
				// Print the custom message
				println("Custom Message:", cmvalidateError.CustomMessage())
			}
		} else {
			// Handle the error
		}
	}

}
