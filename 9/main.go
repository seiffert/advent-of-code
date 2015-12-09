package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	InputLineRegexp = regexp.MustCompile("^([a-zA-Z]+) to ([a-zA-Z]+) = ([0-9]+)$")
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		m := InputLineRegexp.FindAllStringSubmatch(scanner.Text(), -1)
		if len(m) < 1 {
			continue
		}

		setDist(city(m[0][1]), city(m[0][2]), i(m[0][3]))
	}

	t := &Trip{}

	shortestTrip := shortestTripViaAllCities(t)
	fmt.Printf("shortest: %s\n", shortestTrip.String())

	t = &Trip{}
	longestTrip := longestTripViaAllCities(t)
	fmt.Printf("longest: %s\n", longestTrip.String())
}

func i(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
