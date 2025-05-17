package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// This file is part of the Go Bcrypt package.
// It is a simple wrapper around the bcrypt library.
// It is used to hash passwords and verify them.
// It is not intended to be used for anything else.

func main() {
	s := "password123"
	p, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Hashed Pass", string(p))
	fmt.Println("Hashed Pass as slice of byte", p)

	println("----------------------")
	// Verify the password
	loginPswd := "password123"

	err = bcrypt.CompareHashAndPassword(p, []byte(loginPswd))
	if err != nil {
		fmt.Println("Password does not match")
	} else {
		fmt.Println("Password matches")
	}
}
