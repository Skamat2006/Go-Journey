package piscine

/*
func PointOne takes a pointer to an int as argument and gives to this int the value of 1.
*/
func PointOne(n *int) {
	*n = 1
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	n := 0
	piscine.PointOne(&n)
	fmt.Println(n)
}
*/
