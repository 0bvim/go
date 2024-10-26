package main

import "fmt"

// simple type example
type person struct {
	first string
}

// method implementation
func (p person) speak() {
	fmt.Println("My name is", p.first)
}

func main() {
	p1 := person{
		first: "James",
	}

	p2 := person{
		first: "Janny",
	}
	p1.speak()
	p2.speak()
}
