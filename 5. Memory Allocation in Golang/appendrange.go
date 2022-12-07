package piscine

/*
func AppendRange takes an int min and an int max as parameters. The function should return a slice of ints with all the values between the min and max.
Min is included, and max is excluded.
If min is greater than or equal to max, a nil slice is returned.
make is not allowed for this exercise.
*/
func AppendRange(min, max int) []int {
	if min >= max {
		var tab []int = nil
		return tab
	} else {
		tab := []int{}
		for i := min; i < max; i++ {
			tab = append(tab, i)
		}
		return tab
	}
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.AppendRange(5, 10))
	fmt.Println(piscine.AppendRange(10, 5))
}
*/
