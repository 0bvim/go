package main

import "fmt"

func main() {

	mapping := map[string]int{
		"Todd":  42,
		"Lila":  22,
		"Stete": 32,
	}
	// 116 - map
	{
		fmt.Printf("mapping: %v\n", mapping)
		fmt.Printf("mapping Lila: %v\n", mapping["Lila"])

		another := make(map[string]int)
		another["Lucas"] = 28
		another["Steph"] = 37
		fmt.Printf("another: %v\n", another)
		fmt.Println("--------------")
	}
	// 117 adding element in a map and using range
	{
		for key, value := range mapping {
			fmt.Println(key, value)
		}

		for _, value := range mapping {
			fmt.Println(value)
		}

		for key, _ := range mapping {
			fmt.Println(key)
		}
		fmt.Println("--------------")
	}
	// 118 delete item from a map
	{
		delete(mapping, "Todd")
		delete(mapping, "Todd") // using twice don't make program panic, just ignore even if you try to print or something else.

		fmt.Printf("mapping: %v\n", mapping["Todd"])
		fmt.Printf("mapping: %v\n", mapping)
		fmt.Println("--------------")
	}
	// 118 using ok notation to verify valid values in map
	{
		// if you liik up a non-existent key, the zero value will be returned as the value associated
		// with that non-existent key
		if v, ok := mapping["Lila"]; !ok {
			fmt.Println("key didn't exist")
		} else {
			fmt.Println(v)
		}

		delete(mapping, "Lila")
		if v, ok := mapping["Lila"]; ok {
			fmt.Println(v)
		} else {
			fmt.Println("key didn't exist")
		}
	}
}
