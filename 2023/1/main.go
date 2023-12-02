package main

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	acc := 0
	for scanner.Scan() {
		line := scanner.Text()

		first := ""
		last := ""

		for i := 0; i < len(line); i++ {
			letter := (line[i])
			if unicode.IsNumber(rune(letter)) {
				digit := string(letter)
				if first == "" {
					first = digit
				}
				last = digit
			}

		}
		result, _ := strconv.Atoi(first + last)
		acc += result
	}

	print(acc)
	defer file.Close()
}
