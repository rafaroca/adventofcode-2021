package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input1")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	incCount := 0
	lastElem := 100000
	for scanner.Scan() {
		elem, _ := strconv.Atoi(scanner.Text())
		if elem > lastElem {
			incCount++
		}
		lastElem = elem
	}
	fmt.Println(incCount)
}
