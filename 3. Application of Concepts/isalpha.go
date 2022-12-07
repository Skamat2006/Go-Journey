package piscine

/*
func IsAlpha returns true if the string passed as the parameter only contains alphanumerical characters or is empty, otherwise, and returns false.
*/
func IsAlpha(s string) bool {
	runes := []byte(s)
	yes := true
	for i := range s {
		if (runes[i] >= 'A' && runes[i] <= 'Z') || (runes[i] >= 'a' && runes[i] <= 'z') || (runes[i] >= '0' && runes[i] <= '9') {
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
	fmt.Println(piscine.IsAlpha("Hello! How are you?"))
	fmt.Println(piscine.IsAlpha("HelloHowareyou"))
	fmt.Println(piscine.IsAlpha("What's this 4?"))
	fmt.Println(piscine.IsAlpha("Whatsthis4"))

}
*/
