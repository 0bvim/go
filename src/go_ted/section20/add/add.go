package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func DoMath(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func Paradise(loc string) string {
	return fmt.Sprint("My idea of paradise is: ", loc)
}

func main() {
	fmt.Printf("%v\n", Add(2, 3))
	fmt.Printf("%v\n", Subtract(2, 3))
	fmt.Printf("%v\n", DoMath(2, 3, Add))
	fmt.Printf("%v\n", DoMath(2, 3, Subtract))
	fmt.Printf("%v\n", Paradise("Canada"))
}
