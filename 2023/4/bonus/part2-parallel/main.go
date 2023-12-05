package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Card struct {
	Number      int
	FirstArray  []int
	SecondArray []int
}

type CardHeap struct {
	cards []Card
	end   bool
}

func (h *CardHeap) Len() int { return len(h.cards) }

func (h *CardHeap) Push(x Card) {
	h.cards = append(h.cards, x)
}

func (h *CardHeap) Pop() Card {
	if len(h.cards) == 0 {
		return Card{}
	}
	old := h.cards
	n := len(old)
	x := old[n-1]
	h.cards = old[0 : n-1]
	return x
}

func (h *CardHeap) Close() {
	h.end = true
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

var lines []string
var mutex sync.Mutex
var wg sync.WaitGroup
var acc int

func iterHeap(h *CardHeap) {
	defer wg.Done()
	for h.Len() > 0 {
		mutex.Lock()
		pCard := h.Pop()
		mutex.Unlock()
		if pCard.Number == 0 {
			continue
		}
		wins := countDupes(pCard)
		copies := &CardHeap{end: false}
		for j := 1; j <= wins; j++ {
			new := parseCard(lines[pCard.Number-1+j])
			mutex.Lock()
			copies.Push(new)
			acc += 1
			mutex.Unlock()

		}
		wg.Add(1)
		go iterHeap(copies)
	}
}

func main() {
	file, err := os.Open("../input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	copies := &CardHeap{end: false}
	for _, line := range lines {
		card := parseCard(line)
		wins := countDupes(card)
		mutex.Lock()
		acc += 1
		mutex.Unlock()

		wg.Add(1)
		go func() {
			defer wg.Done()
			wg.Add(1)
			go func(h *CardHeap) {
				defer wg.Done()
				for i := 1; i <= wins; i++ {
					cardEarned := parseCard(lines[card.Number-1+i])
					mutex.Lock()
					h.Push(cardEarned)
					acc += 1
					mutex.Unlock()
				}
			}(copies)
			wg.Add(1)
			go iterHeap(copies)
		}()
	}

	wg.Wait()
	fmt.Println(acc)

}
