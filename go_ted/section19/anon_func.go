package main

import "fmt"

func main() {
	// not anonymous
	foo()

	// anonymous
	func() {
		fmt.Println("anonymous function")
	}()

	// anonymous function with param
	func(s string) {
		fmt.Println("Anon func showing my name", s)
	}("Nivi")

	// can assign it to a variable too
	x := func(s string) {
		fmt.Println("Anon func showing my name to a func assigned to a variable", s)
	}
	x("Thompson")
}

func foo() {
	fmt.Println("foo")
}
