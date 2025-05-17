package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Contact struct {
	Email string `validate:"required,email"`
	Phone string `validate:"required"`
}

type Person struct {
	Name string `validate:"required"`
	Age  int    `validate:"required,gte=18"`
	// if you remove dive and runs code it will pass in validation
	Contacts []Contact `validate:"required,dive,required"`
}

func main() {
	validate := validator.New()

	// Creating a sample Person with invalid Contacts
	person := Person{
		Name: "John Doe",
		Age:  25,
		Contacts: []Contact{
			{Email: "", Phone: "1234567890"},    // Email is required but missing
			{Email: "invalid-email", Phone: ""}, // Invalid email format and missing phone
		},
	}

	// Validating the struct
	err := validate.Struct(person)
	if err != nil {
		// This will print out validation errors for each Contact
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Validation failed on field '%s' for condition '%s'\n", err.Field(), err.Tag())
		}
	} else {
		fmt.Println("Validation passed!")
	}
}
