package main

import "fmt"

func main() {
	f := incrementor()
	fmt.Printf("value: %v\tMem: %v\n", f(), &f)
	fmt.Printf("value: %v\tMem: %v\n", f(), &f)
	fmt.Printf("value: %v\tMem: %v\n", f(), &f)

	println("--------------")

	// here, like in 'f' var above the function assigned to the 'g' variable will
	// be in the same memory place.
	// so when you call, it will have the same value stored into deeper variable 'x'
	g := incrementor()
	fmt.Printf("value: %v\tMem: %v\n", g(), &g)
	fmt.Printf("value: %v\tMem: %v\n", g(), &g)
	fmt.Printf("value: %v\tMem: %v\n", g(), &g)

	println("--------------")

	// if you don't assign to a variable it will return initial state of variable 'x'
	// because it's stored in a different place in memory
	fmt.Printf("Mem: %v\n", incrementor())
	fmt.Printf("Mem: %v\n", incrementor())
	fmt.Printf("Mem: %v\n", incrementor())
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
