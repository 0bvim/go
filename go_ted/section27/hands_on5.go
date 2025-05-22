package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func incrementer(wg *sync.WaitGroup, c *int64) {
	defer wg.Done()
	atomic.AddInt64(c, 1)
}

func start() {
	// using atomic in counter
	var counter int64

	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)
	// mutex will no longer needed in this case
	// var mu sync.Mutex

	for range gs {
		go incrementer(&wg, &counter)
	}

	wg.Wait()
	fmt.Printf("Final counter %v\n", counter)
}

func main() {
	start()
}
