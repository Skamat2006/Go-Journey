package main

import (
	"os"

	"github.com/01-edu/z01"
)

/*
Write a program that prints the arguments received in the command line in reverse order.
*/
func main() {
	arguments := os.Args
	counter := 0

	for i := range arguments {
		counter = i
	}
	for i := counter; i > 0; i-- {
		for _, v := range arguments[i] {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}

/*
$ go run . choumi is the best cat
cat
best
the
is
choumi
$
*/
