package day5

import (
	"strconv"
	"strings"
)

type Day5 struct{}

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

var FIELD_SIZE = 1000

func (t Day5) Part1(input []string) int {
	lines := ParseLines(input)
	straightLines := FilterStraightLines(lines)
	field := make([][]uint8, FIELD_SIZE)
	for i := range field {
		field[i] = make([]uint8, FIELD_SIZE)
	}
	for _, line := range straightLines {
		if line.x1 == line.x2 {
			yMin := min(line.y1, line.y2)
			yMax := max(line.y1, line.y2)
			for i := yMin; i <= yMax; i++ {
				field[line.x1][i] += 1
			}
		}
		if line.y1 == line.y2 {
			xMin := min(line.x1, line.x2)
			xMax := max(line.x1, line.x2)
			for i := xMin; i <= xMax; i++ {
				field[i][line.y1] += 1
			}
		}
	}
	return CountIntersections(field)
}

func CountIntersections(field [][]uint8) int {
	intersections := 0
	for x := 0; x < FIELD_SIZE; x++ {
		for y := 0; y < FIELD_SIZE; y++ {
			if field[x][y] > 1 {
				intersections++
			}
		}
	}
	return intersections
}

func (t Day5) Part2(input []string) int {
	lines := ParseLines(input)
	field := make([][]uint8, FIELD_SIZE)
	for i := range field {
		field[i] = make([]uint8, FIELD_SIZE)
	}
	for _, line := range lines {
		if line.x1 == line.x2 {
			yMin := min(line.y1, line.y2)
			yMax := max(line.y1, line.y2)
			for i := yMin; i <= yMax; i++ {
				field[line.x1][i] += 1
			}
		} else if line.y1 == line.y2 {
			xMin := min(line.x1, line.x2)
			xMax := max(line.x1, line.x2)
			for i := xMin; i <= xMax; i++ {
				field[i][line.y1] += 1
			}
		} else {
			var xd int
			if line.x1 <= line.x2 {
				xd = 1
			} else {
				xd = -1
			}
			var yd int
			if line.y1 <= line.y2 {
				yd = 1
			} else {
				yd = -1
			}
			lineLength := line.x2 - line.x1
			if lineLength < 0 {
				lineLength = -lineLength
			}
			lineLength++

			for i := 0; i < lineLength; i++ {
				field[line.x1+(xd*i)][line.y1+(yd*i)] += 1
			}
		}

	}
	return CountIntersections(field)
}

func FilterStraightLines(lines []Line) []Line {
	var straightLines []Line
	for _, line := range lines {
		if line.x1 == line.x2 || line.y1 == line.y2 {
			straightLines = append(straightLines, line)
		}
	}
	return straightLines
}

func (t Day5) InputFilename() string {
	return "day5/input"
}

func ParseLines(rawLines []string) []Line {
	var lines []Line
	for _, rawLine := range rawLines {
		rawCoordinates := strings.Split(rawLine, " -> ")
		if len(rawCoordinates) == 1 {
			continue
		}
		rawCoordinates0 := strings.Split(rawCoordinates[0], ",")
		rawCoordinates1 := strings.Split(rawCoordinates[1], ",")
		x1, _ := strconv.Atoi(rawCoordinates0[0])
		y1, _ := strconv.Atoi(rawCoordinates0[1])
		x2, _ := strconv.Atoi(rawCoordinates1[0])
		y2, _ := strconv.Atoi(rawCoordinates1[1])
		lines = append(lines, Line{x1, y1, x2, y2})
	}
	return lines
}

func min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
