package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func withoutRaceConditionAtomic() {
	var counter int64
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for range gs {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("counter", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("Goroutines nb:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Final Counter:", counter)
}

func main() {
	withoutRaceConditionAtomic()
}
