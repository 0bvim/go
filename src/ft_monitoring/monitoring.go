package main

import "fmt"
import "os"
// module to work with post, get. 
import "net/http"

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
	site := "https://workspaces.42sp.org.br/login"
	// function that do https requisition to site provided as argument.
	// to show the return of http.Get function, you need to put it in a variable.
	// this function have 2 returns, so you need to use two variables.
	// first return is answer from Get call and other is if have an error
	answer, err := http.Get(site)
	fmt.Println(answer, err)
	
}

func showLogs() {
	fmt.Println("Show Logs")
}

func exitCode() {
	fmt.Println("Exiting...")
	fmt.Println("Good bye!")
	os.Exit(0)
}
