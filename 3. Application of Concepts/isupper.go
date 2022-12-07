package piscine

/*
 func IsUpper returns true if the string passed as parameter contains only uppercase characters, otherwise, returns false.
*/
func IsUpper(s string) bool {
	letter := []byte(s)
	for i := 0; i <= len(s)-1; i++ {
		if letter[i] < 65 || letter[i] > 90 {
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
	fmt.Println(piscine.IsUpper("HELLO"))
	fmt.Println(piscine.IsUpper("HELLO!"))

}
*/
