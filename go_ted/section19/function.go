package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	foo()
	foo1("xpto")
	fmt.Println(foo2("I want"))
	rtn, err := foo3("I want", " more")
	if err != nil {
		log.Fatalf("error while calling foo3: %v", err)
	}
	fmt.Println(rtn)
	fmt.Println(dogYears("Vinicius", 31))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7))

	// unfurling a slice
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(x...))

}

// no params and no return
func foo() {
	fmt.Println("I am from foo")
}

// 1 param no return
func foo1(str string) {
	fmt.Println("I am from foo", str)
}

// 1 param and 1 return
func foo2(str string) string {
	return str + " somenthing"
}

// 2 param and 2 return
func foo3(str string, str2 string) (string, error) {
	if str == "" || str2 == "" {
		return "", errors.New("something went wrong")
	}
	return str + str2, nil
}

func dogYears(name string, age int) (string, int) {
	return fmt.Sprint(name, " is this old in dog years "), age * 7
}

// variadic function
func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)

	n := 0
	for _, value := range ii {
		n += value
	}
	return n
}
