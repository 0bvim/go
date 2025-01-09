package main

import "fmt"

type Person struct {
	First string
	Age   int
}

func (p *Person) Speak() {
	fmt.Printf("%s is %d years old\n", p.First, p.Age)
}

func main() {
	person := Person{
		First: "Vinicius",
		Age:   32,
	}
	person.Speak()
}
