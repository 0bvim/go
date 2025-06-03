package main

import "fmt"

func main() {
	// This code causes a deadlock because the channel 'c' is unbuffered.
	// When we try to send 42 to the channel with 'c <- 42', the program
	// blocks waiting for a receiver to read from the channel.
	// However, the receiver code 'fmt.Println(<-c)' only runs after the send,
	// so we have a deadlock situation.
	c := make(chan int)
	c <- 42

	fmt.Println(<-c)
}
