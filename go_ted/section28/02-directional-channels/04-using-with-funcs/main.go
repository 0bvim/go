package main

import "fmt"

func main() {
	// This code demonstrates Go's concurrency model with goroutines and channels:
	// 1. We create an unbuffered channel 'c' that passes integers
	// 2. The foo function runs in a separate goroutine, sending value 42 to channel c
	// 3. The bar function runs in the main goroutine, blocking until it receives a value from c
	// 4. This synchronizes the two goroutines - foo must send before bar can proceed
	// 5. After the value is received and printed, the program continues and exits

	c := make(chan int)

	go foo(c)
	bar(c)

	fmt.Println("about to exit")
}

func foo(c chan<- int) {
	c <- 42
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}
