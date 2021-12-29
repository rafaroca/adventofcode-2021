package day6

import (
	"advent/parser"
)

type Day struct{}

func (t Day) Part1(input []string) int {
	return FishSum(input, 80)
}

func (t Day) Part2(input []string) int {
	return FishSum(input, 256)
}

func FishSum(input []string, iterationCount int) int {
	fishState := ParseSeed(input[0])

	for i := 1; i <= iterationCount; i++ {
		ShiftFishState(&fishState)
	}

	sum := 0
	for _, v := range fishState {
		sum += v
	}
	return sum
}

func ShiftFishState(fishState *[]int) {
	newFish := (*fishState)[0]
	for i := 0; i < 8; i++ {
		(*fishState)[i] = (*fishState)[i+1]
	}
	(*fishState)[6] += newFish
	(*fishState)[8] = newFish
}

func ParseSeed(s string) []int {
	seeds := make([]int, 9)
	inputs := parser.ParseCommaSeparatedLine(s)
	for _, v := range inputs {
		seeds[v]++
	}
	return seeds
}

func (t Day) InputFilename() string {
	return "day6/input"
}
