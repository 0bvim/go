package main

import (
	"fmt"
	"slices"
)

func main() {
	xi := []int{42, 43, 44}
	fmt.Println(xi)
	fmt.Println("----------------------")

	// can appende variadic values
	xi = append(xi, 12, 13, 14, 15, 555, 111, 222)
	fmt.Println(xi)
	fmt.Println("----------------------")

	xi = append(xi, 6943, 234, 2349)
	fmt.Println(xi)
	fmt.Println("----------------------")

	// print len exclusive (one way)
	fmt.Println(xi[len(xi)-1])
	fmt.Println("----------------------")

	// [inclusive:exclusive]
	fmt.Printf("inc/exc xi: %#v\n", xi[0:4])
	fmt.Println("----------------------")

	// [exclusive]
	fmt.Printf("exclusive xi: %#v\n", xi[:7])
	fmt.Println("----------------------")

	// [inclusive]
	fmt.Printf("inclusive xi: %#v\n", xi[7:])
	fmt.Println("----------------------")

	// [all] can be just variable name too
	fmt.Printf("all xi: %#v\n", xi[:])
	fmt.Println("----------------------")

	// to delete or remove
	xi = append(xi[:4], xi[5:]...)
	fmt.Printf("deleting xi: %v\n", xi)
	fmt.Println("----------------------")

	// copy slice
	xiCopy := make([]int, len(xi))
	copy(xiCopy, xi)
	fmt.Println("copy xi:", xiCopy)
	fmt.Println("----------------------")

	// reverse slice
	slices.Reverse(xi)
	fmt.Printf("reverse xi: %v\n", xi)
	fmt.Println("----------------------")

	// sort slice
	slices.Sort(xi)
	fmt.Printf("sorted xi: %v\n", xi)
	fmt.Println("----------------------")
}
