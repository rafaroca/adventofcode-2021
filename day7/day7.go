package day7

import (
	"advent/parser"
	"math"
)

type Day7 struct{}

func (t Day7) Part1(input []string) int {
	positions := parser.ParseCommaSeparatedLine(input[0])
	max := calculateMax(positions)
	fuels := make([]int, max+1)
	for i := 0; i <= max; i++ {
		for _, pos := range positions {
			fuels[i] += abs(i - pos)
		}
	}
	return calculateMin(fuels)
}

func (t Day7) Part2(input []string) int {
	positions := parser.ParseCommaSeparatedLine(input[0])
	max := calculateMax(positions)
	fuels := make([]int, max+1)
	for i := 0; i <= max; i++ {
		for _, pos := range positions {
			fuels[i] += calculateIncreasingFuel(i, pos)
		}
	}
	return calculateMin(fuels)
}

func calculateIncreasingFuel(start, end int) int {
	diff := abs(end - start)
	return diff * (diff + 1) / 2
}

func (t Day7) InputFilename() string {
	return "day7/input"
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func calculateMax(positions []int) int {
	max := 0
	for _, v := range positions {
		if v > max {
			max = v
		}
	}
	return max
}

func calculateMin(fuels []int) int {
	min := math.MaxInt64
	for _, v := range fuels {
		if v < min {
			min = v
		}
	}
	return min
}
