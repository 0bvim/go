package main

import "fmt"

// this exercice is about data type, aka, aggregate data type, aka, complex data type

type Engine struct {
	engineType string
	fuelType   string
	cylinders  int
}

func (e *Engine) Start() {
	fmt.Println("Engine Start")
}

type Car struct {
	Engine     // embedding the Engine Struct
	color      string
	model      string
	doorNumber int
}

func (c *Car) Bip() {
	fmt.Println("Car Bip")
}

func main() {
	car := Car{}
	car.Start()
	car.Bip()

}
