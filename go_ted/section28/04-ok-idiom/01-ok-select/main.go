package main

import "fmt"

func main() {
	// This code works by using Go's goroutines and channels to demonstrate concurrent programming:
	//
	// 1. We create three channels:
	//    - eve: for even numbers
	//    - odd: for odd numbers
	//    - quit: to signal when processing is complete
	//
	// 2. We start a goroutine with send() that:
	//    - Loops through numbers 0-9
	//    - Sends even numbers to the eve channel
	//    - Sends odd numbers to the odd channel
	//    - Closes the quit channel when done
	//
	// 3. The receive() function:
	//    - Uses a select statement to wait for data from any channel
	//    - Prints values from each channel as they arrive
	//    - Terminates when the quit channel closes
	//
	// The select statement allows non-blocking reception from multiple channels,
	// demonstrating Go's concurrent communication patterns.

	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(eve, odd, quit)

	receive(eve, odd, quit)
}

func receive(e, o <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("from even the eve channel:", v)
		case v := <-o:
			fmt.Println("from even the odd channel:", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("from comma ok:", i, ok)
				return
			} else {
				fmt.Println("from comma ok:", i)
			}
		}
	}
}

func send(e, o chan<- int, quit chan<- bool) {
	for i := range 10 {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(quit)
}
