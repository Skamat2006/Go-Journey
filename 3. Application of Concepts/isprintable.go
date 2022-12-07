package piscine

/*
 func IsPrintable returns true if the string passed as a parameter contains only printable characters, otherwise, returns false.
*/
func IsPrintable(s string) bool {
	runes := []byte(s)
	yes := true
	for i := range s {
		if runes[i] >= 32 && runes[i] <= 127 {
			yes = true
		} else {
			return false
		}
	}
	return yes
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsPrintable("Hello"))
	fmt.Println(piscine.IsPrintable("Hello\n"))

}
*/
