package piscine

/*
func Compare behaves like the Compare function.
*/
func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Compare("Hello!", "Hello!"))
	fmt.Println(piscine.Compare("Salut!", "lut!"))
	fmt.Println(piscine.Compare("Ola!", "Ol"))
}
*/
