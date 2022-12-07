package piscine

/*
   func UltimateDivMod will divide the int a and b.
   The result of this division will be stored in the int pointed by a.
   The remainder of this division will be stored in the int pointed by b.
*/
func UltimateDivMod(a *int, b *int) {
	quotientAB := *a / *b
	remainderAB := *a % *b
	*a = quotientAB
	*b = remainderAB
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := 13
	b := 2
	piscine.UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
*/
