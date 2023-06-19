package algorithms

func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if arr[minIndex] != arr[i] {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}

	}
}
