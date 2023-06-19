package algorithms

import "fmt"

// Функция для слияния двух отсортированных массивов
func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0

	// Пока есть элементы в обоих массивах
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	// Добавляем оставшиеся элементы
	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}

	return result
}

// Функция для сортировки слиянием
func MergeSort(arr []int) []int {
	fmt.Printf("MyArr = %v\n", arr)
	if len(arr) <= 1 {
		fmt.Printf("\nreturn from MergeSort with arr = %d\n", arr)
		return arr
	}

	// Находим средний индекс массива
	middle := len(arr) / 2

	fmt.Printf("middle = %d\n", middle)

	// Рекурсивно сортируем левую и правую половины
	fmt.Printf("leftArr = %v\n", arr[:middle])
	left := MergeSort(arr[:middle])
	fmt.Printf("left = %v\n", left)

	fmt.Printf("\nMyArr = %v, middle = %v\n", arr, middle)
	fmt.Printf("rightArr = %v\n", arr[middle:])
	right := MergeSort(arr[middle:])
	fmt.Printf("right = %v\n", right)

	// Сливаем отсортированные половины

	fmt.Printf("\n\nLeftArr = %v, RightArr = %v\n", left, right)
	fmt.Printf("\n\nAfter merge = %v\n", merge(left, right))
	return merge(left, right)
}
