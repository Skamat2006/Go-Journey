package piscine

/*
   func TrimAtoi transforms numbers within a string, into an int.
   If the - sign is encountered before any number it should determine the sign of the returned int.
   This function should only return an int. In the case of an invalid input, the function should return 0.
   Note: There will never be more than one sign in a string in the tests.
*/
func TrimAtoi(s string) int {
	strrune := []rune(s)
	var ints []rune
	bool1 := false
	bool2 := false
	n := 0
	for y := 0; y < len(strrune); y++ {
		if strrune[y] == '-' && !bool2 {
			bool1 = true
		}
		if strrune[y] >= '0' && strrune[y] <= '9' {
			if !bool1 {
				bool2 = true
			}
		}

	}
	for j := 0; j < len(strrune); j++ {
		if strrune[j] >= '0' && strrune[j] <= '9' {
			ints = append(ints, strrune[j])
		}
	}
	for i := 0; i < len(ints); i++ {
		y := ints[i] - '0'
		n = n*10 + int(y)
	}
	if bool1 {
		n = -n
	}
	return n
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.TrimAtoi("12345"))
	fmt.Println(piscine.TrimAtoi("str123ing45"))
	fmt.Println(piscine.TrimAtoi("012 345"))
	fmt.Println(piscine.TrimAtoi("Hello World!"))
	fmt.Println(piscine.TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(piscine.TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(piscine.TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(piscine.TrimAtoi("sdx1+fa2W3s4"))
}
*/
