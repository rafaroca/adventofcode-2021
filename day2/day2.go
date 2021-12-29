package day2

import (
	"strconv"
	"strings"
)

type Day struct{}

func (t Day) Part1(input []string) int {
	depth := 0
	horizontal := 0
	for _, instrRaw := range input {
		if len(instrRaw) == 0 {
			continue
		}
		instr := strings.Split(instrRaw, " ")
		command := instr[0]
		value, _ := strconv.Atoi(instr[1])
		switch command {
		case "forward":
			horizontal += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}
	return depth * horizontal
}

func (t Day) Part2(input []string) int {
	aim := 0
	horizontal := 0
	depth := 0
	for _, instrRaw := range input {
		if len(instrRaw) == 0 {
			continue
		}
		instr := strings.Split(instrRaw, " ")
		command := instr[0]
		value, _ := strconv.Atoi(instr[1])
		switch command {
		case "forward":
			horizontal += value
			depth += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}
	return depth * horizontal
}

func (t Day) InputFilename() string {
	return "day2/input"
}
