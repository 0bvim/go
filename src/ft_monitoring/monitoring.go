package main

import (
	"fmt"
	"os"
	// module to work with post, get. 
	"net/http"
	// time module
	"time"
)
//creating constant variable in go

const	check = 5
const	delay = 5

func main(){
	showIntro()
	// in golang don't have while loop, only for. 
	// if you want a infinity loop, you need to use for withou conditionals
	for {

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
	fmt.Println("")

	return readOpt
}

func monitoring() {
	fmt.Println("Monitoring")
	sites := []string{"https://intra.42.fr/","https://google.com/","https://find-peers.codam.nl/São-Paulo","https://game.42sp.org.br/", "https://workspaces.42sp.org.br/login"}
	// use range in for instead old for i := 0; i < x; i++
	for i := 0; i < check; i++ {

		for i, site := range sites {
			fmt.Println("Checking Status", i, ":", site)
			siteChecker(site)
			fmt.Println("")
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
	//site := "https://google.com"
	// to verify more than one website we need to create and array of string
	// all arrays in goland have a fixed value
	// var sites [4]string
	// site[0] = "https://intra.42.fr/"
	// site[1] = "https://42evaluators.com/"
	// site[2] = "https://find-peers.codam.nl/São-Paulo"
	// site[3] = "https://game.42sp.org.br/"
	// site[4] = "https://workspaces.42sp.org.br/login"
	// in golang we have another structure that is better than array
	// it's called slice, it works like an array but don't need a fixed value
	
	
	// function that do https requisition to site provided as argument.
	// to show the return of http.Get function, you need to put it in a variable.
	// this function have 2 returns, so you need to use two variables.
	// first return is answer from Get call and other is if have an error
}

func siteChecker(site string) {
	answer, _ := http.Get(site)

	if answer.StatusCode == 200 {
		fmt.Println("Site:", site, "Status Ok")
	} else {
		fmt.Println("Site:", site, "Not working. Status Code:", answer.StatusCode)
	}
}
func showLogs() {
	fmt.Println("Show Logs")
}

func exitCode() {
	fmt.Println("Exiting...")
	fmt.Println("Good bye!")
	os.Exit(0)
}
