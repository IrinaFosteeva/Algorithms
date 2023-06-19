package main

import "Algorithms/algorithms"

func main() {
	//mySlice := []int{6, 10, 1, 6, 4, 1, 8}
	//fmt.Println(algorithms.Quicksort(mySlice))
	//fmt.Println(algorithms.MergeSort(mySlice))

	//algorithms.SelectionSort(mySlice)
	//fmt.Println(mySlice)

	//mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//searchedElement := 20
	//fmt.Printf("Searched index for number %d is %d", searchedElement, algorithms.BinarySearch(mySlice, searchedElement))

	graph := [][]int{
		{1, 2},    // Node 0 is connected to nodes 1 and 2
		{0, 3, 4}, // Node 1 is connected to nodes 0, 3, and 4
		{0, 5},    // Node 2 is connected to nodes 0 and 5
		{1},       // Node 3 is connected to node 1
		{1},       // Node 4 is connected to node 1
		{2},       // Node 5 is connected to node 2
	}

	startNode := 0
	algorithms.BFS(graph, startNode)

}
