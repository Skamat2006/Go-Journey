package piscine

/*
func ForEach: for an int slice, applies a function on each element of that slice.
*/
func ForEach(f func(int), a []int) {
	for _, element := range a {
		f(element)
	}
}

func PrintNbr(a int) {
}

/*
package main

import "piscine"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	piscine.ForEach(piscine.PrintNbr, a)
}
*/
/*
$ go run .
123456
$
*/
