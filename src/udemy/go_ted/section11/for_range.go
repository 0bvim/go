package main

import (
	"fmt"
)

func main() {
	// for range loop
	// data structures - slice
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Println("ranging over a slice", i, v)
	}

	fmt.Println() // just a line to better output
	// here ignoring counter with underscore
	xx := []int{42, 43, 44, 45, 46, 47}
	for _, v := range xx {
		fmt.Println("ranging over a slice", v)
	}

	fmt.Println()
	// for range loop
	// data structures - map
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	// the same of the slices
	for k, v := range m {
		fmt.Println("ranging over a map", k, v)
	}

}
