package main

import "fmt"

func main() {
	foo()
	foo1("xpto")
}

// no params and no return
func foo() {
	fmt.Println("I am from foo")
}

// no return
func foo1(str string) {
	fmt.Println("I am from foo", str)
}
