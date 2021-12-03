package day1

import "testing"

func TestDay1Part1(t *testing.T) {
	result := Part1([]string{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"})

	if result != 7 {
		t.Errorf("result is not quite right: %d", result)
	}
}

func TestDay1SlidingWindow(t *testing.T) {
	test_input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected_output := []int{607, 618, 618, 617, 647, 716, 769, 792}
	result := SlidingWindow(test_input)

	for i, v := range result {
		if v != expected_output[i] {
			t.Errorf("%d != %d at index %d", v, expected_output[i], i)
		}
	}
}
