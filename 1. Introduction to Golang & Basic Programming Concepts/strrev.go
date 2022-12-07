package piscine

/*
   func StrRev reverses a string and return the reversed string.
*/
func StrRev(s string) string {
	r := []rune(s)
	length := len(r) - 1
	var newr []rune
	for i := length; i >= 0; i-- {
		newr = append(newr, r[i])
	}
	return string(newr)
}

/*
func StrRev(s string) string {
	// convert string to byte string
	chars := []rune(s)
	// start iterating the byte string
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		// swap the first element with the last element of converted byte string
		chars[i], chars[j] = chars[j], chars[i]
	}
	// convert the byte string to string and return it
	return string(chars)
}
*/

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := "Hello World!"
	s = piscine.StrRev(s)
	fmt.Println(s)
}
*/
