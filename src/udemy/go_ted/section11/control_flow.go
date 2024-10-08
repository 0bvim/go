package main

import "fmt"

// digital ocean article talking about init function
// https://www.digitalocean.com/community/tutorials/understanding-init-in-go

func init() {
	fmt.Println("This function run before all things in main")
}

func main() {
	fmt.Println("Hello")
}
