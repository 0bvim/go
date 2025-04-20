package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func MarhsalProcess() {
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
	bs2, _ := json.Marshal(family)
	fmt.Println(string(bs))
	fmt.Println(string(bs2))
}

func UnmarshalProcess() {
	s := `[{"First":"Nivi","Last":"Pereira","Age":32},{"First":"Stete","Last":"Franco","Age":24},{"First":"Lila","Last":"Freitas","Age":4}]`
	bs := []byte(s)

	people := []person{}

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println("Error to unmarshal")
	}

	fmt.Println(people)

	for i, v := range people {
		fmt.Println("\nPerson Nb:", i+1)
		fmt.Println(v.First, v.Last, v.Age)
	}
}

func main() {
	MarhsalProcess()
	UnmarshalProcess()
}
