package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type AgroMap struct {
	matrix [][]int
}

func (a AgroMap) findSrc(val int) int {
	for i := 0; i < len(a.matrix); i++ {
		rg := a.matrix[i][2]
		source := a.matrix[i][1]
		end := a.matrix[i][0]
		if val >= end && val <= end+rg {
			return (val - end) + source
		}
	}
	return val
}

func parseSeeds(line string) []int {
	parts := strings.Fields(line)
	integers := []int{}

	for _, part := range parts {
		// Check if the part can be converted to an integer
		if integer, err := strconv.Atoi(part); err == nil {
			integers = append(integers, integer)
		}
	}
	return integers
}

type Loc struct {
	val   int
	shift int
}

func InsertSorted(s []Loc, e Loc) []Loc {
	i := sort.Search(len(s), func(i int) bool { return s[i].val > e.val })
	s = append(s, Loc{}) // Append an empty struct
	copy(s[i+1:], s[i:])
	s[i] = e
	return s
}

func main() {
	file, err := os.Open("../input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	var seeds []int
	mapPattern := regexp.MustCompile(`(\w+-to-\w+)`)
	intPattern := regexp.MustCompile(`\d+`)

	agroMaps := make(map[string]AgroMap)
	currentMap := ""
	for idx, line := range lines {
		if idx == 0 {
			seeds = parseSeeds(line)
			continue
		}
		matches := mapPattern.FindAllString(line, -1)
		if len(matches) > 0 {
			currentMap = matches[0]
			agroMaps[currentMap] = AgroMap{}
		}
		matches = intPattern.FindAllString(line, -1)
		if len(matches) > 0 {
			var nMatches = []int{}
			for _, i := range matches {
				j, err := strconv.Atoi(i)
				if err != nil {
					panic(err)
				}
				nMatches = append(nMatches, j)
			}
			am := agroMaps[currentMap]
			am.matrix = append(am.matrix, nMatches)
			agroMaps[currentMap] = am
		}
	}

	var locRanges []Loc
	for _, mp := range agroMaps["humidity-to-location"].matrix {
		locRanges = InsertSorted(locRanges, Loc{val: mp[0], shift: mp[2]})
	}
	// fmt.Println(locRanges)

	for loc := 0; loc <= locRanges[len(locRanges)-1].val; loc++ {
		// fmt.Println("loc", loc)

		humidity := agroMaps["humidity-to-location"].findSrc(loc)
		temp := agroMaps["temperature-to-humidity"].findSrc(humidity)
		light := agroMaps["light-to-temperature"].findSrc(temp)
		water := agroMaps["water-to-light"].findSrc(light)
		fert := agroMaps["fertilizer-to-water"].findSrc(water)
		soil := agroMaps["soil-to-fertilizer"].findSrc(fert)
		seed := agroMaps["seed-to-soil"].findSrc(soil)

		for k := 0; k < len(seeds)-1; k++ {

			if k%2 == 0 && seed >= seeds[k] && seed <= seeds[k]+seeds[k+1]-1 {
				fmt.Println(loc) //99751241
				return
			}
		}
	}

	// fmt.Println(finalLoc)
}
