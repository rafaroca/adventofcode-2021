package day1

import (
	"strings"
	"testing"
)

func TestDayPart1(t *testing.T) {
	result := Day{}.Part1(strings.Split(testInput, "\n"))

	if result != 7 {
		t.Errorf("result is not 7 but: %d", result)
	}
}

func TestDayPart2(t *testing.T) {
	result := Day{}.Part2(strings.Split(testInput, "\n"))

	if result != 5 {
		t.Errorf("There should be 5 sums larger than the previous but there were: %d", result)
	}
}

func TestDaySlidingWindow(t *testing.T) {
	test_input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected_output := []int{607, 618, 618, 617, 647, 716, 769, 792}
	result := SlidingWindow(test_input)

	for i, v := range result {
		if v != expected_output[i] {
			t.Errorf("%d != %d at index %d", v, expected_output[i], i)
		}
	}
}

const testInput = `199
200
208
210
200
207
240
269
260
263

`
