package main

func main() {
	println(factorial(4))
}

func factorial(n int) int {
	println("this is n", n)
	if n == 0 {
		return 1
	}
	return factorial(n-1) * n
}
