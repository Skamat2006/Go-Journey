package piscine

/*
func AlphaCount counts the letters of a string and returns the count.
The letters are from the Latin alphabet list only, any other characters, symbols or empty spaces shall not be counted.
*/
func AlphaCount(s string) int {
	count := 0
	letter := []byte(s)
	for i := 0; i <= len(s)-1; i++ {
		if (letter[i] >= 65 && letter[i] <= 90) || (letter[i] >= 97 && letter[i] <= 122) {
			count = count + 1
		}
	}
	return count
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := "Hello 78 World!    4455 /"
	nb := piscine.AlphaCount(s)
	fmt.Println(nb)
}
*/
