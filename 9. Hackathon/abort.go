package piscine

// func Abort returns the median of five int arguments
func Abort(a, b, c, d, e int) int {
	array := []int{a, b, c, d, e}
	count := 5
	temp := 0
	for i := 0; i < count-1; i++ {
		temp = i
		for j := i + 1; j < count; j++ {
			if array[j] < array[temp] {
				temp = j
			}
		}
		array[i], array[temp] = array[temp], array[i]
	}
	return array[2]
}

/*
import "fmt"
	   "piscine"

func main(){
	middle := piscine.Abort(2, 3, 8, 7, 5)
	fmt.Println(middle)
} /*
/*
func Abort(a, b, c, d, e int) int {
	array := []int{a, b, c, d, e}
	length := 0
	for l := range array {
		length = l + 1
	}
	for i, j := 0, 1; j < length; i, j = i+1, j+1 {
		if array[i] > array[j] {
			temp := array[i]
			array[i] = array[j]
			array[j] = temp
		}
	}
	return array[2]
}
*/
