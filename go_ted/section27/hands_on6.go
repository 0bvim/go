package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go Arch:", runtime.GOARCH)
	fmt.Println("Go Os:", runtime.GOOS)
	fmt.Println("Go Root:", runtime.GOROOT())
}
