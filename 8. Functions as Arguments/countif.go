package piscine

/*
func CountIf returns, the number of elements of a string slice, for which the f function returns true.
*/
func CountIf(f func(string) bool, tab []string) int {
	total := 0
	for _, each := range tab {
		if f(each) {
			total++
		}
	}
	return total
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This","1", "is", "4", "you"}
	answer1 := piscine.CountIf(piscine.IsNumeric, tab1)
	answer2 := piscine.CountIf(piscine.IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}
*/
