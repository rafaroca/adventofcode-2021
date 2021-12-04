package main

import (
	"advent/day1"
	"advent/day3"
	"bufio"
	"fmt"
	"os"
)

type Day interface {
	InputFilename() string
	Part1([]string) int
	Part2([]string) int
}

var days = []Day{day1.Day1{}, day3.Day3{}}

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
	for i, v := range days {
		dayCount := i + 1
		input := ReadInputFile(v.InputFilename())
		fmt.Printf("Day %d: %d %d\n", dayCount, v.Part1(input), v.Part2(input))
	}
}
