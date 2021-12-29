package day6

import (
	"advent/parser"
	"fmt"
	"strings"
	"testing"
)

var day6 = Day{}

func TestFish(t *testing.T) {
	firstIteration := parser.ParseCommaSeparatedLine(testInput[0])
	iteration := Uint8FromInt(firstIteration)
	fmt.Println("First Iteration", iteration)
	iter := Iter(&iteration)
	fmt.Println("Second Iteration", iter)
	iter = Iter(iter)
	fmt.Println("Third Iteration", iter)
}

func TestPart1(t *testing.T) {
	result := day6.Part1(testInput)
	if result != 5934 {
		t.Error("Sum ain't right, should be 5934 but is", result)
	}
}

func TestPart2(t *testing.T) {
	result := day6.Part2(testInput)
	// if result != 26 {
	if result != 26984457539 {
		t.Error("Sum ain't right, should be 26984457539 but is", result)
		// t.Error("Sum ain't right, should be 26 but is", result)
	}
}

const rawTestInput = `3,4,3,1,2
`

var testInput = strings.Split(rawTestInput, "\n")
