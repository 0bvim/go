package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(50)

	// common switch
	switch {
	case x > 42:
		fmt.Println("x > 42")
	case x < 42:
		fmt.Println("x < 42")
	case x == 42:
		fmt.Println("x == 42")
	default:
		fmt.Println("default case for x")
	}

	// less readable switch
	switch x {
	case 40:
		fmt.Println("40")
	case 41:
		fmt.Println("41")
	case 42:
		fmt.Println("42")
	default:
		fmt.Println("default case for x")
	}

	// fallthrough
	y := 40
	switch y {
	case 40:
		fmt.Println("40")
		fallthrough
	case 41:
		fmt.Println("41")
	case 42:
		fmt.Println("42")
	default:
		fmt.Println("default case for x")
	}

	//full fallthrough
	switch y {
	case 40:
		fmt.Println("40")
		fallthrough
	case 41:
		fmt.Println("41")
		fallthrough
	case 42:
		fmt.Println("42")
		fallthrough
	default:
		fmt.Println("default case for x")
	}
}
