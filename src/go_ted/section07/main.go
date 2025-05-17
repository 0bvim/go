package main

// import multiple packages
import (
	"fmt"
	"runtime"
)

// get os and arch using runtime package
func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
