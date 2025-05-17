package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Jenny",
		last:  "MoneyPenny",
		age:   27,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("%T\t%#v\n", p1, p2)
	fmt.Printf("p1: %v\n", p1)

	fmt.Printf("p1.first: %v\n", p1.first)
	fmt.Printf("p1.last: %v\n", p1.last)
	fmt.Printf("p1.age: %v\n", p1.age)
}
