package main

import "fmt"

func main() {
	// This code works because:
	// 1. We create a channel 'c' to communicate between goroutines
	// 2. We launch foo() as a goroutine which sends values to the channel
	// 3. The main goroutine reads from the channel using range
	// 4. The range loop automatically terminates when foo() closes the channel
	// 5. The program will exit only after all values have been processed and printed
	c := make(chan int)

	go foo(c)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func foo(c chan<- int) {
	for i := range 100 {
		c <- i + 1
	}
	close(c)
}
