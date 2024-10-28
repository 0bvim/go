package main

import (
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.first))
	return err
}
func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer file.Close()
}
