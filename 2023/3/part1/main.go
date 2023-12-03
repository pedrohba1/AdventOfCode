package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func findSpecialChar(line []rune) (rune, int) {
	for idx, r := range line {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '.' {
			return r, idx
		}
	}
	return 0, -1
}

type Position struct {
	i          int
	j          int
	r          rune
	fullNumber int
}

type SliceType interface {
	~string | ~int | ~float64 // add more *comparable* types as needed
}

func removeDuplicateValues(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func inRange(i, j int, m [][]rune) bool {
	if i < 0 || i >= len(m) {
		return false
	}
	if j < 0 || j >= len(m[i]) {
		return false
	}
	return true
}

func fulfillNumber(p *Position, lines [][]rune) {
	numberStr := []rune{}
	for k := 0; ; k++ {
		if !inRange(p.i, p.j-k, lines) {
			break
		}
		r := lines[p.i][p.j-k]
		if !unicode.IsDigit(r) {
			break
		}
		numberStr = append([]rune{r}, numberStr...)
	}

	for k := 1; ; k++ {
		if !inRange(p.i, p.j+k, lines) {
			break
		}
		r := lines[p.i][p.j+k]
		if !unicode.IsDigit(r) {
			break
		}
		numberStr = append(numberStr, r)
	}

	res, _ := strconv.Atoi(string(numberStr))

	*&p.fullNumber = res
}

func calcSquare(i, j int, lines [][]rune) int {

	var toExpand []Position
	for _, p := range []Position{
		{i: i + 1, j: j + 1, r: lines[i+1][j+1]},
		{i: i - 1, j: j + 1, r: lines[i-1][j+1]},
		{i: i - 1, j: j - 1, r: lines[i-1][j-1]},
		{i: i + 1, j: j - 1, r: lines[i+1][j-1]},
		{i: i, j: j - 1, r: lines[i][j-1]},
		{i: i, j: j + 1, r: lines[i][j+1]},
		{i: i + 1, j: j, r: lines[i+1][j]},
		{i: i - 1, j: j, r: lines[i-1][j]},
	} {
		if unicode.IsDigit(p.r) {
			toExpand = append(toExpand, p)
		}
	}

	expanded := []int{}
	for _, p := range toExpand {
		fulfillNumber(&p, lines)
		expanded = append(expanded, p.fullNumber)

	}

	expanded = removeDuplicateValues(expanded)
	fmt.Println(expanded)

	sum := 0
	for _, number := range expanded {
		sum += number
	}

	return sum
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
	var lines [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		rs := []rune(line)
		lines = append(lines, rs)
	}

	for i, line := range lines {
		r, j := findSpecialChar(line)
		if r != 0 {
			acc += calcSquare(i, j, lines)
		}
	}

	fmt.Print(acc)
}
