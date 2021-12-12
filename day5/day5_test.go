package day5

import (
	"strings"
	"testing"
)

var day = Day5{}

func TestLinesPart1(t *testing.T) {
	result := day.Part1(strings.Split(testInput, "\n"))
	if result != 5 {
		t.Error("Resulting lines count should be 5 but was", result)
	}
}

func TestStraightLines(t *testing.T) {
	rawLines := ParseLines(strings.Split(testInput, "\n"))
	lines := FilterStraightLines(rawLines)
	if len(lines) != 6 {
		t.Error("Should be 6 straight lines but are", len(lines))
	}
}

const testInput = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2

`
