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

	//TODO: OPTMIZE
	// 	store your seed maps as ranges
	// when part of a range gets mapped to something new, slice out the old part of the range and create a new mapped range
	// Easy: just brute force it if your code is fast.
	// Hard: Sort the maps by source in normal ascending order.
	//  There are six cases you need to consider to either do nothing,
	// stop checking map ranges, modify the seed range,
	//  split it once below, once above,
	// or even both below and above. Make sure to update the upper part of the seed range so the next map range can work with it.
	// visual explanation on https://www.reddit.com/r/adventofcode/comments/18b9ohu/comment/kc2ybz6/
	//pseudo code:

	// for each remap in the game
	// for each remap segment
	//     new domain <- []
	//     for each domain segment
	//         compute overlap (between remap segment and domain segment) and remainder (of domain segment)
	//         add overlap (with remap segment's offset) and remainder (zero offset) to new domain
	//     domain <- new domain
	// apply offsets to domain

	// 	// another explanation
	// 	if you have seeds 1-4 (1,2,3,4),
	// and your map says seeds 3-9 (3,4,5,6,7,8,9) should turn into soil 13-19 (13,14,15,16,17,18,19),
	// then you now have a soil range of 13-14 (seeds 3,4 get turned into 13,14),
	// and a seed range of 1-2 (these did not get turned into soil).

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
