package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	ptrDoor.state = CLOSE
	return true
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	ptrDoor.state = OPEN
	return true
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	ptrDoor.state = CLOSE
	return true
}

func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...")
	ptrDoor.state = CLOSE
	return true
}

func main() {
	door := &Door{}
	OpenDoor(door)
	IsDoorClose(door)
	IsDoorOpen(door)
	CloseDoor(door)
}

const (
	OPEN  = true
	CLOSE = true
)
