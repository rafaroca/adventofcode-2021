package day4

import (
	"strings"
	"testing"
)

var day = Day{}

func TestDayPart1(t *testing.T) {
	if result := day.Part1(testInput); result != 4512 {
		t.Errorf("Result was not 4512 but %d.", result)
	}
}

func TestDayPart2(t *testing.T) {
	if result := day.Part2(testInput); result != 1924 {
		t.Errorf("Result was not 1924 but %d.", result)
	}
}

func TestParseField(t *testing.T) {
	field := ParseField(testInput[2:8])
	expectedFirstRow := []int{22, 13, 17, 11, 0}
	if len(field.numbers) != 5 {
		t.Error("Field should contain 5 rows but contained", len(field.numbers), field.numbers)
	}
	if len(expectedFirstRow) != len(field.numbers[0]) {
		t.Error("First row should contain 5 columns but contained", len(field.numbers[0]), field.numbers)
	}
	for i, v := range expectedFirstRow {
		if field.numbers[0][i] != v {
			t.Error("First field should be", expectedFirstRow, "but is", field.numbers[0], field.numbers)
		}
	}
}

func TestCalculateScore(t *testing.T) {
	field := Field{
		numbers: [][]int{
			{14, 21, 17, 24, 4},
			{10, 16, 15, 9, 19},
			{18, 8, 23, 26, 20},
			{22, 11, 13, 6, 5},
			{2, 0, 12, 3, 7}},
		marked: map[int]struct{}{7: {}, 4: {}, 9: {}, 5: {}, 11: {}, 17: {}, 23: {}, 2: {}, 0: {}, 14: {}, 21: {}, 24: {}},
	}
	if score := CalculateScore(field, 24); score != 4512 {
		t.Error("Score should be 4512 but was", score)
	}
}

const rawTestInput = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7

`

var testInput = strings.Split(rawTestInput, "\n")
