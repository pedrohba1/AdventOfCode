package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()

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

	fmt.Println(acc)
}
