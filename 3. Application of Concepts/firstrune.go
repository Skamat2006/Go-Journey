package piscine

/*
func FirstRune returns the first rune of a string
*/
func FirstRune(s string) rune {
	MyRune := []rune(s)
	return MyRune[0]
}

/*
package main

import (
	"github.com/01-edu/z01"

	"piscine"
)

func main() {
	z01.PrintRune(piscine.FirstRune("Hello!"))
	z01.PrintRune(piscine.FirstRune("Salut!"))
	z01.PrintRune(piscine.FirstRune("Ola!"))
	z01.PrintRune('\n')
}
*/
