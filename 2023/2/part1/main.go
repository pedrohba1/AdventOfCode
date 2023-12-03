package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	Id   int
	Sets []Set
}

type Set struct {
	Colors map[string]int
}

func parseGame(line string) Game {
	gameRe := regexp.MustCompile(`(\d+ \w+)`)
	textRe := regexp.MustCompile(`red|blue|green`)
	numberRe := regexp.MustCompile(`\d+`)

	parts := strings.Split(line, ":")
	id, _ := strconv.Atoi(numberRe.FindString(parts[0]))

	game := Game{
		Id:   id,
		Sets: make([]Set, 0),
	}

	game.Id = id
	strSets := strings.Split(parts[1], ";")

	for _, strElem := range strSets {

		newSet := Set{
			Colors: make(map[string]int),
		}

		sets := gameRe.FindAllString(strElem, -1)

		for _, elem := range sets {
			color := textRe.FindString(elem)
			val, _ := strconv.Atoi(numberRe.FindString(elem))
			newSet.Colors[color] += val
		}
		game.Sets = append(game.Sets, newSet)
	}

	return game
}

func validSets(game Game) bool {
	for _, set := range game.Sets {
		if set.Colors["red"] > 12 || set.Colors["green"] > 13 || set.Colors["blue"] > 14 {
			return false
		}
	}
	return true
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

		game := parseGame(line)
		if validSets(game) {
			acc += game.Id
		}
	}
	fmt.Println(acc)
}
