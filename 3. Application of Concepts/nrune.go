package piscine

/*
func NRune returns the nth rune of a string. If not possible, it returns 0.
*/
func NRune(s string, n int) rune {
	a := []rune(s)
	counter := 0
	findindex := 0

	for range s {
		counter++
	}
	if n > counter {
		return 0
	} else if n <= 0 {
		return 0
	}
	for i := 0; i < n; i++ {
		findindex = i
	}
	return a[findindex]
}

/*
package main

import (
	"github.com/01-edu/z01"

	"piscine"
)

func main() {
	z01.PrintRune(piscine.NRune("Hello!", 3))
	z01.PrintRune(piscine.NRune("Salut!", 2))
	z01.PrintRune(piscine.NRune("Bye!", -1))
	z01.PrintRune(piscine.NRune("Bye!", 5))
	z01.PrintRune(piscine.NRune("Ola!", 4))
	z01.PrintRune('\n')
}
*/
