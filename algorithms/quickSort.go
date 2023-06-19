package algorithms

import "fmt"

func Quicksort(arr []int) []int {

	if len(arr) < 2 {
		//workTwo
		return arr
	}
	pivot := arr[0]
	var less, greater []int

	for _, el := range arr[1:] {
		if el > pivot {
			greater = append(greater, el)
		} else {
			less = append(less, el)
			fmt.Println("Test00")
			fmt.Println("Test")
		}
	}

	less = Quicksort(less)
	greater = Quicksort(greater)

	return append(append(less, pivot), greater...)
}
