package main

import (
	"advent/day1"
	"advent/day15"
	"advent/day2"
	"advent/day3"
	"advent/day4"
	"advent/day5"
	"advent/day6"
	"advent/day7"
	"advent/day8"
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
	day1.Day{},
	day2.Day{},
	day3.Day{},
	day4.Day{},
	day5.Day{},
	day6.Day{},
	day7.Day{},
	day8.Day{},
	day15.Day{},
}

func main() {
	for _, v := range days {
		input := parser.ReadInputFile(v.InputFilename())
		start := time.Now()
		fmt.Printf("%3s: %d %d\n", v.InputFilename(), v.Part1(input), v.Part2(input))
		fmt.Println("This day took", time.Since(start))
	}
}
