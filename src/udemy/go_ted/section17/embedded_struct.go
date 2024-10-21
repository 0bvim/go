package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	// agent person // this way you need to use dots to access
	person // this way is really embedded, you just need to populate and can access inside field directly
	ltk    bool
	first  string // you can hava a parameter with same name that have in embedded struct param
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   27,
		},
		first: "Steve vai",
		ltk:   true,
	}

	p2 := person{
		first: "Jenny",
		last:  "MoneyPenny",
		age:   27,
	}

	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Printf("%T\t%#v\n", sa1, p2)
	fmt.Printf("sa1: %v\n", sa1)

	fmt.Printf("sa1.first: %v\n", sa1.age)
	fmt.Printf("sa1.last: %v\n", sa1.last)
	fmt.Printf("sa1.age: %v\n", sa1.age)
	fmt.Println(sa1.first, sa1.person.first)
}
