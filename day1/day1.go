package day1

func Part1(input []int) int {
	incCount := 0
	lastElem := 100000
	for _, v := range input {
		if v > lastElem {
			incCount++
		}
		lastElem = v
	}
	return incCount
}

func Part2(input []int) int {
	slidingWindow := SlidingWindow(input)

	return Part1(slidingWindow)
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
