package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
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
		panic(err)
	}

	lines := strings.Split(string(in), "\n")

	var withoutFuel, withFuel int
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		mass, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Sprintf("not a number: %s", line))
		}

		withoutFuel += FuelForMass(mass)
		withFuel += FuelForMassAndFuel(mass)
	}

	println("FuelOfMass:", withoutFuel)
	println("FuelOfMassAndFuel:", withFuel)
}
