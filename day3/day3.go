package day3

import (
	"fmt"
	"strconv"
)

type Day3 struct {
}

func (t Day3) InputFilename() string {
	return "day3/input"
}

func (t Day3) Part1(input []string) int {
	bitCount := len(input[0])
	fmt.Printf("Running on %d bits.\n", bitCount)

	resultBits := make([]int, bitCount)
	for _, v := range input {
		for i := 0; i < bitCount; i++ {
			if v[i] == '1' {
				resultBits[i]++
			} else {
				resultBits[i]--
			}
		}
	}
	gamma := 0
	for i, v := range resultBits {
		if v > 0 {
			gamma |= 1 << (bitCount - i - 1)
		}
	}
	epsilon := 0
	for i, v := range resultBits {
		if v < 0 {
			epsilon |= 1 << (bitCount - i - 1)
		}
	}
	return gamma * epsilon
}

func (t Day3) Part2(input []string) int {
	bitCount := len(input[0])
	fmt.Printf("Running on %d bits.\n", bitCount)

	remainingCandidates := input

	for i := 0; i < bitCount; i++ {
		bitSum := sumOfBitsAtPosition(remainingCandidates, i)
		var bit string
		if bitSum >= 0 {
			bit = "1"
		} else {
			bit = "0"
		}
		remainingCandidates = FilterCandidates(remainingCandidates, i, bit)
		if len(remainingCandidates) == 1 {
			fmt.Println("Went through", i, "bits to filter until", remainingCandidates)
			break
		}
	}

	oxygen, _ := strconv.ParseInt(remainingCandidates[0], 2, 32)

	remainingCandidates = input

	for i := 0; i < bitCount; i++ {
		bitSum := sumOfBitsAtPosition(remainingCandidates, i)
		var bit string
		if bitSum >= 0 {
			bit = "0"
		} else {
			bit = "1"
		}
		remainingCandidates = FilterCandidates(remainingCandidates, i, bit)
		if len(remainingCandidates) == 1 {
			fmt.Println("Went through", i, "bits to filter until", remainingCandidates)
			break
		}
	}

	co2, _ := strconv.ParseInt(remainingCandidates[0], 2, 32)

	return int(oxygen * co2)
}

func FilterCandidates(candidates []string, index int, bit string) []string {
	var result []string
	for _, v := range candidates {
		if string(v[index]) == bit {
			result = append(result, v)
		}
	}
	return result
}

func sumOfBitsAtPosition(candidates []string, i int) int {
	var result int
	for _, v := range candidates {
		if v[i] == '1' {
			result++
		} else {
			result--
		}
	}
	return result
}
