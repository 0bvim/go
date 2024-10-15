package main

import "fmt"

func main() {
	si := []string{"a", "b", "c"}
	fmt.Printf("si: %v\n", si)

	xi := make([]int, 0, 10)

	fmt.Printf("xi: %v\n", xi)
	fmt.Printf("len(xi): %v\n", len(xi))
	fmt.Printf("cap(xi): %v\n", cap(xi))
	xi = append(xi, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("xi: %v\n", xi)

	fmt.Println("----------------")

	xi = append(xi, 11, 12, 13, 14, 15)
	fmt.Printf("xi: %v\n", xi)
	fmt.Printf("len(xi): %v\n", len(xi))
	fmt.Printf("cap(xi): %v\n", cap(xi))
}
