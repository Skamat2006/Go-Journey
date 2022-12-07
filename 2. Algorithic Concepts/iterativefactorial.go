package piscine

/*
func IterativeFactorial returns the factorial of the int passed as parameter.
Errors (non possible values or overflows) will return 0.
*/
func IterativeFactorial(nb int) int {
	if nb > 25 || nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	}
	x := 1
	for i := 1; i <= nb; i++ {
		x = x * (i)
	}
	return x
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	arg := 4
	fmt.Println(piscine.IterativeFactorial(arg))
}
*/
