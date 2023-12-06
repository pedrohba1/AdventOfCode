package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"regexp"
)

type Race struct {
	time     big.Int
	distance big.Int
}

func ways(time, dist *big.Float) *big.Int {
	// Convert necessary values to big.Float
	four := big.NewFloat(4.0)
	two := big.NewFloat(2.0)

	// Calculate operand using big.Float
	// operand = sqrt(b*b - 4*a*c)
	a := big.NewFloat(1.0)
	b := new(big.Float).Neg(time)
	c := dist
	operand := new(big.Float).Sub(new(big.Float).Mul(b, b), new(big.Float).Mul(four, new(big.Float).Mul(a, c)))
	operand.Sqrt(operand)

	// Calculate r1 and r2
	r1 := new(big.Float).Quo(new(big.Float).Sub(new(big.Float).Neg(b), operand), new(big.Float).Mul(two, a))
	r2 := new(big.Float).Quo(new(big.Float).Add(new(big.Float).Neg(b), operand), new(big.Float).Mul(two, a))

	// Convert to big.Int for comparison and increment/decrement
	r1Int := new(big.Int)
	r1.Int(r1Int)
	r2Int := new(big.Int)
	r2.Int(r2Int)

	// Adjust bounds
	r1Ceil := new(big.Int).Add(r1Int, big.NewInt(1))
	r2Floor := new(big.Int).Sub(r2Int, big.NewInt(1))

	// Calculate number of ways
	ways := new(big.Int).Sub(r2Floor, r1Ceil)
	ways.Add(ways, big.NewInt(1))

	return ways
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
	itime := new(big.Int)
	itime, ok := itime.SetString(time, 10)
	idistance := new(big.Int)
	idistance, ok = idistance.SetString(distance, 10)

	if !ok {
		fmt.Println("SetString: error")
		return
	}
	races = append(races, Race{
		time:     *itime,
		distance: *idistance,
	})

	var result big.Int
	result.SetInt64(1)
	for _, race := range races {

		timeBig := new(big.Float).SetInt(&race.time)
		distanceBig := new(big.Float).SetInt(&race.distance)
		ways := ways(timeBig, distanceBig)
		result.Mul(&result, ways)
	}

	fmt.Println("Result:", &result)
}
