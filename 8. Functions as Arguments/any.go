package piscine

/*
func Any that returns true, for a string slice :
if, when that string slice is passed through an f function, at least one element returns true.
*/
func Any(f func(string) bool, a []string) bool {
	for _, each := range a {
		if f(each) {
			return true
		}
	}
	return false
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := piscine.Any(piscine.IsNumeric, a1)
	result2 := piscine.Any(piscine.IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
*/
/*
$ go run .
false
true
$
*/
