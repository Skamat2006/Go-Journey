package piscine

/*
func RecursiveFactorial returns the factorial of the int passed as parameter.
Errors (non possible values or overflows) will return 0.
for is forbidden for this exercise.
*/
func RecursiveFactorial(x int) int {
	if x > 0 {
		return x * RecursiveFactorial(x-1)
	} else {
		return 1
	}
	return 0
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	arg := 4
	fmt.Println(piscine.RecursiveFactorial(arg))
}
*/
