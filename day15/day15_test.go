package day15

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var day = Day15{}

func TestPart1(t *testing.T) {
	result := day.Part1(strings.Split(testInput, "\n"))
	if result != 40 {
		t.Error("Shortest path should be 40 but was", result)
	}
}

func TestPart2(t *testing.T) {
	result := day.Part2(strings.Split(testInput, "\n"))
	if result != 0 {
		t.Error("Should align on fuel 168 but was", result)
	}
}

func TestNeighborsCorner(t *testing.T) {
	result := Neighbors(Node{0, 0}, 10)
	assert.ElementsMatch(t, []Node{{1, 0}, {0, 1}}, result)
}

func TestNeighborsEdge(t *testing.T) {
	result := Neighbors(Node{1, 0}, 10)
	assert.ElementsMatch(t, []Node{{2, 0}, {0, 0}, {1, 1}}, result)
}

func TestNeighborsMiddle(t *testing.T) {
	result := Neighbors(Node{1, 1}, 10)
	assert.ElementsMatch(t, []Node{{1, 0}, {0, 1}, {2, 1}, {1, 2}}, result)
}

const testInput = `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581

`
