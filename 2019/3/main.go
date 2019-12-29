package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/2019/lib"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		lib.Abort("error reading input file: %w", err)
	}

	inputLines := strings.Split(string(input), "\n")
	g := NewGrid()
	for i, line := range inputLines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		g.AddWire(strconv.Itoa(i), line)
	}

	fmt.Println("distance to closest intersection",
		g.FindClosestIntersectionDistance(),
	)

	fmt.Println("shortest path to intersection",
		g.FindShortestIntersectingPath(),
	)
}
