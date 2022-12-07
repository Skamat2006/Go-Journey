package piscine

/*
func IsSorted returns true, if the slice of int is sorted, otherwise returns false.
The function passed in the parameter returns a positive int if a (the first argument) is greater than to b (the second argument), it returns 0 if they are equal and it returns a negative int otherwise.
To do your testing you have to write your own f function.
*/
func IsSorted(f func(a, b int) int, tab []int) bool {
	length := 0
	for i := range tab {
		length = i + 1
	}
	left := false
	right := false
	for i := 0; i < length-1; i++ {
		if f(tab[i], tab[i+1]) > 0 {
			left = true
		} else if f(tab[i], tab[i+1]) < 0 {
			right = true
		}
	}
	if left && right {
		return false
	} else {
		return true
	}
}

// package piscine

// func f(a, b int) int {
// 	return a - b
// }

// func IsSorted(f func(a, b int) int, a []int) bool {
// 	for i := 0; i < len(a)-2; i++ {
// 		if f(a[i], a[i+1])*f(a[i+1], a[i+2]) < 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := piscine.IsSorted(f, a1)
	result2 := piscine.IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
*/
