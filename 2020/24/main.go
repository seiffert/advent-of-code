package main

import (
	"fmt"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	g := NewHexGrid()
	for _, instr := range strings.Split(lib.MustReadFile("input.txt"), "\n") {
		g.Flip(instr)
	}
	fmt.Println("number of black tiles:", g.Count())

	for i := 0; i < 100; i++ {
		g.Evolute()
	}
	fmt.Println("after 100 days:", g.Count())
}

type HexGrid struct {
	BoolGrid3D
}

func NewHexGrid() *HexGrid {
	return &HexGrid{make(BoolGrid3D)}
}

func (g HexGrid) Flip(instr string) {
	var x, y, z int
	for i := 0; i < len(instr); i++ {
		switch instr[i] {
		case 'e':
			x++
			y--
		case 'w':
			x--
			y++
		case 's':
			i++
			if instr[i] == 'e' {
				z++
				y--
				continue
			}
			x--
			z++
		case 'n':
			i++
			if instr[i] == 'e' {
				z--
				x++
				continue
			}
			z--
			y++
		}
	}
	g.BoolGrid3D[coords{x, y, z}] = !g.BoolGrid3D[coords{x, y, z}]
}

func (g HexGrid) Evolute() {
	stack := make([]coords, 0, len(g.BoolGrid3D))
	for c := range g.BoolGrid3D {
		stack = append(stack, c)
	}

	newGrid := BoolGrid3D{}
	for len(stack) > 0 {
		c := stack[0]
		stack = stack[1:]
		if _, ok := newGrid[c]; ok {
			continue
		}

		var ns int
		n := neighbors(c)
		for _, nc := range n {
			if g.BoolGrid3D[nc] {
				ns++
			} else if g.BoolGrid3D[c] {
				if _, ok := newGrid[nc]; !ok {
					stack = append(stack, nc)
				}
			}
		}

		switch {
		case g.BoolGrid3D[c] && (ns == 0 || ns > 2):
			newGrid[c] = false
		case !g.BoolGrid3D[c] && ns == 2:
			newGrid[c] = true
		default:
			newGrid[c] = g.BoolGrid3D[c]
		}
	}
	for c, b := range newGrid {
		g.BoolGrid3D[c] = b
	}
}

type BoolGrid3D map[coords]bool
type coords struct{ x, y, z int }

func (g BoolGrid3D) Count() int {
	var count int
	for _, b := range g {
		if b {
			count++
		}
	}
	return count
}

func neighbors(c coords) []coords {
	return []coords{
		{c.x + 1, c.y - 1, c.z},
		{c.x + 1, c.y, c.z - 1},
		{c.x, c.y - 1, c.z + 1},
		{c.x, c.y + 1, c.z - 1},
		{c.x - 1, c.y, c.z + 1},
		{c.x - 1, c.y + 1, c.z},
	}
}
