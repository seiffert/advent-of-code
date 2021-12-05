package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	fmt.Printf("without diagonals: found %d overlaps\n",
		NewGame(lib.MustReadFile("input.txt"), false).CountOverlaps(),
	)
	fmt.Printf("with diagonals: found %d overlaps\n",
		NewGame(lib.MustReadFile("input.txt"), true).CountOverlaps(),
	)
}

var lineRE = regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

type Game map[Coords]int
type Coords struct{ X, Y int }

func NewGame(in string, diagonals bool) Game {
	g := Game{}

	for _, l := range strings.Split(in, "\n") {
		matches := lib.MustAllInts(lineRE.FindStringSubmatch(l)[1:])
		from, to := Coords{matches[0], matches[1]}, Coords{matches[2], matches[3]}

		if !diagonals && from.X != to.X && from.Y != to.Y {
			continue
		}

		var xDiff, yDiff int
		if to.Y != from.Y {
			yDiff = (to.Y - from.Y) / int(math.Abs(float64(to.Y-from.Y)))
		}
		if to.X != from.X {
			xDiff = (to.X - from.X) / int(math.Abs(float64(to.X-from.X)))
		}

		for ; from != to; from.X, from.Y = from.X+xDiff, from.Y+yDiff {
			g[from]++
		}

		g[to]++
	}

	return g
}

func (g Game) CountOverlaps() (out int) {
	for _, f := range g {
		if f > 1 {
			out++
		}
	}
	return
}
