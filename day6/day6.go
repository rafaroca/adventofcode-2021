package day6

import (
	"strconv"
	"strings"
)

const ITERATIONS = 80

func Part1(input []string) int {
	return 0
}

func ReadInputFish(input []string) []uint8 {
	rawFish := strings.Split(input[0], ",")
	fish := make([]uint8, len(rawFish))
	for i, rawSingleFish := range rawFish {
		convertedFish, _ := strconv.Atoi(rawSingleFish)
		fish[i] = uint8(convertedFish)
	}
	return fish
}

func Iter(input []uint8) []uint8 {
	result := input
	newborns := 0
	for _, v := range input {
		v--
		if v == 0 {
			v = 6
			newborns++
		}
		result = append(result, v-1)
	}
	for i := 0; i < newborns; i++ {
		result = append(result, 8)
	}
	return result
}
