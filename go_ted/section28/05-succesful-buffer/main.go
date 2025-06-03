package main

import "fmt"

func main() {
	// This code works because we're creating a buffered channel with a capacity of 2.
	// A buffered channel can store multiple values without needing a receiver waiting.
	// We send 42 and 43 to the channel, which both fit within its capacity.
	// Then we read just one value (42) from the channel with the receive operation <-c.
	// If the channel wasn't buffered or had capacity < 2, we would get a deadlock
	// because we'd be sending values with no goroutine to receive them.
	c := make(chan int, 2)
	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
}
