package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 40

	// statement; statement
	if z := 2 * rand.Intn(x); z >= x {
		fmt.Println("z is greater or equal x")
	} else {
		fmt.Println("z is lesser than x")
	}
}
