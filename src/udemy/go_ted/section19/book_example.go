package main

import (
	"fmt"
	"log"
	"strconv"
)

// link of example https://www.alexedwards.net/blog/interfaces-explained

// this is my implementation of a method that satisfy
// String() string (fmt.Stringer()) interface
type Item struct {
	Type string
	Name string
}

func (it Item) String() string {
	return fmt.Sprintf("Type: %s - Name: %s", it.Type, it.Name)
}

// Declare a Book type which satisfies the fmt.Stringer interface.
type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

// Declare a Count type which satisfies the fmt.Stringer interface.
type Count int

func (c Count) String() string {
	return strconv.Itoa(int(c))
}

// Declare a writeLog function which takes any object that satisfies
// the fmt.Stringer interface.
func WriteLog(s fmt.Stringer) {
	log.Println(s.String())
}

func main() {
	// Initialize a Count object and pass it to WriteLog().
	book := Book{"Alice in Wonderland", "Lewis Carrol"}
	WriteLog(book)

	// Initialize a Count object and pass it to WriteLog().
	count := Count(3)
	WriteLog(count)

	item := Item{
		Type: "Eletronic",
		Name: "Notebook",
	}
	WriteLog(item)
}
