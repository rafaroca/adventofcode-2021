package day4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day4 struct{}

func (t Day4) InputFilename() string {
	return "day4/input"
}

func (t Day4) Part1(input []string) int {
	rawDrawnNumbers := strings.Split(input[0], ",")
	drawnNumbers := make([]int, len(rawDrawnNumbers))
	for i, rawDrawnNumber := range rawDrawnNumbers {
		drawnNumbers[i], _ = strconv.Atoi(rawDrawnNumber)
	}

	// Don't count header lines, every field has 6 lines because of the trailing empty line
	fieldCount := (len(input) - 2) / 6
	rawFields := input[2:]
	fields := make([]Field, fieldCount)
	for i := 0; i < fieldCount; i++ {
		fields[i] = ParseField(rawFields[i*6 : (i*6)+6])
	}

	// play games
	for _, drawnNumber := range drawnNumbers {
		for i := range fields {
			fields[i].marked[drawnNumber] = true
		}
		// check for winning field
		// TODO: check only row and column of last drawn number
		winningField := -1
		for iField, field := range fields {
			for _, rowNumbers := range field.numbers {
				winningRow := true
				for _, number := range rowNumbers {
					if _, marked := field.marked[number]; marked == false {
						winningRow = false
						break
					}
				}
				if winningRow {
					winningField = iField
					break
				}
			}

			for i := 0; i < 5; i++ {
				winningColumn := true
				for j := 0; j < 5; j++ {
					if _, marked := field.marked[field.numbers[j][i]]; marked == false {
						winningColumn = false
						break
					}
				}
				if winningColumn {
					winningField = iField
					break
				}
			}
			if winningField >= 0 {
				fmt.Printf("Field %d won at winning number %d\n", winningField, drawnNumber)
				return CalculateScore(fields[winningField], drawnNumber)
			}
		}
	}
	return -1
}

func CalculateScore(field Field, drawnNumber int) int {
	unmarkedSum := 0
	for _, rowNumbers := range field.numbers {
		for _, number := range rowNumbers {
			if _, marked := field.marked[number]; marked == false {
				unmarkedSum += number
			}
		}
	}
	return unmarkedSum * drawnNumber
}

type Field struct {
	numbers [][]int
	marked  map[int]interface{}
}

func ParseField(rawField []string) Field {
	field := Field{}
	for row := 0; row < 5; row++ {
		rawColumnValues := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(rawField[row]), -1)
		var columnValues []int
		for _, rawValue := range rawColumnValues {
			intValue, _ := strconv.Atoi(rawValue)
			columnValues = append(columnValues, intValue)
		}
		field.numbers = append(field.numbers, columnValues)
	}
	field.marked = make(map[int]interface{})
	return field
}

func (t Day4) Part2(input []string) int {
	return 0
}
