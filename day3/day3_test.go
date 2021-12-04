package day3

import "testing"

var test_input = []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}

func TestDay3Part1(t *testing.T) {
	day := Day3{}
	if result := day.Part1(test_input); result != 198 {
		t.Errorf("The result should be 198 but was '%d'", result)
	}
}

func TestFilterCandidates(t *testing.T) {
	result := FilterCandidates([]string{"001", "101", "111"}, 0, "1")
	if len(result) != 2 || result[0] != "101" || result[1] != "111" {
		t.Error("Should only consist of 101 and 111, but is ", result)
	}
}

func TestDay3Part2(t *testing.T) {
	day := Day3{}
	if result := day.Part2(test_input); result != 230 {
		t.Errorf("The result should be 230 but was '%d'", result)
	}
}
