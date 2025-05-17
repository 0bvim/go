// import package main to function main works
package main

// import modules to use some functions like print, read, open like lib in clang
// fmt is for use 'fmt.Println("args")'
import "fmt"
// 'os' is a module to use 'Exit' and others
import "os"

/*
to compile your code you can use 'go build filename.go' then execute the file (bin file have same name of source file but without extension)

If you wanna just run your code and verify if result is that you've spected, use 'go run filename.go' it will be compiled and will give the output in your terminal and don't keep bin file.

*/
// to call a function just type func
// curl bracket is always at the end of function, golang don't accepet below function like in clang and others.
func main(){
	// variable declaration 'var' 'varname' 'vartype'
	// ';' is optional
	var name string = "GoodNameHere"
	
	// you can declare a variable in other two ways
	var name2 = "AnotherGoodName"
	// with implicit declaration of type and var
	name3 := "MoreOneGoodName"
	// to print in stdout
	fmt.Println("Hello Mr.", name)

	// to concat with variable
	fmt.Println("Hello Mr.", name2)

	// in both cases of print go already set to use a line break at the end '\n', so you don't need to use it, only if you want more than one.

	// golang have 2 type of float, 32 and 64
	var version float32 = 1.1

	// in this I'll use float32 because the number is small
	fmt.Println("Program version:", version)

	// if you note, we don't need too put a space at end of string that will print a variable after. Golang put it automatically.

	// declare a empty variable always will return zero in golangif you execute it, answer will be 'Your age is: 0'. But go don't compile if you set a variable that is not being used.
	var age int
	fmt.Println("Hello Mr.", name3, "Your age is:", age)

	// to get input from user just specify variable that'll store it and use 'fmt.Scan' from "fmt" module. And parameter is address of command variable. You need to figure it out (address)
	var command int
	fmt.Scan(&command)
	
	fmt.Println("Value of command var:", command)

	// working with if/else

	if command == 1 {
		fmt.Println("Positive value")
	} else if command == 2 {
		fmt.Println("Another Positive value")
	} else {
		fmt.Println("Omg, idk what's this value.")
	}
	// if/else and for don't need of parentheses to conditionals
	// conditional in if don't accept only variable as a conditional like in clang e.g. --> if (name), conditionals always have to return a boolean value
	// to exit from program when success number (0) if fail (-1)
	os.Exit(0)
}
