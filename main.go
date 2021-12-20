package main

import (
	"advent/day1"
	"advent/day15"
	"advent/day3"
	"advent/day4"
	"advent/day5"
	"advent/day6"
	"advent/day7"
	"advent/parser"
	"fmt"
	"time"
)

type Day interface {
	InputFilename() string
	Part1([]string) int
	Part2([]string) int
}

var days = []Day{
	day1.Day1{},
	day3.Day3{},
	day4.Day4{},
	day5.Day5{},
	day6.Day6{},
	day7.Day7{},
	day15.Day15{},
}

func main() {
	for _, v := range days {
		input := parser.ReadInputFile(v.InputFilename())
		start := time.Now()
		fmt.Printf("%3s: %d %d\n", v.InputFilename(), v.Part1(input), v.Part2(input))
		fmt.Println("This day took", time.Since(start))
	}
}
