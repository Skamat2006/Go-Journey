package piscine

/*
func Index behaves like the Index function.
*/
func Index(s string, toFind string) int {
	str := []rune(s)
	substr := []rune(toFind)
	strcounter := 0
	substrcounter := 0

	for range substr {
		substrcounter++
	}
	if substrcounter == 0 {
		return 0
	}

	for range str {
		strcounter++
	}

	for index, letter := range str {
		if letter == substr[0] && strcounter >= substrcounter+index-1 {
			a := 1
			for b := 1; b < substrcounter; b++ {
				if substr[b] == str[index+b] {
					a++
				}
			}
			if a == substrcounter {
				return index
			}
		}
	}
	return -1
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Index("Hello!", "l"))
	fmt.Println(piscine.Index("Salut!", "alu"))
	fmt.Println(piscine.Index("Ola!", "hOl"))
}
*/
