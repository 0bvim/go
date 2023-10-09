package main

import "fmt"
import "os"

func main(){
	showIntro()
	showMenu()
	option := readOption()

	switch option {
	case 1:
		monitoring()
	case 2:
		showLogs()
	case 0:
		exitCode()
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

func monitoring() {
	fmt.Println("Monitoring")
}

func showLogs() {
	fmt.Println("Show Logs")
}

func exitCode() {
	fmt.Println("Exiting...")
	fmt.Println("Good bye!")
	os.Exit(0)
}
