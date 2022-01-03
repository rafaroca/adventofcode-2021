package day8

import (
	"fmt"
	"strings"
)

type Day struct{}

func (t Day) Part1(input []string) int {
	onefourseveneights := 0
	for _, rawLine := range input {
		_, outputs := parseSignalsAndOutput(rawLine)
		for _, output := range outputs {
			switch len(output) {
			case 2, 3, 4, 7:
				onefourseveneights++
			}
		}
	}
	return onefourseveneights
}

func parseSignalsAndOutput(rawLine string) (signals, outputs []string) {
	if len(rawLine) == 0 {
		return
	}
	linePart := strings.Split(rawLine, " | ")
	signals = strings.Split(linePart[0], " ")
	outputs = strings.Split(linePart[1], " ")
	return
}

var zeroSegs = "abcefg"
var oneSegs = "cf"
var twoSegs = "acdeg"
var threeSegs = "acdfg"
var fourSegs = "bcdf"
var fiveegs = "abdfg"
var sixSegs = "abdefg"
var sevenSegs = "acf"
var eightSegs = "abcdefg"
var nineSegs = "abcdfg"

func (t Day) Part2(input []string) int {
	//outputSum := 0

	// map of letter to possible candidate letter a -> abcdefg, b -> abcdefg, ...
	candidates := InitializeCandidates()

	for _, rawLine := range input {
		signals, _ := parseSignalsAndOutput(rawLine)
		for _, signal := range signals {
			AdjustCandidates(&candidates, signal)
		}
	}
	// remove possible candidates on each signal. "eb" then e -> cf, b -> cf. All others may not contain cf anymore
	// repeat until all signals correspond to a single letter only
	// decode output, sum up
	return 0
}

func AdjustCandidates(candidates *map[string]string, signal string) {

}

func Intersection(a, b string) string {
	var intersection string
	for _, char := range a {
		if strings.Contains(b, string(char)) {
			intersection = fmt.Sprintf("%s%c", intersection, char)
		}
	}
	return intersection
}

func InitializeCandidates() map[string]string {
	candidates := make(map[string]string)
	for _, segment := range strings.Split(eightSegs, "") {
		candidates[segment] = eightSegs
	}
	return candidates
}

func (t Day) InputFilename() string {
	return "day8/input"
}
