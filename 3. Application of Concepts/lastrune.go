package piscine

/*
func LastRune returns the last rune of a string.
*/
func LastRune(s string) rune {
	n := len(s)
	MyRune := []rune(s)
	return MyRune[n-1]
}

/*
package main

import (
	"github.com/01-edu/z01"

	"piscine"
)

func main() {
	z01.PrintRune(piscine.LastRune("Hello!"))
	z01.PrintRune(piscine.LastRune("Salut!"))
	z01.PrintRune(piscine.LastRune("Ola!"))
	z01.PrintRune('\n')
}
*/
