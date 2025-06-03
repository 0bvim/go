package main

import "fmt"

func main() {
	// This code doesn't work because c is declared as a send-only channel (chan<- int)
	// We're trying to receive from it using <-c in the fmt.Println statements
	// But we can't receive from a send-only channel
	// To fix this, we should declare c as a bidirectional channel:
	// c := make(chan int, 2)
	c := make(chan<- int, 2)
	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
}
