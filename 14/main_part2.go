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

	deers := []Deer{}
	for s.Scan() {
		l := s.Text()

		m := InputLineRegexp.FindAllStringSubmatch(l, -1)
		deers = append(deers, Deer{
			Name:     m[0][1],
			Speed:    i(m[0][2]),
			FlyTime:  i(m[0][3]),
			RestTime: i(m[0][4]),
		})
	}
	points := map[Deer]int{}
	for t := 1; t <= time; t++ {
		positions := map[Deer]int{}
		for _, d := range deers {
			positions[d] = calcPos(t, d)
		}
		maxPos := 0
		maxPosDeers := []Deer{}
		for d, p := range positions {
			if maxPos == p {
				maxPosDeers = append(maxPosDeers, d)
			} else if maxPos < p {
				maxPosDeers = []Deer{d}
				maxPos = p
			}
		}
		for _, d := range maxPosDeers {
			points[d]++
		}
	}

	maxPoints := 0
	var maxPointsDeer Deer
	for d, p := range points {
		if p > maxPoints {
			maxPoints = p
			maxPointsDeer = d
		}
	}

	fmt.Printf("%s got %d points", maxPointsDeer.Name, maxPoints)
}

func calcPos(t int, d Deer) int {
	period := d.RestTime + d.FlyTime
	periods := int(math.Floor(float64(t / period)))

	return periods*d.Speed*d.FlyTime + int(math.Min(float64(t%period), float64(d.FlyTime)))*d.Speed
}

type Deer struct {
	Name     string
	Speed    int
	FlyTime  int
	RestTime int
}

func i(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
