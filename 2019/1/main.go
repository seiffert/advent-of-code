package main

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/2019/lib"
)

func FuelForMass(mass int) int {
	return mass/3 - 2
}

func FuelForMassAndFuel(mass int) int {
	fuel := FuelForMass(mass)

	var result int
	for fuel > 0 {
		result += fuel
		fuel = FuelForMass(fuel)
	}
	return result
}

func main() {
	in, err := ioutil.ReadFile("input.txt")
	if err != nil {
		lib.Abort("Error reading input file result: %w", err)
	}

	lines := strings.Split(string(in), "\n")

	var withoutFuel, withFuel int
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		mass, err := strconv.Atoi(line)
		if err != nil {
			lib.Abort("Not a number: %s", line)
		}

		withoutFuel += FuelForMass(mass)
		withFuel += FuelForMassAndFuel(mass)
	}

	println("FuelOfMass:", withoutFuel)
	println("FuelOfMassAndFuel:", withFuel)
}
