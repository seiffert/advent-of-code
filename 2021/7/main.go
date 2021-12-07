package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	fmt.Printf("required fuel under linear usage: %d\n", RequiredFuel(lib.MustReadFile("input.txt"), true))
	fmt.Printf("required fuel under increasing usage: %d\n", RequiredFuel(lib.MustReadFile("input.txt"), false))
}

func RequiredFuel(in string, linear bool) int {
	positions := lib.MustAllInts(strings.Split(in, ","))

	min, max := math.MaxInt, math.MinInt
	for _, pos := range positions {
		if pos < min {
			min = pos
		}
		if pos > max {
			max = pos
		}
	}

	minFuelRequired := math.MaxInt
	for center := min; center <= max; center++ {
		if f := fuelRequired(positions, center, linear); f < minFuelRequired {
			minFuelRequired = f
		}
	}

	return minFuelRequired
}

func fuelRequired(crabs []int, pos int, linear bool) (out int) {
	for _, c := range crabs {
		dist := int(math.Abs(float64(c - pos)))

		if linear {
			out += dist
			continue
		}

		out += dist * (dist + 1) / 2
	}
	return
}
