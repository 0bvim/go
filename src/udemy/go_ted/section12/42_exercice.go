package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 42; i++ {
		x := rand.Intn(5)

		switch x {
		case 0:
			fmt.Printf("iteration %v \t x is %v x is 0\n", i, x)
		case 1:
			fmt.Printf("iteration %v \t x is %v x is 1\n", i, x)
		case 2:
			fmt.Printf("iteration %v \t x is %v x is 2\n", i, x)
		case 3:
			fmt.Printf("iteration %v \t x is %v x is 3\n", i, x)
		case 4:
			fmt.Printf("iteration %v \t x is %v x is 4\n", i, x)
		default:
			fmt.Printf("iteration %v \t x is %v x isn't\n", i, x)
		}
	}
}
