package piscine

/*  func CollatzCountdown returns the number of steps necessary to reach 1 using the collatz countdown.
    It must return -1 if start is equal to 0 or negative.
	**Collatz conjecture applies to all positive integers and speculates that it is always possible to get "back to 1" if you follow these steps:
	- if n is 1 stop
	- otherwise if n is even, repeat this process on n/2
	- otherwise if n is odd, repeat this process on 3n+1
*/

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	steps := 0

	for start != 1 {
		// for even numbers
		if start%2 == 0 {
			start /= 2
		} else {
			// for odd numbers
			start = 3*start + 1
		}
		steps++
	}
	return steps
}

/*
import (
	"fmt"
	"piscine"
)

func main() {
	steps := piscine.CollatzCountdown(12)
	fmt.Println(steps)
}
*/
