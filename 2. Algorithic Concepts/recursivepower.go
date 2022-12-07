package piscine

/*
func RecursivePower returns the value of nb to the power of power.
Negative powers will return 0. Overflows do not have to be dealt with.
for is forbidden for this exercise.
*/
func RecursivePower(base int, expo int) int {
	if expo == 0 {
		return 1
	} else {
		return base * RecursivePower(base, expo-1)
	}
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.RecursivePower(4, 3))
}
*/
