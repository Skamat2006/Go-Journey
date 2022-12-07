package main

import (
	"fmt"
	"os"
)

/*
Write a program comcheck that displays on the standard output Alert!!! followed by newline ('\n')
if at least one of the arguments passed in parameter matches the string:
    01, galaxy or galaxy 01.
    If none of the parameters match, the program displays nothing.
*/

func main() {
	words := []string{"01", "galaxy", "galaxy 01"}
	count := 0
	arguments := os.Args[1:]

	for i := range arguments {
		for _, s := range words {
			if arguments[i] == s {
				count++
			}
		}
	}
	if count >= 1 {
		fmt.Println("Alert!!!")
	}
}
