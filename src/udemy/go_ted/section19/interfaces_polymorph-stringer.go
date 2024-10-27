package main

import (
	"fmt"
	"strconv"
)

// basic struct type
type book struct {
	title string
}

// method implemetation
func (b book) String() string {
	return fmt.Sprint("Book title: ", b.title)
}

// basic type
type count int

// method implementation
func (c count) String() string {
	return fmt.Sprint("The number is ", strconv.Itoa(int(c)))
}

func main() {
	b := book{
		title: "Learning to learn",
	}

	var c count = 42

	fmt.Println(b)
	fmt.Println(c)
}
