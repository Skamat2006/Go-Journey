package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileName := "quest8.txt"

	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	in, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(string(in))
}

/*
Write a program that displays, on the standard output, the content of a file given as argument.
*/

/*
$ go run .
File name missing
$ echo "Almost there!!" > quest8.txt
$ go run . quest8.txt main.go
Too many arguments
$ go run . quest8.txt
Almost there!!
$
*/
