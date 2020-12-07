package main

import (
	"fmt"
	"strings"

	"github.com/seiffert/advent-of-code/2020/lib"
)

func main() {
	input := lib.MustReadFile("input.txt")

	m := NewMap(input)

	result31 := m.CountTrees(Slope{3, 1})
	fmt.Printf("would encounter %d trees with a 3/1 slope\n", result31)

	result11 := m.CountTrees(Slope{1, 1})
	fmt.Printf("would encounter %d trees with a 1/1 slope\n", result11)

	result51 := m.CountTrees(Slope{5, 1})
	fmt.Printf("would encounter %d trees with a 5/1 slope\n", result51)

	result71 := m.CountTrees(Slope{7, 1})
	fmt.Printf("would encounter %d trees with a 7/1 slope\n", result71)

	result12 := m.CountTrees(Slope{1, 2})
	fmt.Printf("would encounter %d trees with a 1/2 slope\n", result12)

	fmt.Printf("%d * %d * %d * %d * %d = %d\n",
		result31, result11, result51, result71, result12,
		result31*result11*result51*result71*result12,
	)

}

type Map [][]bool
type Slope struct{ right, down int }

func NewMap(input string) Map {
	var (
		lines = strings.Split(input, "\n")
		m     = make(Map, 0, len(lines))
	)
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		row := make([]bool, 0, len(line))
		for _, char := range line {
			row = append(row, char == '#')
		}
		m = append(m, row)
	}
	return m
}

func (m Map) CountTrees(s Slope) int {
	for trees, x, y := 0, 0, 0; ; y += s.down {
		if y >= len(m) {
			return trees
		}
		if m[y][x%len(m[y])] {
			trees++
		}
		x += s.right
	}
}
