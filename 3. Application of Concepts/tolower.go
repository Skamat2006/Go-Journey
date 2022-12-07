package piscine

/*
func ToLower lower cases for each letter in a string.
*/
func ToLower(s string) string {
	x := []rune(s)

	for i := 0; i < len(s); i++ {
		if x[i] >= 65 && x[i] <= 90 {
			x[i] = x[i] + 32
		}
	}
	return string(x)
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.ToLower("Hello! How are you?"))
}
*/
