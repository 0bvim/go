package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	xb, err := readFile("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(xb)
	fmt.Println(string(xb))
}

func readFile(fileName string) ([]byte, error) {
	xb, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return xb, nil
}
