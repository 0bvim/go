package main

import "fmt"

func main() {
	// if you wanna use only one of the returned types instead both
	// just put '_' (underscore) to 'ignore it'
	_, age := nameStateAge()
	fmt.Println("My name is", "I'm", age, "years old.")
}
// this function return more than one variable
// need to declare all types returned and return types
func nameStateAge() (string, int) {
	name := "MyName"
	age := 30
	
	return name, age
}
