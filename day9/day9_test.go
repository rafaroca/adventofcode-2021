package day9

import (
	"strings"
	"testing"
)

var day = Day{}

func TestPart1(t *testing.T) {
	result := day.Part1(strings.Split(testInput, "\n"))
	if result != 15 {
		t.Error("sum of heights should be 15 but is", result)
	}
}

func TestPart2(t *testing.T) {
	result := day.Part2(strings.Split(testInput, "\n"))
	if result != 1134 {
		t.Error("three larges basin sizes multiplied should be 1134 but were", result)
	}
}

const testInput = `2199943210
3987894921
9856789892
8767896789
9899965678

`
