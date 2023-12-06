package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Race struct {
	time     int
	distance int
}

func ways(time float64, dist float64) int64 {
	a := 1.0
	b := -time
	c := dist
	operand := math.Sqrt(b*b - 4*a*c)

	r1 := ((-b - operand) / 2 * a)
	r2 := ((-b + operand) / 2 * a)

	if r1 == float64(int(r1)) {
		r1++
	}
	if r2 == float64(int(r2)) {
		r2--
	}
	// get the upper and lower limit of exclusive values
	r1 = math.Ceil(r1)  // lower limit
	r2 = math.Floor(r2) // upper limit
	return int64(r2 - r1 + 1)
}

func main() {
	var races []Race
	fmt.Println("- part 2 solution with bhaskara algorithm on big boi")

	file, err := os.Open("../bigboi.txt")
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

	var time string
	var distance string
	for i := 0; i < len(distances); i++ {
		time += times[i]
		distance += distances[i]

	}
	itime, _ := strconv.Atoi(time)
	idistance, _ := strconv.Atoi(distance)

	races = append(races, Race{
		time:     itime,
		distance: idistance,
	})

	// this is an attempt on using the bhaskara formula to solve the problem.
	// we can infer by a series of equantios
	// total_time = time_pressing + travel_time
	// velocity = time_pressing * 1
	// distance = travel_time * velocity
	// distance = (total_time - time_pressing) * velocity
	// distance = (total_time - time_pressing) * time_pressing
	// distance = total_time * time_pressing - time_pressing^2
	// which can be resumed to this for the solution
	// x = time_pressing
	// total_time * x - x ** 2 > distance
	// quadratic inequality:
	// x ^ 2 + (- total_time) * x + distance > 0

	// since for every case I have distance and travel time,
	// it turns into an equation 2nd degree equation
	// let's say for the race in the example:
	// 9 < 7 *time_pressing - time_pressing^2
	//0  -time_pressing^2 + 7*time_pressing -9

	// we can apply bhaskara's formula:
	// (-7 +/- sqrt((-7)^2 - 4*(-1) *(-9))/ 2*(-1)
	// and then we just need to floor the - root, and ceil the + root;
	// after that, just ge the integer values in between the ways.

	var result int64
	result = 1
	for _, race := range races {
		ways := ways(float64(race.time), float64(race.distance))
		result *= ways
	}

	fmt.Println(result)
}
