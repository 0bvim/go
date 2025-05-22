package main

import (
	"fmt"
	"sync"
)

func incrementer(wg *sync.WaitGroup, c *int, mu *sync.Mutex) {
	defer wg.Done()

	mu.Lock()
	*c++
	mu.Unlock()
}

func start() {
	counter := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)
	var mu sync.Mutex

	for range gs {
		go incrementer(&wg, &counter, &mu)
	}

	wg.Wait()
	fmt.Printf("Final counter %v\n", counter)
}

func main() {
	start()
}
