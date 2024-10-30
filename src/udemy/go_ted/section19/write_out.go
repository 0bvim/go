package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) {
	_, err := w.Write([]byte(p.first))
	if err != nil {
		return
	}
}

func main() {
	p := person{
		first: "Jenny",
	}

	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var b bytes.Buffer

	p.writeOut(file)
	p.writeOut(&b)
	fmt.Println(b.String())
}
