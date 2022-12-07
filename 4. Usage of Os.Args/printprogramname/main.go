package main

import (
	"os"

	"github.com/01-edu/z01"
)

/*
Write a program that prints the name of the program.
*/
func main() {
	arguments := os.Args[0]
	aR := []rune(arguments)

	for i, v := range aR {
		if i > 1 {
			z01.PrintRune(rune(v))
		}
	}
	z01.PrintRune('\n')
}

/*
student/piscine/printprogramname$ go build main.go
student/piscine/printprogramnane$ ./main
main
student/piscine/printprogramname$ go build
student/piscine/printprogramname$ ./printprogramname | cat -e
printprogramname$
student/piscine/printprogramname$ go build -o Nessy
student/piscine/printprogramname$ ./Nessy
Nessy
student/piscine/printprogramname$
*/
