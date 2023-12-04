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
	file, err := os.Open("../input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	scanner := bufio.NewScanner(file)

	acc := 0
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	copies := &CardHeap{}
	for _, line := range lines {
		card := parseCard(line)
		wins := countDupes(card)

		acc += 1

		for i := 1; i <= wins; i++ {
			cardEarned := parseCard(lines[card.Number-1+i])
			copies.Push(cardEarned)
			acc += 1
		}

		for copies.Len() > 0 {
			pCard := copies.Pop()
			wins := countDupes(pCard)
			for j := 1; j <= wins; j++ {
				new := parseCard(lines[pCard.Number-1+j])
				copies.Push(new)
				acc += 1
			}
		}
	}

	fmt.Println(acc)

}
