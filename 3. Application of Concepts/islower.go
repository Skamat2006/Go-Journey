package piscine

/*
func IsLower returns true if the string passed as the parameter contains only lowercase characters, otherwise, returns false.
*/
func IsLower(s string) bool {
	letter := []byte(s)
	for i := 0; i <= len(s)-1; i++ {
		if letter[i] < 97 || letter[i] > 122 {
			return false
		}
	}
	return true
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsLower("hello"))
	fmt.Println(piscine.IsLower("hello!"))

}
*/
