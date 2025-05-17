package main

import "fmt"
import "os"

func main(){
	var name string = "GoodNameHere"
	var name2 = "AnotherGoodName"
	name3 := "MoreOneGoodName"

	fmt.Println("Hello Mr.", name)
	fmt.Println("Hello Mr.", name2)

	var version float32 = 1.1
	fmt.Println("Program version:", version)

	var age int
	fmt.Println("Hello Mr.", name3, "Your age is:", age)

	var command int
	fmt.Scan(&command)
	fmt.Println("Value of command var:", command)

	if command == 1 {
		fmt.Println("Positive value")
	} else if command == 2 {
		fmt.Println("Another Positive value")
	} else {
		fmt.Println("Omg, idk what's this value.")
	}
	os.Exit(0)
}
