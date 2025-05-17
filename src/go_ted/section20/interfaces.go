package main

import (
	"fmt"
	"math"
)

type Square struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

func (s Square) area() float64 {
	return s.width * s.length
}

func (c Circle) area() float64 {
	return math.Pow(c.radius, 2.0) * math.Pi
}

type Shape interface {
	area() float64
}

func info(s Shape) {
	fmt.Printf("area: %f\n", s.area())
}

func main() {
	ball := Circle{radius: 5}
	square := Square{length: 5, width: 5}

	info(ball)
	info(square)
}
