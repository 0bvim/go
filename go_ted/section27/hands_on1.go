package main

import (
	"fmt"
	"sync"
)

func toPrint(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Print something", num)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go toPrint(1, &wg)
	go toPrint(2, &wg)

	wg.Wait()
}
