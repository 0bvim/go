package main

import "fmt"

func main() {
	// when we use defer (like 'adiar' in pt_br) it's like delay when the action of function will happens
	{
		foo()
		bar()
	}
	// in this case foo will be printed after bar
	{
		defer foo()
		bar()
	}
}

// func prototype
// func (r receiver) identifier(p parameter(s)) (r return(s)) { <code? }

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
