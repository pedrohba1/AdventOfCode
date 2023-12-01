package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	scanner := bufio.NewScanner(file)
	acc := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		firstDigit := ""
		lastDigit := ""

		for i := 0; i < len(line); i++ {
			letter := (line[i])
			if unicode.IsNumber(rune(letter)) {
				digit := string(letter)

				if firstDigit == "" {
					firstDigit = digit
				}

				lastDigit = digit
			}

		}
		result, _ := strconv.Atoi(firstDigit + lastDigit)
		acc += result
	}

	print(acc)
	defer file.Close()
}
