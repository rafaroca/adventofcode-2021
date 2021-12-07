package parser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func ParseCommaSeparatedLine(input string) []int {
	rawNumbers := strings.Split(input, ",")
	numbers := make([]int, len(rawNumbers))
	for i, rawNumber := range rawNumbers {
		numbers[i], _ = strconv.Atoi(rawNumber)
	}
	return numbers
}
