package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Race struct {
	time     int
	distance int
}

func main() {
	var races []Race

	file, err := os.Open("../input.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	scanner := bufio.NewScanner(file)

	intPattern := regexp.MustCompile(`\d+`)
	var times []string
	var distances []string
	for i := 1; scanner.Scan(); i++ {
		if i == 1 {
			times = intPattern.FindAllString(scanner.Text(), -1)
		} else {
			distances = intPattern.FindAllString(scanner.Text(), -1)
		}
	}

	for i := 0; i < len(distances); i++ {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])

		races = append(races, Race{
			time:     time,
			distance: distance,
		})
	}

	result := 1
	for _, race := range races {
		waysToWin := 0
		for holdTime := 0; holdTime < race.time; holdTime++ {
			speed := holdTime
			travelTime := race.time - holdTime
			if speed*travelTime > race.distance {
				waysToWin++
			}
		}
		result *= waysToWin
	}

	fmt.Println("Total ways to win all races:", result)
}
