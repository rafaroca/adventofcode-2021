package day15

import (
	"math"
	"strconv"
	"strings"
)

type Day struct{}

type Node struct {
	x, y int
}

type OpenListNode struct {
	node Node
	f    int
}

func (t Day) Part1(input []string) int {
	lineLength := len(input[0])
	lengths := parseLengths(input)
	return ShortestPathFromTopLeftToBottomRight(lengths, lineLength)
}

func ShortestPathFromTopLeftToBottomRight(lengths map[Node]int, lineLength int) int {
	shortestPaths := InitializeShortestPaths(lineLength)
	queue := []Node{{0, 0}}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		neighbors := Neighbors(node, lineLength)
		for _, neighborNode := range neighbors {
			currentLength := shortestPaths[node]
			lengthToNeighbor := currentLength + lengths[neighborNode]
			if lengthToNeighbor < shortestPaths[neighborNode] {
				shortestPaths[neighborNode] = lengthToNeighbor
				queue = append(queue, neighborNode)
			}
		}
	}
	return shortestPaths[Node{lineLength - 1, lineLength - 1}]
}

func InitializeShortestPaths(lineLength int) map[Node]int {
	paths := make(map[Node]int)
	for i := 0; i < lineLength; i++ {
		for j := 0; j < lineLength; j++ {
			paths[Node{i, j}] = math.MaxInt
		}
	}
	paths[Node{0, 0}] = 0
	return paths
}

func Neighbors(node Node, lineLength int) []Node {
	var neighbors []Node

	if node.x != lineLength-1 {
		neighbors = append(neighbors, Node{node.x + 1, node.y})
	}
	if node.x != 0 {
		neighbors = append(neighbors, Node{node.x - 1, node.y})
	}
	if node.y != lineLength-1 {
		neighbors = append(neighbors, Node{node.x, node.y + 1})
	}
	if node.y != 0 {
		neighbors = append(neighbors, Node{node.x, node.y - 1})
	}

	return neighbors
}

func parseLengths(input []string) map[Node]int {
	weights := make(map[Node]int)

	for y, line := range input {
		if len(line) == 0 {
			continue
		}
		lineWeightsRaw := strings.Split(line, "")

		for x, char := range lineWeightsRaw {
			weights[Node{x, y}], _ = strconv.Atoi(char)
		}
	}
	return weights
}

func (t Day) Part2(input []string) int {
	lengths := parseLengths(input)
	originalLineLength := len(input[0])
	enlargedField := EnlargeField(lengths, originalLineLength)
	enlargedLineLength := originalLineLength * 5
	return ShortestPathFromTopLeftToBottomRight(enlargedField, enlargedLineLength)
}

func EnlargeField(input map[Node]int, lineLength int) (enlarged map[Node]int) {
	enlarged = make(map[Node]int)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for y := 0; y < lineLength; y++ {
				for x := 0; x < lineLength; x++ {
					incOriginal := input[Node{x, y}] + i + j
					enlarged[Node{x + (lineLength * i), y + (lineLength * j)}] = incOriginal%10 + incOriginal/10
				}
			}
		}
	}
	return
}

func (t Day) InputFilename() string {
	return "day15/input"
}
