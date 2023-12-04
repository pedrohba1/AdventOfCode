package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	Number      int
	FirstArray  []int
	SecondArray []int
	amount      int
}

type CardHeap []Card

func (h CardHeap) Len() int { return len(h) }
func (h *CardHeap) Push(x Card) {

	*h = append(*h, x)
}

func (h *CardHeap) Pop() Card {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func parseCard(line string) Card {
	parts := strings.SplitN(line, ":", 2)
	cardNumber, _ := strconv.Atoi(strings.TrimSpace(parts[0][4:]))
	numberParts := strings.Split(parts[1], "|")
	firstArray := parseNumbers(numberParts[0])
	secondArray := parseNumbers(numberParts[1])
	return Card{
		Number:      cardNumber,
		FirstArray:  firstArray,
		SecondArray: secondArray,
		amount:      1,
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

func countDupes(c Card) int {
	count := 0
	for _, num1 := range c.FirstArray {
		for _, num2 := range c.SecondArray {
			if num1 == num2 {
				count++
				break
			}
		}
	}
	return count
}

func main() {
	fmt.Println("--- part 2 faster algorithm ---")
	file, err := os.Open("../input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	scanner := bufio.NewScanner(file)

	var cards []Card

	for scanner.Scan() {
		line := scanner.Text()
		card := parseCard(line)
		cards = append(cards, card)
	}

	for i, card := range cards {
		wins := countDupes(card)

		if wins > 0 {
			for j := i + 1; j <= i+wins; j++ {
				card := cards[j]
				card.amount += cards[i].amount
				cards[j] = card
			}
		}
	}

	acc := 0
	for _, card := range cards {
		acc += card.amount
	}

	fmt.Println(acc)

}
