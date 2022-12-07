package piscine

/*
func Fibonacci is a recursive function that returns the value at the position index in the fibonacci sequence.
The first value is at index 0.
The sequence starts this way: 0, 1, 1, 2, 3 etc...
A negative index will return -1.
for is forbidden for this exercise.
*/
func Fibonacci(x int) int {
	if x < 0 {
		return -1
	} else if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else {
		return Fibonacci(x-1) + Fibonacci(x-2)
	}
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	arg1 := 4
	fmt.Println(piscine.Fibonacci(arg1))
}
*/
