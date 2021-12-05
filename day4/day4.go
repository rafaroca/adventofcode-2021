package day4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day4 struct{}

type Field struct {
	numbers [][]int
	marked  map[int]struct{}
}

type WonField struct {
	score              int
	winningDrawnNumber int
	id                 int
}

func (t Day4) InputFilename() string {
	return "day4/input"
}

func (t Day4) Part1(input []string) int {
	drawnNumbers := parseDrawnNumbers(input)
	fields := parseFields(input)

	// play games
	for _, drawnNumber := range drawnNumbers {
		for i := range fields {
			fields[i].marked[drawnNumber] = struct{}{}
		}
		// check for winning field
		// TODO: check only row and column of last drawn number
		wonFields := checkForWinningField(fields, drawnNumber)
		for _, wonField := range wonFields {
			return wonField.score
		}
	}
	return -1
}

func (t Day4) Part2(input []string) int {
	drawnNumbers := parseDrawnNumbers(input)
	fields := parseFields(input)

	// play games
	for _, drawnNumber := range drawnNumbers {
		for i := range fields {
			fields[i].marked[drawnNumber] = struct{}{}
		}
		// check for winning field
		wonFieldsOfRound := checkForWinningField(fields, drawnNumber)
		fields = removeWonFields(fields, wonFieldsOfRound)
		if len(fields) == 0 {
			fmt.Println("All fields won on number", drawnNumber)
			for _, wonField := range wonFieldsOfRound {
				if wonField.winningDrawnNumber == drawnNumber {
					fmt.Println("Field", wonField.id, "won.")
					return wonField.score
				}
			}
		}
	}
	return -1
}

func removeWonFields(fields []Field, wonFieldsOfRound map[int]WonField) []Field {
	var unwonFields []Field
	for i, field := range fields {
		if _, won := wonFieldsOfRound[i]; !won {
			unwonFields = append(unwonFields, field)
		}
	}
	return unwonFields
}

func parseFields(input []string) []Field {
	// Don't count header lines, every field has 6 lines because of the trailing empty line
	fieldCount := (len(input) - 2) / 6
	rawFields := input[2:]
	fields := make([]Field, fieldCount)
	for i := 0; i < fieldCount; i++ {
		fields[i] = ParseField(rawFields[i*6 : (i*6)+6])
	}
	return fields
}

func checkForWinningField(fields []Field, drawnNumber int) map[int]WonField {
	winningField := -1
	wonFields := make(map[int]WonField)

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
			wonFields[winningField] = WonField{
				score:              CalculateScore(fields[winningField], drawnNumber),
				winningDrawnNumber: drawnNumber,
				id:                 winningField,
			}
			winningField = -1
		}
	}
	return wonFields
}

func parseDrawnNumbers(input []string) []int {
	rawDrawnNumbers := strings.Split(input[0], ",")
	drawnNumbers := make([]int, len(rawDrawnNumbers))
	for i, rawDrawnNumber := range rawDrawnNumbers {
		drawnNumbers[i], _ = strconv.Atoi(rawDrawnNumber)
	}
	return drawnNumbers
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
	field.marked = make(map[int]struct{})
	return field
}
