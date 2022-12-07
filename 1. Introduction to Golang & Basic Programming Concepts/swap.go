package piscine

/*
func Swap takes two pointers to an int (*int) and swaps their contents.
*/
func Swap(a *int, b *int) {
	c := *b
	*b = *a
	*a = c
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := 0
	b := 1
	piscine.Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
*/
