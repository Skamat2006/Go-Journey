package piscine

/*
func ToUpper capitalizes each letter in a string.
*/
func ToUpper(s string) string {
	x := []rune(s)

	for i := 0; i < len(s); i++ {
		if x[i] >= 97 && x[i] <= 122 {
			x[i] = x[i] - 32
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
	fmt.Println(piscine.ToUpper("Hello! How are you?"))
}
*/
