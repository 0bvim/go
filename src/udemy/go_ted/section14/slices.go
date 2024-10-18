package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	{
		for i, v := range x {
			fmt.Printf("%v \t %T \t %#v\n", v, v, i)
		}

		// sharp-v to print type of variable or data type in go-like-syntax
		name := "string"
		nb := -1
		p := Person{"Ken"}
		fmt.Printf("%#v\n %#v\n %#v\n %#v\n", x, name, nb, p)
		fmt.Printf("%v\n %v\n %v\n %v\n", x, name, nb, p)
		fmt.Println("------------")
	}

	//inclusive exclusive
	{
		fmt.Printf("x: %v\n", x[:5])
		fmt.Printf("x: %v\n", x[5:])
		fmt.Printf("x: %v\n", x[2:7])
		fmt.Printf("x: %v\n", x[1:6])
		fmt.Println("------------")
	}

	// appending
	{
		x = append(x, 52)
		fmt.Printf("x: %v\n", x)
		x = append(x, 53, 54, 55)
		fmt.Printf("x: %v\n", x)

		y := []int{56, 57, 58, 59, 60}
		x = append(x, y...)
		fmt.Printf("x: %v\n", x)
		fmt.Println("------------")
	}

	// append and slice to delete
	{
		x = append(x[:3], x[6:]...)
		fmt.Printf("x: %v\n", x)

	}
	{
		x1 := make([]int, 50)
		fmt.Printf("x1: %v\n", x1)
		fmt.Printf("x1: %v\n", len(x1))
		fmt.Printf("x1: %v\n", cap(x1))

		x2 := make([]int, 0, 50)
		fmt.Printf("x2: %v\n", x2)
		fmt.Printf("x2: %v\n", len(x2))
		fmt.Printf("x2: %v\n", cap(x2))

		xs := make([]string, 0, 50)

		xs = append(xs,
			"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut",
			"Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa",
			"Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan",
			"Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire",
			"New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio",
			"Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota",
			"Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia",
			"Wisconsin", "Wyoming",
		)
		fmt.Printf("xs: %v\n", xs)
		fmt.Printf("xs: %v\n", cap(xs))
		fmt.Printf("xs: %v\n", len(xs))

		for i := 0; i < len(xs); i++ {
			fmt.Printf("%v - %v\n", i+1, xs[i])
		}
	}
}
