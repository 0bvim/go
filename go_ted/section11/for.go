package main

import "fmt"

func main() {
	var y int
	// basic for loop
	// I try use pre increment, and it not work in go I guess
	fmt.Println("Like C for")
	for i := 0; i < 5; i++ {
		fmt.Printf("Value of i == %d\n", i)
	}

	// like a while loop in C
	fmt.Println("\nLike C while")
	for y < 10 {
		fmt.Printf("Value of y == %d\n", y)
		y++
	}

	// like a for(;;) in C
	fmt.Println("\nLike C 'for(;;)'")
	for {
		fmt.Printf("Y value is: %d\n", y)
		if y > 20 {
			break
		}
		y++
	}

	fmt.Println("\n continue")
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("Even numbers: %d\n", i)
	}

	fmt.Println("\n nested loops")
	for i := 0; i < 5; i++ {
		fmt.Println("--")
		for j := 0; j < 5; j++ {
			fmt.Printf("Outer loop %v \t inner loop %v\n", i, j)
		}
	}
}
