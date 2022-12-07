package piscine

import "github.com/01-edu/z01"

/*
func PrintNbrInOrder prints the digits of an int passed in parameter in ascending order. All possible values of type int have to go through, excluding negative numbers. Conversion to int64 is not allowed.
*/
func PrintNbrInOrder(n int) {
	my_arr := [10]int{}

	if n < 10 && n >= 0 {
		z01.PrintRune(rune(n + 48))
	} else {

		for {
			if n == 0 {
				break
			}
			my_arr[n%10]++
			n /= 10
		}

		for index, num := range my_arr {
			if num != 0 {
				for i := 0; i < num; i++ {
					z01.PrintRune(rune(index + 48))
				}
			}
		}
	}
}

/*
package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n >= 0 {
		var array []int
		var mod int
		for i := 0; i < 20; i++ {
			if n > 9 {
				mod = n % 10
				array = append(array, mod)
				n = n / 10
			}
		}
		array = append(array, n)
		for i := 0; i < len(array)-1; i++ {
			for j := 0; j < len(array)-i-1; j++ {
				if array[j] > array[j+1] {
					array[j], array[j+1] = array[j+1], array[j]
				}
			}
		}
		for _, value := range array {
			z01.PrintRune(rune(value + '0'))
		}
	}
}
*/

/*
package main

import "piscine"

func main() {
	piscine.PrintNbrInOrder(321)
	piscine.PrintNbrInOrder(0)
	piscine.PrintNbrInOrder(321)
}
*/
