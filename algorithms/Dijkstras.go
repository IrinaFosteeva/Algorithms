package algorithms

import (
	"fmt"
	"math"
)

// Graph представляет граф, на котором будет выполнен алгоритм Дейкстры.
type Graph struct {
	nodes []Node
	edges map[Node]map[Node]float64
}

// Node представляет узел графа.
type Node struct {
	name string
}

// AddNode добавляет новый узел в граф.
func (g *Graph) AddNode(n Node) {
	g.nodes = append(g.nodes, n)
}

// AddEdge добавляет новое ребро между двумя узлами графа.
func (g *Graph) AddEdge(n1, n2 Node, weight float64) {
	if g.edges == nil {
		g.edges = make(map[Node]map[Node]float64)
	}
	if g.edges[n1] == nil {
		g.edges[n1] = make(map[Node]float64)
	}
	g.edges[n1][n2] = weight
}

// Dijkstra реализует алгоритм Дейкстры для нахождения кратчайшего пути от начального узла до всех остальных узлов графа.
func Dijkstra(g Graph, start Node) map[Node]float64 {
	distances := make(map[Node]float64)
	visited := make(map[Node]bool)

	for _, node := range g.nodes {
		distances[node] = math.Inf(1)
	}
	distances[start] = 0

	for {
		currNode := minDistance(distances, visited)
		if currNode == (Node{}) {
			break
		}
		visited[currNode] = true

		for neighbor, weight := range g.edges[currNode] {
			if distances[currNode]+weight < distances[neighbor] {
				distances[neighbor] = distances[currNode] + weight
			}
		}
	}

	return distances
}

// minDistance находит узел с наименьшим значением расстояния из непосещенных узлов.
func minDistance(distances map[Node]float64, visited map[Node]bool) Node {
	minDist := math.Inf(1)
	var minNode Node

	for node, dist := range distances {
		if !visited[node] && dist < minDist {
			minDist = dist
			minNode = node
		}
	}

	return minNode
}

func main() {
	// Создаем граф
	graph := Graph{}

	// Добавляем узлы
	nodeA := Node{"A"}
	nodeB := Node{"B"}
	nodeC := Node{"C"}
	nodeD := Node{"D"}
	graph.AddNode(nodeA)
	graph.AddNode(nodeB)
	graph.AddNode(nodeC)
	graph.AddNode(nodeD)

	// Добавляем ребра и их веса
	graph.AddEdge(nodeA, nodeB, 10)
	graph.AddEdge(nodeA, nodeC, 15)
	graph.AddEdge(nodeB, nodeD, 12)
	graph.AddEdge(nodeC, nodeD, 2)

	// Вызываем алгоритм Дейкстры для поиска кратчайшего пути от узла A до всех остальных узлов
	distances := Dijkstra(graph, nodeA)

	// Выводим результаты
	for node, dist := range distances {
		fmt.Printf("Shortest distance from A to %s: %v\n", node.name, dist)
	}
}
