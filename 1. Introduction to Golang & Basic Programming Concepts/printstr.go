package piscine

import "github.com/01-edu/z01"

/*
func PrintStr prints one by one the characters of a string on the screen.
*/
func PrintStr(s string) {
	for _, word := range s {
		z01.PrintRune(word)
	}
}

/*
package main

import "piscine"

func main() {
	piscine.PrintStr("Hello World!")
}
*/
