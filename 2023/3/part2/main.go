package main

import (
	"bufio"
	"fmt"
	"os"
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
	i int
	j int
	r rune
}

func checkSurroundings(i, j int, lines [][]rune) int {

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

	return 0
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
			checkSurroundings(i, j, lines)
		}
	}

	fmt.Println(acc)
}
