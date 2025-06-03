package main

import "fmt"

func main() {
	// This code works because:
	// 1. We create a buffered channel with capacity 1, allowing one value to be sent
	//    without a corresponding receiver being ready.
	// 2. We send the value 42 into the channel, which is stored in the buffer.
	// 3. We then receive from the channel, which retrieves the value from the buffer.
	// If this were an unbuffered channel (capacity 0), we would need a separate
	// goroutine to receive the value, otherwise it would deadlock.
	c := make(chan int, 1)
	c <- 42

	fmt.Println(<-c)
}
