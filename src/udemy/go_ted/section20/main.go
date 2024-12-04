package main

import (
	. "fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	Println(sum(numbers))
	Println(bar())
	Println(foo())
	Println(foo2(numbers...))
	Println(bar2(numbers))

}

// hands on 57
// this named return function don't need to
// specify 'return total' just return do the job!
func sum(ii []int) (total int) {
	total = 0
	for _, v := range ii {
		total += v
	}
	return
}

// hands on 58
func foo() int {
	return 42
}

func bar() (int, string) {
	return 42, "Jinboo" +
		""
}

func foo2(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func bar2(numbers []int) (total int) {
	total = 0
	for _, n := range numbers {
		total += n
	}
	return
}
