package main

import "fmt"

func main() {
	// This code won't work because:
	// 1. We're creating a receive-only channel with '<-chan int'
	// 2. You can't send values to a receive-only channel
	c := make(<-chan int, 2)
	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
}
