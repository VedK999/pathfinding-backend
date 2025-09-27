package algorithms

import (
	"container/heap"
	"math"
)

type priorityQueue []*Node

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

func Dijkstra(grid [][]Node, start, end Node) ([]Node, []Node, bool) {
	for i := range grid {
		for j := range grid[i] {
			grid[i][j].Distance = math.Inf(1)
			grid[i][j].Visited = false
			grid[i][j].Previous = nil
		}
	}

	grid[start.Row][start.Col].Distance = 0

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &grid[start.Row][start.Col])

	visitedNodes := []Node{}

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Node)

		if current.Visited {
			continue
		}

		current.Visited = true
		visitedNodes = append(visitedNodes, *current)

		if current.Row == end.Row && current.Col == end.Col {
			path := constructPath(current)
			return path, visitedNodes, true
		}

		neighbors := getNeighbours(grid, *current)
		for _, neighbor := range neighbors {
			if !grid[neighbor.Row][neighbor.Col].Visited {
				newDistance := current.Distance + 1

				if newDistance < grid[neighbor.Row][neighbor.Col].Distance {
					grid[neighbor.Row][neighbor.Col].Distance = newDistance
					grid[neighbor.Row][neighbor.Col].Previous = current
					heap.Push(pq, &grid[neighbor.Row][neighbor.Col])
				}
			}
		}
	}

	return []Node{}, []Node{}, false
}
