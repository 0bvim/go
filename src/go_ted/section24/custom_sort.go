package main

import (
	"fmt"
	"sort"
)

// ByAge implements sort.Interface for []Person based on
// the Age field.

type Person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func (p Person) String() string {
	return fmt.Sprintf("First: %s, Last: %s, Age: %d", p.First, p.Last, p.Age)
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].First < a[j].First }

func main() {
	p1 := Person{
		First: "Nivi",
		Age:   32,
	}

	p2 := Person{
		First: "Stete",
		Age:   24,
	}

	p3 := Person{
		First: "Lila",
		Age:   4,
	}

	family := []Person{p1, p2, p3}

	fmt.Println("Before sorting by age", family)
	println("================================")

	sort.Sort(ByAge(family))
	fmt.Println("After sorting by age", family)
	println("================================")

	fmt.Println("Before sorting by name", family)
	println("================================")

	sort.Sort(ByName(family))
	fmt.Println("After sorting by name", family)
	println("================================")
}
