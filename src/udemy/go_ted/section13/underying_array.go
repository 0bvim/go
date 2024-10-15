package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := make([]int, len(a))
	copy(b, a)
	b[0] = 7
	fmt.Printf("b: %v\n", b)
	fmt.Printf("a: %v\n", a)

}
