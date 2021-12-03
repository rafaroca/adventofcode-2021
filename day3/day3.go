package day3

import "fmt"

func Part1(input []string) int {
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

func Part2(input []string) int {
	return 0
}
