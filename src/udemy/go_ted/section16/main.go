package main

import "fmt"

func main() {
	favorites := make(map[string][]string)

	favorites["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	favorites["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	favorites["no_dr"] = []string{"cats", "ice cream", "sunsets"}

	for key, value := range favorites {
		fmt.Println("Record for", key)
		for i, value := range value {
			fmt.Printf("%v - %v\n", i, value)
		}
	}
	fmt.Println("-------------------------")

	delete(favorites, "bond_james")

	for key, value := range favorites {
		fmt.Println("Record for", key)
		for i, value := range value {
			fmt.Printf("%v - %v\n", i, value)
		}
	}
	fmt.Println("-------------------------")

}
