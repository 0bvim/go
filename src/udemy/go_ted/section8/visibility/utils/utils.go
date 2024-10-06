package utils

import "fmt"

// this is to "say" when a function in visible or not visible

// not capitalized function are like "private" functions
// can only be called in the same package
func printName(n string) {
	fmt.Println(n)
}

// capitalized functions are like public function
// can be called by another packages
func NameToBePrinted(name string) {
	printName(name)
}
