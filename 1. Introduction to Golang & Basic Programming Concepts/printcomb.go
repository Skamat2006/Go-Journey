package piscine

import "github.com/01-edu/z01"

/*
func PrintComb: prints, in ascending order and on a single line: all unique combinations of three different digits so that, the first digit is lower than the second, and the second is lower than the third.
These combinations are separated by a comma and a space.
*/

func PrintComb() {
	for a := '0'; a <= '7'; a++ {
		for b := a + 1; b <= '8'; b++ {
			for c := b + 1; c <= '9'; c++ {
				if a < b && b < c {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(c)
					if a == '7' && b == '8' && c == '9' {
						z01.PrintRune(10) // ascii value = /n
					} else {
						z01.PrintRune(44) // ascii value = ','
						z01.PrintRune(32) // ascii value = ' '
					}
				}
			}
		}
	}
}

/* Test
package main

import "piscine"

func main() {
	piscine.PrintComb()
}
*/
