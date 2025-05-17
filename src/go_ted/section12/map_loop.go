package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
		"Lila":       4,
		"Stete":      24,
		"Nivi":       18,
	}

	for i, v := range m {
		fmt.Printf("Name %v\tAge %v\n", i, v)
	}

	fmt.Println()

	if age, ok := m["James"]; ok {
		fmt.Println("The age is", age)
	} else {
		fmt.Println("Not match")
	}

	if age, ok := m["O"]; ok {
		fmt.Println("The age is", age)
	} else {
		fmt.Println("Not match")
	}

	if age, ok := m["n"]; !ok {
		fmt.Println("Not match")
	} else {
		fmt.Println("The age is", age)
	}
}
