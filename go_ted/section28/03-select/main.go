package main

import "fmt"

func main() {
	// This code demonstrates Go's concurrency model using channels:
	// 1. We create three channels: eve (for even numbers), odd (for odd numbers), and quit (for termination)
	// 2. The send() function runs in a separate goroutine, sending even numbers to eve channel,
	//    odd numbers to odd channel, and finally sending a signal to quit channel
	// 3. The receive() function selects from all three channels, handling values as they arrive
	// 4. When the quit channel receives a value, the program terminates
	// This showcases Go's select statement which allows waiting on multiple channel operations,
	// and goroutines which enable concurrent execution
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)

	receive(eve, odd, quit)
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from even the eve channel:", v)
		case v := <-o:
			fmt.Println("from even the odd channel:", v)
		case v := <-q:
			fmt.Println("from quit channel:", v)
			return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := range 100 {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0
}
