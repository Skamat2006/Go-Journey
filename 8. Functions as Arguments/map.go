package piscine

/*
func Map: for an int slice, applies a function of this type func(int) bool on each element of that slice and returns a slice of all the return values.
*/
func Map(f func(int) bool, a []int) (result []bool) {
	for _, el := range a {
		result = append(result, f(el))
	}
	return
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := piscine.Map(piscine.IsPrime, a)
	fmt.Println(result)
}
*/
/*
$ go run .
[false true true false true false]
$
*/
