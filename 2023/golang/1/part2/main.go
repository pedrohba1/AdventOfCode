package main

import (
	"bufio"
	"os"
	"strings"
)

func findNumInLine(line string) int {

	termMap := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	firstMatchIndex := len(line)
	lastIndex := -1
	var first string
	var last string

	for key := range termMap {
		currentFirstIdx := strings.Index(line, key)
		currentLastIdx := strings.LastIndex(line, key)
		if currentFirstIdx != -1 {
			if currentFirstIdx < firstMatchIndex {
				firstMatchIndex = currentFirstIdx
				first = key

			}
			if currentLastIdx > lastIndex {
				lastIndex = currentLastIdx
				last = key
			}
		}
	}

	return termMap[first]*10 + termMap[last]
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	acc := 0
	for scanner.Scan() {
		line := scanner.Text()

		result := findNumInLine(line)

		acc += result
	}

	print(acc)
}
