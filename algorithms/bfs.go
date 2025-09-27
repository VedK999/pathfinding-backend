package algorithms

import (
	"fmt"
)

func getNeighbours(grid [][]Node, current Node) []Node {
	neighbours := []Node{}
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, dir := range directions {
		newRow := current.Row + dir[0]
		newCol := current.Col + dir[1]

		if newRow >= 0 && newRow < len(grid) && newCol >= 0 && newCol < len(grid[0]) && !grid[newRow][newCol].IsWall {
			neighbours = append(neighbours, grid[newRow][newCol])
		}
	}

	return neighbours

}

func constructPath(node *Node) []Node {
	path := []Node{}
	curr := node

	for curr != nil {
		path = append([]Node{*curr}, path...)
		curr = curr.Previous
	}

	fmt.Printf("Reconstructed path length: %d\n", len(path))
	for i, node := range path {
		fmt.Printf("Path[%d]: (%d, %d)\n", i, node.Row, node.Col)
	}

	return path
}

func BFS(grid [][]Node, start, end Node) ([]Node, []Node, bool) {
	for i := range grid {
		for j := range grid[0] {
			grid[i][j].Visited = false
			grid[i][j].Previous = nil
		}
	}

	queue := []Node{start}
	grid[start.Row][start.Col].Visited = true
	visitedNodes := []Node{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		visitedNodes = append(visitedNodes, current)

		if current.Row == end.Row && current.Col == end.Col {
			path := constructPath(&grid[current.Row][current.Col])
			return path, visitedNodes, true
		}

		neighbours := getNeighbours(grid, current)
		for _, neighbour := range neighbours {
			if !grid[neighbour.Row][neighbour.Col].Visited {
				grid[neighbour.Row][neighbour.Col].Visited = true
				grid[neighbour.Row][neighbour.Col].Previous = &grid[current.Row][current.Col]
				queue = append(queue, grid[neighbour.Row][neighbour.Col])
			}
		}

	}

	return []Node{}, visitedNodes, false
}
