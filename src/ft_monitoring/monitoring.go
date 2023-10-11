package main

import (
	"fmt"
	"os"
	// module to work with post, get. 
	"net/http"
	// time module
	"time"
	"io"
	// to open file and read all in a row, is a good choice.
	//another module to read file
	"bufio"
	// module to use string function
	"strings"
	// module to convert from many kind of types to string
	"strconv"
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
	version := 1.3
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
	// this slice (sites) are commented because now I'll implement a slice to read a file
	// sites := []string{"https://intra.42.fr/","https://google.com/","https://find-peers.codam.nl/São-Paulo","https://game.42sp.org.br/", "https://workspaces.42sp.org.br/login"}
	// use range in for instead old for i := 0; i < x; i++

	sites := readFile()
	for i := 0; i < check; i++ {

		for i, site := range sites {
			fmt.Println("Checking Status", i, ":", site)
			siteChecker(site)
			fmt.Println("")
		
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
}

func	readFile() []string {

	var sites []string
	// to read a file have two function, from 'os' module and from 'io/util' module
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	// io/util this function return an array of bytes
	//file, err := ioutil.ReadFile("sites.txt")
	// to print it, just print like a string
	// handling errors
	//another function to read file line per line
	
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		// if EOL if a \n (empty line) program crash cuz try to do a http call to a incorrect link
		if line == "" {
			break
		}
		// use string module to trim lines	
		line = strings.TrimSpace(line)
		// append read/trimmed content to site slice
		sites = append(sites, line)
		// io.EOF is the function that detect when file ends.
		if err == io.EOF {
			break
		}
		// string(file) is like a casting in clang
	}
	// now we need to colose fd from read
	file.Close()
	return sites
}

func siteChecker(site string) {
	answer, err := http.Get(site)

	if err != nil {
		fmt.Println("Error:", err)
	}
	if answer.StatusCode == 200 {
		fmt.Println("Site:", site, "Status Ok")
		logReg (site, true)
	} else {
		fmt.Println("Site:", site, "Not working. Status Code:", answer.StatusCode)
		logReg (site, true)
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

func logReg(site string, status bool) {
	// opening file with OpenFile, with this function you can create a file if it not exist
	// O_RDWR to read and write in this file, O_APPEND to keep all lines, O_CREATE to creat this file if not exist 0666 is the permission of file
	file, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	
	// if error return err printed in stdou
	if err != nil {
		fmt.Println(err)
	}
	// this function is to write a string inside of a file
	// need to import module strconv to use function to convert bool to string
	file.WriteString(site + "- online: " + strconv.FormatBool(status) + "\n")
	// if not error, print content of file
	file.Close()
}
