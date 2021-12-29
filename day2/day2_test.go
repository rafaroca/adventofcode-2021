package day2

import (
	"strings"
	"testing"
)

var day = Day{}

func TestPart1(t *testing.T) {
	result := day.Part1(strings.Split(testInput, "\n"))
	if result != 150 {
		t.Error("Depth should be 150 but was", result)
	}
}

func TestPart2(t *testing.T) {
	result := day.Part2(strings.Split(testInput, "\n"))
	if result != 900 {
		t.Error("Depth should be 900 but was", result)
	}
}

const testInput = `forward 5
down 5
forward 8
up 3
down 8
forward 2

`
