package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	Number         int
	WinningNumbers []int
	MyNumbers      []int
}

func parseCard(line string) Card {
	parts := strings.SplitN(line, ":", 2)
	cardNumber, _ := strconv.Atoi(strings.TrimSpace(parts[0][4:]))

	numberParts := strings.Split(parts[1], "|")
	firstArray := parseNumbers(numberParts[0])

	secondArray := parseNumbers(numberParts[1])

	return Card{
		Number:         cardNumber,
		WinningNumbers: firstArray,
		MyNumbers:      secondArray,
	}
}

func parseNumbers(s string) []int {
	strNumbers := strings.Fields(s)
	var numbers []int
	for _, strNum := range strNumbers {
		num, _ := strconv.Atoi(strNum)
		numbers = append(numbers, num)
	}
	return numbers
}

func countWins(c Card) int {
	count := 0
	for _, num1 := range c.WinningNumbers {
		for _, num2 := range c.MyNumbers {
			if num1 == num2 {
				count++
				break
			}
		}
	}
	return count
}

func main() {
	file, err := os.Open("../input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	scanner := bufio.NewScanner(file)

	acc := 0
	for scanner.Scan() {
		line := scanner.Text()
		card := parseCard(line)
		winCount := countWins(card)
		if winCount == 0 {
			continue
		}
		acc += int(math.Pow(float64(2), float64(winCount-1)))
	}

	fmt.Println(acc)
}
