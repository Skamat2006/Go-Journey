package piscine

/*
   func BasicAtoi simulates the behaviour of the Atoi function in Go. Atoi transforms a number defined as a string in a number defined as an int.

   Atoi returns 0 if the string is not considered as a valid number. For this exercise only valid string will be tested. They will only contain one or several digits as characters.

   For this exercise the handling of the signs + or - does not have to be taken into account.

   This function will only have to return the int.

   For this exercise the error return of Atoi is not required.
*/
func BasicAtoi(s string) int {
	var newSliceOfString []rune
	var newInt int
	for _, character := range s {
		newSliceOfString = append(newSliceOfString, character-'0')
	}
	for _, char := range newSliceOfString {
		newInt = newInt*10 + int(char)
	}
	return newInt
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.BasicAtoi("12345"))
	fmt.Println(piscine.BasicAtoi("0000000012345"))
	fmt.Println(piscine.BasicAtoi("000000"))
}
*/
