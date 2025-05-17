package main

import "fmt"
// module to verify types of variables
import "reflect"

func main() {
	//slices are like arrays, but better.
	names := []string {"Jack", "Black", "Linus", "Turing"}
	// in slices you can use append function to insert new items
	names = append(names, "Bill Gates")
	names = append(names, "Steve Jobs")
	fmt.Println(names)
	// to verify type of variable and print
	fmt.Println(reflect.TypeOf(names))
	// to print lenght of slice
	fmt.Println("Slice lenght", len(names))
	// to print capacity of a slice
	fmt.Println("Slice capacity", len(names))
}
