package algorithms

import "fmt"

func Quicksort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var less, greater []int

	////SOME
	////CODE

	for _, el := range arr[1:] {
		if el > pivot {
			greater = append(greater, el)
		} else {
			less = append(less, el)
			fmt.Println("TEST")
		}
	}

	less = Quicksort(less)
	greater = Quicksort(greater)

	return append(append(less, pivot), greater...)
}
