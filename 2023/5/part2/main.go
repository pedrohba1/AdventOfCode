package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type AgroMap struct {
	matrix [][]int
}

func (a AgroMap) findDest(val int) int {
	for i := 0; i < len(a.matrix); i++ {
		rg := a.matrix[i][2]
		source := a.matrix[i][1]
		end := a.matrix[i][0]
		if val >= source && val <= source+rg {
			return (val - source) + end
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
func main() {
	file, err := os.Open("../example.txt")
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

	// acc := 0
	var allSeeds []int
	for i, seed := range seeds {
		if i%2 == 0 {
			for j := 0; j <= seed+seeds[i+1]; j++ {
				allSeeds = append(allSeeds, seed+j)
			}
		}
	}
	fmt.Println(allSeeds)

	fmt.Println("seeds: ", allSeeds)
	finalLocation := math.MaxInt64
	for _, seed := range seeds {
		soil := agroMaps["seed-to-soil"].findDest(seed)
		fert := agroMaps["soil-to-fertilizer"].findDest(soil)
		water := agroMaps["fertilizer-to-water"].findDest(fert)
		light := agroMaps["water-to-light"].findDest(water)
		temp := agroMaps["light-to-temperature"].findDest(light)
		humidity := agroMaps["temperature-to-humidity"].findDest(temp)
		location := agroMaps["humidity-to-location"].findDest(humidity)
		if location < finalLocation {
			finalLocation = location
		}
	}

	fmt.Println(finalLocation)
}
