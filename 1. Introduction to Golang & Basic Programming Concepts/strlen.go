package piscine

/*
func StrLen counts the runes of a string and that returns that count.
*/
func StrLen(s string) int {
	r := []rune(s)
	amount := len(r)
	return amount
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	l := piscine.StrLen("Hello World!")
	fmt.Println(l)
}
*/
