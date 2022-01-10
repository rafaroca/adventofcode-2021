package day9

import "strconv"

type Day struct{}

func (t Day) Part1(input []string) int {
	if len(input[len(input)-1]) == 0 {
		input = input[:len(input)-1]
	}
	lineLength := len(input[0])
	lineCount := len(input)

	heightCount := 0
	for y, rawLine := range input {
		for x, rawHeight := range rawLine {
			height, _ := strconv.Atoi(string(rawHeight))
			neighbors := Neighbors(x, y, &input, lineLength, lineCount)
			if IsLowPoint(height, neighbors) {
				heightCount += height + 1
			}
		}
	}
	return heightCount
}

func IsLowPoint(height int, neighbors []int) bool {
	isLow := true
	for _, neighborHeight := range neighbors {
		if height >= neighborHeight {
			isLow = false
		}
	}
	return isLow
}

func Neighbors(x int, y int, field *[]string, lineLength int, lineCount int) (neighbors []int) {
	if x != lineLength-1 {
		rightNeighbor, _ := strconv.Atoi(string((*field)[y][x+1]))
		neighbors = append(neighbors, rightNeighbor)
	}
	if x != 0 {
		leftNeighbor, _ := strconv.Atoi(string((*field)[y][x-1]))
		neighbors = append(neighbors, leftNeighbor)
	}
	if y != lineCount-1 {
		lowerNeighbor, _ := strconv.Atoi(string((*field)[y+1][x]))
		neighbors = append(neighbors, lowerNeighbor)
	}
	if y != 0 {
		upperNeighbor, _ := strconv.Atoi(string((*field)[y-1][x]))
		neighbors = append(neighbors, upperNeighbor)
	}
	return
}

func (t Day) Part2(input []string) int {
	outputSum := 0

	return outputSum
}

func (t Day) InputFilename() string {
	return "day9/input"
}
