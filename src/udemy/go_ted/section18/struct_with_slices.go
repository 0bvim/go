package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

func main() {
	p1 := person{
		firstName:   "John",
		lastName:    "Smith",
		favIceCream: []string{"Vanilla", "Chocolatte"},
	}

	p2 := person{
		firstName:   "Paul",
		lastName:    "McCartney",
		favIceCream: []string{"Strawberry", "Mint"},
	}

	{
		fmt.Println(p1.firstName, p1.lastName+" likes of following flavors:")
		for i, v := range p1.favIceCream {
			fmt.Printf("%v - %v\n", i+1, v)
		}

		println()
		fmt.Println(p2.firstName, p2.lastName+" likes of following flavors:")
		for i, v := range p2.favIceCream {
			fmt.Printf("%v - %v\n", i+1, v)
		}
	}
	fmt.Println("----------------------")
	{
		people_store := map[string]person{
			p1.lastName: p1,
			p2.lastName: p2,
		}

		fmt.Println(people_store[p1.lastName].firstName, people_store[p1.lastName].lastName)
		for _, value := range people_store {
			for _, v2 := range value.favIceCream {
				fmt.Println(v2)
			}
		}
	}
}
