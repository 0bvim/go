package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Nivi",
		Last:  "Pereira",
		Age:   32,
	}

	p2 := person{
		First: "Stete",
		Last:  "Franco",
		Age:   24,
	}

	p3 := person{
		First: "Lila",
		Last:  "Freitas",
		Age:   4,
	}

	family := []person{p1, p2, p3}

	bs, _ := json.MarshalIndent(family, "", " ")
	fmt.Println(string(bs))
}
