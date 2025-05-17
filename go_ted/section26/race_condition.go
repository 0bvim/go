package main

import (
	"fmt"
	"runtime"
	"sync"
)

func withoutRaceCondition() {
	counter := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)
	var mu sync.Mutex

	for range gs {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines nb:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Final Counter:", counter)
}
func withRaceCondition() {
	counter := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for range gs {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines nb:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Final Counter:", counter)
}

// code with race condition
func main() {
	// withRaceCondition()
	withoutRaceCondition()
}
