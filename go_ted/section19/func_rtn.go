package main

import . "fmt"

func main() {
	x := foo()
	Println(x)

	y := bar()
	Println(y())
	Printf("%T\n", foo())
	Printf("%T\n", bar())
	Printf("%T\n", y)
}

func foo() int {
	return 42
}

func bar() func() int {
	return func() int {
		return 43
	}
}
