package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	check = 2
	delay = 5
)

func init() {
	name := "Nivi"
	version := 1.5
	fmt.Println("Hello,", name)
	fmt.Println("Program Version:", version)
}

func main() {
	for {
		showMenu()
		option := readOption()

		switch option {
		case 1:
			monitoring()
		case 2:
			printLogs()
		case 0:
			exitCode()
		default:
			fmt.Println("Unknown option")
			os.Exit(-1)
		}
	}
}

func showMenu() {
	fmt.Println("1 - Check Status")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit")
}

func readOption() int {
	var readOpt int
	fmt.Scan(&readOpt)
	fmt.Println("Chosen option:\n", readOpt)

	return readOpt
}

func monitoring() {
	fmt.Println("Monitoring")
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
	}
}

func readFile() []string {

	var sites []string
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			sites = append(sites, line)
		}

		if err := scanner.Err(); err != nil {
			return sites
		}
	}
	return sites
}

func siteChecker(site string) {
	answer, err := http.Get(site)

	if err != nil {
		fmt.Println("Error:", err)
	}
	if answer.StatusCode == 200 {
		fmt.Println("Site:", site, "Status Ok")
		logReg(site, true)
	} else {
		fmt.Println("Site:", site, "Not working. Status Code:", answer.StatusCode)
		logReg(site, true)
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
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	file.WriteString(time.Now().Format("01/02/2006 03:04:05PM") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func printLogs() {
	file, err := os.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))
}
