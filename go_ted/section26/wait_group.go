package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Print the Go version and OS information
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(2)
	go foo(&wg)
	go bar(&wg)

	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo(wg *sync.WaitGroup) {
	for i := range 10 {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar(wg *sync.WaitGroup) {
	for i := range 10 {
		fmt.Println("bar:", i)
	}
	wg.Done()
}
