package main

import "fmt"

func main() {
	// This code will cause a deadlock because:
	// 1. We create a buffered channel 'c' with capacity 1
	// 2. We send value 42 to the channel (this fills the buffer)
	// 3. We try to send value 43, but since the buffer is full and no one is receiving,
	//    this operation blocks
	// 4. The main goroutine is now stuck trying to send to a full channel
	// 5. No other goroutines exist to receive from the channel
	// Result: deadlock!
	c := make(chan int, 1)
	c <- 42
	c <- 43

	fmt.Println(<-c)
}
