package main

import "fmt"

func main() {
	// slices in slices
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "Moneypenny", "Guiness", "Wolverine Tracks"}
	fmt.Printf("jb: %v\n", jb)
	fmt.Printf("jm: %v\n", jm)

	xp := [][]string{jb, jm}
	xxp := [][]string{jb, jm}
	xxxp := [][][]string{xp, xxp}
	fmt.Printf("xxp: %v\n", xxp)
	fmt.Printf("xp: %v\n", xp)
	fmt.Printf("xxxp: %v\n", xxxp)

}
