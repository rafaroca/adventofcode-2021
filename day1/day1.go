package day1

import "strconv"

func convertToInt(input []string) []int {
	result := make([]int, len(input))
	for i, v := range input {
		result[i], _ = strconv.Atoi(v)
	}
	return result
}

func Part1(input []string) int {
	return Part1int(convertToInt(input))
}

func Part1int(input []int) int {
	incCount := 0
	lastElem := 100000
	for _, elem := range input {
		if elem > lastElem {
			incCount++
		}
		lastElem = elem
	}
	return incCount
}

func Part2(input []string) int {
	elems := convertToInt(input)
	slidingWindow := SlidingWindow(elems)
	return Part1int(slidingWindow)
}

func SlidingWindow(input []int) []int {
	if len(input) < 3 {
		return input
	}

	result := make([]int, len(input)-2)

	for i := range input {
		if i >= len(input)-2 {
			continue
		}
		elem0 := input[i]
		elem1 := input[i+1]
		elem2 := input[i+2]

		result[i] = elem0 + elem1 + elem2
	}

	return result
}
