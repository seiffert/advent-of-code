package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	input := lib.MustReadFile("input.txt")
	fmt.Println("Required fuel for modules:", CalcFuelSum(input, false))
	fmt.Println("Total required fuel:", CalcFuelSum(input, true))
}

func CalcFuelSum(input string, countFuel bool) (result int) {
	for _, l := range strings.Split(input, "\n") {
		m, _ := strconv.Atoi(l)
		result += CalcFuel(m, countFuel)
	}
	return
}
func CalcFuel(m int, countFuel bool) int {
	fuel := int(math.Floor(float64(m)/float64(3))) - 2
	if countFuel {
		fuel += CalcFuelFuel(fuel)
	}
	return fuel
}
func CalcFuelFuel(fuel int) (result int) {
	for {
		fuel = CalcFuel(fuel, false)
		if fuel <= 0 {
			return
		}
		result += fuel
	}
}
