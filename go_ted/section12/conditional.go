package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(300)
	fmt.Printf("Value of x is %v\t", x)

	if x <= 100 {
		fmt.Print("between 0 and 100\n")
	} else if x <= 200 {
		fmt.Print("between 101 and 200\n")
	} else if x <= 250 {
		fmt.Print("between 201 and 250\n")
	} else {
		fmt.Println("this was more than 250")
	}
}
