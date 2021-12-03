package day3

import "testing"

var test_input = []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}

func TestDay3(t *testing.T) {
	if result := Part1(test_input); result != 198 {
		t.Errorf("The result should be 198 but was '%d'", result)
	}
}
