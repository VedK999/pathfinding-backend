package algorithms

import (
	"container/heap"
	"math"
)

type Node struct {
	Row      int     `json:"row"`
	Col      int     `json:"col"`
	IsWall   bool    `json:"isWall"`
	Distance float64 `json:"distance,omitempty"`
	Visited  bool    `json:"visited,omitempty"`
	Previous *Node   `json:"previous,omitempty"`
	FScore   float64 `json:"fscore,omitempty"`
	GScore   float64 `json:"gscore,omitempty"`
}

type AStarQueue []*Node

func (pq AStarQueue) Len() int {
	return len(pq)
}

func (pq AStarQueue) Less(i, j int) bool {
	return pq[i].FScore < pq[j].FScore
}

func (pq *AStarQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *AStarQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

func (pq AStarQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func hscore(a, b Node) float64 {
	return math.Abs(float64(a.Row-b.Row) + float64(a.Col-b.Col))
}

func AStar(grid [][]Node, start, end Node) ([]Node, []Node, bool) {

	for i := range grid {
		for j := range grid[0] {
			grid[i][j].GScore = math.Inf(1)
			grid[i][j].FScore = math.Inf(1)
			grid[i][j].Visited = false
			grid[i][j].Previous = nil
		}
	}

	grid[start.Row][start.Col].GScore = 0
	grid[start.Row][start.Col].FScore = hscore(start, end)

	pq := &AStarQueue{}
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
				tentativeGScore := current.GScore + 1

				if tentativeGScore < grid[neighbor.Row][neighbor.Col].GScore {
					grid[neighbor.Row][neighbor.Col].Previous = current
					grid[neighbor.Row][neighbor.Col].GScore = tentativeGScore
					grid[neighbor.Row][neighbor.Col].FScore = tentativeGScore + hscore(neighbor, end)
					heap.Push(pq, &grid[neighbor.Row][neighbor.Col])
				}
			}
		}
	}

	return []Node{}, []Node{}, false
}
