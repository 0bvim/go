package main

import "fmt"

// simple type example
type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

// method implementation
func (p person) speak() {
	fmt.Println("My name is", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("Im a secret agent", sa.first)
}

// interface help us to have polymorphism
// interfaces in Go define a set of method signatures.

type human interface {
	speak()
}

// function to receive a parameter of human interface and "speak"
// when a type implements the called method in interfaces, it are of this interfaces type
// like secreAgent and person

func saySomething(h human) {
	h.speak()
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}

	p2 := person{
		first: "Janny",
	}

	saySomething(sa1)
	saySomething(p2)
}
