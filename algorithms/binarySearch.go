package algorithms

func BinarySearch(arr []int, searchElement int) int {
	lowIndex := 0
	highIndex := len(arr) - 1

	for highIndex >= lowIndex {
		middle := (highIndex + lowIndex) / 2
		if searchElement == arr[middle] {
			return middle
		} else if searchElement < arr[middle] {
			highIndex = middle - 1
		} else {
			lowIndex = middle + 1
		}
	}
	return -1
}
