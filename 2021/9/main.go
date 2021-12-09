package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	m := NewHeightMap(lib.MustReadFile("input.txt"))

	fmt.Printf("sum of risk levels: %d\n", m.SumRiskLevels(m.LowPoints()))
	fmt.Printf("3 biggest sinks: %d\n", m.MultiplyBiggestBasinSizes(3))
}

type HeightMap map[Coord]int
type Coord struct{ X, Y int }

func NewHeightMap(in string) HeightMap {
	heightMap := make(HeightMap)
	for y, l := range strings.Split(in, "\n") {
		for x, c := range l {
			heightMap[Coord{x, y}] = lib.MustInt(string(c))
		}
	}
	return heightMap
}

func (hm HeightMap) LowPoints() (out []Coord) {
	for c, h := range hm {
		if above, ok := hm[Coord{c.X, c.Y - 1}]; ok && above <= h {
			continue
		}
		if below, ok := hm[Coord{c.X, c.Y + 1}]; ok && below <= h {
			continue
		}
		if left, ok := hm[Coord{c.X - 1, c.Y}]; ok && left <= h {
			continue
		}
		if right, ok := hm[Coord{c.X + 1, c.Y}]; ok && right <= h {
			continue
		}

		out = append(out, c)
	}
	return
}

func (hm HeightMap) SumRiskLevels(ps []Coord) (out int) {
	for _, p := range ps {
		out += hm[p] + 1
	}
	return
}

func (hm HeightMap) MultiplyBiggestBasinSizes(n int) int {
	lps := hm.LowPoints()

	var basins []int
	for _, lp := range lps {
		basins = append(basins, hm.basinSize(lp))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))

	return basins[0] * basins[1] * basins[2]
}

func (hm HeightMap) basinSize(c Coord) (out int) {
	basin := map[Coord]bool{}
	todo := []Coord{c}

	for len(todo) != 0 {
		c, todo = todo[0], todo[1:]
		if basin[c] {
			continue
		}

		basin[c] = true

		if above, ok := hm[Coord{c.X, c.Y - 1}]; ok && above >= hm[c] && above != 9 {
			todo = append(todo, Coord{c.X, c.Y - 1})
		}
		if below, ok := hm[Coord{c.X, c.Y + 1}]; ok && below >= hm[c] && below != 9 {
			todo = append(todo, Coord{c.X, c.Y + 1})
		}
		if left, ok := hm[Coord{c.X - 1, c.Y}]; ok && left >= hm[c] && left != 9 {
			todo = append(todo, Coord{c.X - 1, c.Y})
		}
		if right, ok := hm[Coord{c.X + 1, c.Y}]; ok && right >= hm[c] && right != 9 {
			todo = append(todo, Coord{c.X + 1, c.Y})
		}
	}

	return len(basin)
}
