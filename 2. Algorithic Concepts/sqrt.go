package piscine

/*
func Sqrt returns the square root of the int passed as parameter, if that square root is a whole number. Otherwise it returns 0.
*/
func Sqrt(number int) int {
	var root float32 = float32(number) / 2
	var temp, temp2 float32

	for {
		temp = root
		root = (temp + (float32(number) / temp)) / 2
		if (temp - root) == 0 {
			break
		}
	}
	temp2 = float32(int(root))
	if root-temp2 > 0 {
		return 0
	} else {
		return int(root)
	}
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Sqrt(4))
	fmt.Println(piscine.Sqrt(3))
}
*/
