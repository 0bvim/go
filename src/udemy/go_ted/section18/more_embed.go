package main

import "fmt"

type engine struct {
	eletric bool
}

type vehicle struct {
	engine engine
	make   string
	model  string
	color  string
	doors  int8
}

func main() {
	v1 := vehicle{
		engine: engine{
			eletric: false,
		},
		make:  "idk",
		model: "Ferrari Black",
		color: "Dark black deep blue",
		doors: 4,
	}
	v2 := vehicle{
		engine: engine{
			eletric: true,
		},
		make:  "idk",
		model: "Mazda Speed",
		color: "Red",
		doors: 5,
	}
	{
		fmt.Println(v1, v2)
		fmt.Printf("V1 Specs:\neletric: %v\nMake: %v\nModel: %v\nColor: %v\nDoors: %v\n",
			v1.engine.eletric, v1.make, v1.model, v1.color, v1.doors)
		fmt.Printf("\nV2 Specs:\neletric: %v\nMake: %v\nModel: %v\nColor: %v\nDoors: %v\n",
			v2.engine.eletric, v2.make, v2.model, v2.color, v2.doors)
	}

	// about anonymous struct again
	{
		p1 := struct {
			first     string
			friends   map[string]int
			favDrinks []string
		}{
			first: "vinicius",
			friends: map[string]int{
				"stete": 24,
				"lila":  4,
			},
			favDrinks: []string{"Spiced Chai", "Fruit Juice"},
		}
		fmt.Println(p1)
		fmt.Printf("My first name is %v\n", p1.first)
		for key, value := range p1.friends {
			fmt.Printf("My friend %v is %v years old: \n", key, value)
		}

		fmt.Println("favDrinks:")
		for _, drink := range p1.favDrinks {
			fmt.Println(drink)
		}
	}
}
