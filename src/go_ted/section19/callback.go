package main

func main() {
	println(doMath(41, 1, add))
	println(doMath(43, 1, subtract))
	println(doMath(21, 2, multiply))
	println(doMath(84, 2, division))
}

// callback function are like function that receive a function pointer as param in 'c'
func doMath(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func division(a, b int) int {
	return a / b
}

func multiply(a, b int) int {
	return a * b
}
