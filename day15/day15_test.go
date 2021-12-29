package day15

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var day = Day{}

func TestPart1(t *testing.T) {
	result := day.Part1(strings.Split(testInput, "\n"))
	if result != 40 {
		t.Error("Shortest path should be 40 but was", result)
	}
}

func TestPart2(t *testing.T) {
	result := day.Part2(strings.Split(testInput, "\n"))
	if result != 315 {
		t.Error("Shortest path should be 315 but was", result)
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

func TestEnlargeField(t *testing.T) {
	m := EnlargeField(map[Node]int{{0, 0}: 7}, 1)
	if len(m) != 25 {
		t.Error("Enlarged field should have 25 entries but has", len(m))
	}
	if result := m[Node{4, 4}]; result != 6 {
		t.Error("Last field should be 6 but is", result)
	}
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
