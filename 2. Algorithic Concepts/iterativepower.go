package piscine

/*
func IterartivePower returns the value of nb to the power of power.
Negative powers will return 0. Overflows do not have to be dealt with.
*/
func IterativePower(base int, expo int) int {
	if expo < 0 {
		return 0
	} else if expo == 0 {
		return 1
	}
	val := base
	for i := 1; i < expo; i++ {
		base = base * val
	}
	return base
}

/*
package main

import (
	"fmt"
	"piscine"
)


func main() {
	fmt.Println(piscine.IterativePower(4, 3))
}
*/
