package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 100; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)

		fmt.Printf("iteration %v \tValue of x %v and y %v\t", i, x, y)

		if x < 4 && y < 4 {
			fmt.Println("Both less than 4")
		} else if x > 6 && y > 6 {
			fmt.Println("Both greater than 6")
		} else if x >= 4 && x <= 6 {
			fmt.Println("between 4 and 6")
		} else if y != 5 {
			fmt.Println("y is differ from 5")
		} else {
			fmt.Println("none of the previous cases were met")
		}
	}
}
