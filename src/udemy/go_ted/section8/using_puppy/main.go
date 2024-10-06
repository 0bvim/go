package main

import (
	"fmt"

	"github.com/0bvim/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()

	fmt.Println(s1)
	fmt.Println(s2)

	// also like this
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())

	// go mod depend 2
	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()

	fmt.Println(s3)
	fmt.Println(s4)
}
