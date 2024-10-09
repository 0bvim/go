package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(300)
	fmt.Printf("Value of x is %v\t", x)

	switch {
	case x <= 100:
		fmt.Println("between 0 and 100")
	case x <= 200:
		fmt.Println("between 101 and 200")
	case x <= 250:
		fmt.Println("between 201 and 250")
	default:
		fmt.Println("this was more than 250")
	}
	//TODO: select statement
}
