package main

import (
	"advent/day1"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadInputFile() []int {
	file, err := os.Open("day1/input1")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var result []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		elem, _ := strconv.Atoi(scanner.Text())
		result = append(result, elem)
	}
	return result
}

func main() {
	input := ReadInputFile()
	fmt.Println("Day 1:", day1.Part1(input), day1.Part2(input))
}
