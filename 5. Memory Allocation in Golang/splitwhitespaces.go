package piscine

/*
func SplitWhiteSpaces separates the words of a string and puts them in a string slice.
The separators are spaces, tabs and newlines.
*/
func SplitWhiteSpaces(s string) []string {
	result := []string{}
	word := ""
	prev_space := false
	for _, element := range s {
		if element == 32 || element == 9 || element == '\n' {
			if !prev_space {
				result = append(result, word)
				prev_space = true
			}
			word = ""
		} else {
			word = word + string(element)
			prev_space = false
		}
	}
	if len(word) > 0 {
		result = append(result, word)
	}
	return result
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello how are you?"))
}
*/
