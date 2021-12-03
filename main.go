package main

import (
	"advent/day1"
	"bufio"
	"fmt"
	"os"
)

func ReadInputFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}

func main() {
	input1 := ReadInputFile("day1/input")
	fmt.Println("Day 1:", day1.Part1(input1), day1.Part2(input1))
}
