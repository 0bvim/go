package main

import (
	"fmt"
	"sort"
)

func main() {
	is := []int{4, 2, 5, 3, 1}
	sort.Ints(is)
	fmt.Println("Sorted int slice", is)

	println("================================")

	ss := []string{"Nivi", "Stete", "Lila"}
	sort.StringSlice(ss).Sort()
	fmt.Println("Sorted string slice", ss)

}
