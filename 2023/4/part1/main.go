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
	Number      int
	FirstArray  []int
	SecondArray []int
}

func parseCardLine(line string) Card {
	parts := strings.SplitN(line, ":", 2)
	cardNumber, _ := strconv.Atoi(strings.TrimSpace(parts[0][4:]))

	numberParts := strings.Split(parts[1], "|")
	firstArray := parseNumbers(numberParts[0])

	secondArray := parseNumbers(numberParts[1])

	return Card{
		Number:      cardNumber,
		FirstArray:  firstArray,
		SecondArray: secondArray,
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

func countCommon(firstArray, secondArray []int) int {
	count := 0
	for _, num1 := range firstArray {
		for _, num2 := range secondArray {
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
		card := parseCardLine(line)
		winCount := countCommon(card.SecondArray, card.FirstArray)
		if winCount == 0 {
			continue
		}
		acc += int(math.Pow(float64(2), float64(winCount-1)))
	}

	fmt.Print(acc)
}
