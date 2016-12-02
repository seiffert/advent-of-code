package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coords struct {
	X, Y int
}
func (c *Coords) Distance() int {
	return int(math.Abs(float64(c.X)) + math.Abs(float64(c.Y)))
}
func (c *Coords) Step(dir int) {
	switch dir {
	case 0:
		c.Y++
	case 1:
		c.X++
	case 2:
		c.Y--
	case 3:
		c.X--
	}
}

type City map[Coords]bool

func (c City) Visited(coords Coords) bool {
	return c[coords]
}
func (c City) Visit(coords Coords) {
	c[coords] = true
}

func CalculateDistance(input string) int {
	var (
		city         = &City{}
		coord        = &Coords{}
		dir          int
		instructions = strings.Split(input, ", ")
	)
outerLoop:
	for _, i := range instructions {
		turn, numS := i[0], i[1:]
		num, _ := strconv.Atoi(numS)
		if turn == 'R' {
			dir = (dir + 1) % 4
		} else {
			dir -= 1
			if dir < 0 {
				dir = 4 + dir
			}
		}
		for step := 0; step < num; step++ {
			coord.Step(dir)
			if city.Visited(*coord) {
				break outerLoop
			}
			city.Visit(*coord)
		}
	}
	return coord.Distance()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		t := scanner.Text()
		fmt.Printf("End: %d\n", CalculateDistance(t))
	}
}