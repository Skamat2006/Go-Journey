package piscine

/*
func IsNumeric returns true if the string passed as a parameter contains only numerical characters, otherwise, returns false.
*/
func IsNumeric(s string) bool {
	runes := []byte(s)
	yes := true
	for i := range s {
		if runes[i] >= '0' && runes[i] <= '9' {
			yes = true
		} else {
			return false
		}
	}
	return yes
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsNumeric("010203"))
	fmt.Println(piscine.IsNumeric("01,02,03"))
}
*/
