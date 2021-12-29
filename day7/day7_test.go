package day7

import (
	"strings"
	"testing"
)

var day = Day{}

func TestPart1(t *testing.T) {
	result := day.Part1(strings.Split(testInput, "\n"))
	if result != 37 {
		t.Error("Should align on fuel 37 but was", result)
	}
}

func TestPart2(t *testing.T) {
	result := day.Part2(strings.Split(testInput, "\n"))
	if result != 168 {
		t.Error("Should align on fuel 168 but was", result)
	}
}

const testInput = `16,1,2,0,4,2,7,1,2,14`
