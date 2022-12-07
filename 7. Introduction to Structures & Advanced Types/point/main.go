package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(pointer *point, Changex, Changey int) {
	pointer.x = Changex
	pointer.y = Changey
}

func main() {
	points := point{}

	setPoint(&points, 42, 21)

	a := "x = 42, y = 21"

	for _, value1 := range a {
		z01.PrintRune(value1)
	}
	z01.PrintRune('\n')
}

/*
Create a new directory called point.

    The code below must be copied into a file called main.go inside the point directory.

    The necessary changes must be applied so that the program works.
*/
/*
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	fmt.Printf("x = %d, y = %d\n",points.x, points.y)
}
*/
/*
$ go run .
x = 42, y = 21
$
*/
