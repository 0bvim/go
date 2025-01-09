package main

import "fmt"

func example() {
  for i := 0; i < 4; i++ {
    defer fmt.Print(i)
  }
}

func main() {
  example()
  
  {
    defer fmt.Println("Bye 5") 
  }
  defer fmt.Println("Bye 4") 
  defer fmt.Println("Bye 3") 
  defer fmt.Println("Bye 2") 
  defer fmt.Println("Bye 1") 
  fmt.Println("Bye 0")
}
