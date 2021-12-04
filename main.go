package main

import (
	"advent/day1"
	"advent/day3"
	"advent/day4"
	"bufio"
	"fmt"
	"os"
)

type Day interface {
	InputFilename() string
	Part1([]string) int
	Part2([]string) int
}

var days = []Day{day1.Day1{}, day3.Day3{}, day4.Day4{}}

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
	for _, v := range days {
		input := ReadInputFile(v.InputFilename())
		fmt.Printf("%3s: %d %d\n", v.InputFilename(), v.Part1(input), v.Part2(input))
	}
}
