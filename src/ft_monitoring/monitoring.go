package main

import "fmt"
import "os"

func main(){
	showIntro()
	showMenu()
	option := readOption()

	switch option {
	case 1:
		fmt.Println("Monitoring")
	case 2:
		fmt.Println("Show Logs")
	case 0:
		fmt.Println("Exit")
		fmt.Println("Good bye!")
		os.Exit(0)
	default:
		fmt.Println("Unknown option")
		os.Exit(-1)
	}
}

func showIntro() {
	name := "Nivi"
	version := 1.0
	fmt.Println("Hello,", name)
	fmt.Println("Program Version:", version)
}

func showMenu() {
	fmt.Println("1- Check Status")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")
}

func readOption() int {
	var readOpt int
	fmt.Scan(&readOpt)
	fmt.Println("Chosen option:", readOpt)

	return readOpt
}
