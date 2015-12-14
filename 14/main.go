package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

var (
	InputLineRegexp = regexp.MustCompile("^([^ ]+) can fly ([0-9]+) km/s for ([0-9]+) seconds, but then must rest for ([0-9]+) seconds.$")
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	time := i(os.Args[1])

	maxDist := 0
	maxDistDeer := ""
	for s.Scan() {
		l := s.Text()

		m := InputLineRegexp.FindAllStringSubmatch(l, -1)
		dist := calcDist(time, i(m[0][2]), i(m[0][3]), i(m[0][4]))
		if dist > maxDist {
			maxDistDeer = m[0][1]
			maxDist = dist
		}
	}

	fmt.Printf("%s ran %d km\n", maxDistDeer, maxDist)
}

func i(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func calcDist(time, speed, flytime, resttime int) int {
	period := resttime + flytime
	periods := int(math.Floor(float64(time / period)))

	return periods*speed*flytime + int(math.Min(float64(time%period), float64(flytime)))*speed
}
