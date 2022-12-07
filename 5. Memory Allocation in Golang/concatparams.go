package piscine

/*
func ConcatParams takes the arguments received in parameters and returns them as a string. The string is the result of all the arguments concatenated with a newline (\n) between.
*/
func ConcatParams(args []string) string {
	str := ""
	for i, rep := range args {
		str += string(rep)
		if i != len(args)-1 {
			str += "\n"
		}
	}
	return str
}

/*
package main

import (
	"fmt"
	"piscine"
)

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(piscine.ConcatParams(test))
}
*/
