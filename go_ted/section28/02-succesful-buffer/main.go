package main

import "fmt"

func main() {
	// This code works because:
	// 1. We create a channel 'c' that can send and receive integers.
	// 2. We start a goroutine (concurrent function) that sends the value 42 to the channel.
	// 3. The main goroutine waits to receive a value from the channel (with <-c).
	// 4. The program doesn't exit early because the channel operation synchronizes the goroutines.
	// 5. Once the value is received, it's printed and the program completes.
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
