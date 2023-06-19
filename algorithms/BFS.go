package algorithms

import "fmt"

type Queue []int

func (q *Queue) Enqueue(v int) {
	*q = append(*q, v)
}

func (q *Queue) Dequeue() int {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func BFS(graph [][]int, start int) {
	visited := make([]bool, len(graph))
	queue := Queue{}
	queue.Enqueue(start)
	visited[start] = true

	fmt.Println("BFS Traversal:")

	for len(queue) > 0 {
		vertex := queue.Dequeue()
		fmt.Printf("%d ", vertex)

		for _, v := range graph[vertex] {
			if !visited[v] {
				queue.Enqueue(v)
				visited[v] = true
			}
		}
	}

	fmt.Println()
}
