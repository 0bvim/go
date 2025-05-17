package main

import "fmt"

func main() {
	xs := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)"}
	fmt.Println(xs)
	fmt.Printf("Type of variable\t%T\n", xs)

	for _, v := range xs {
		fmt.Printf("%v\n", v)
	}

	fmt.Println("----------------------")
	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])
	fmt.Println("len of slice", len(xs))

	fmt.Println("----------------------")
	for i := 0; i < len(xs); i++ {
		fmt.Println(xs[i])
	}
}
