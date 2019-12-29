package main

import (
	"math"
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/2019/lib"
)

type (
	Grid struct {
		cells map[coords]*cell
		wires map[string]*wire

		topRight   coords
		bottomLeft coords
	}
	wire struct {
		end    coords
		length int
	}
	cell struct {
		wires map[string]int
	}
	coords struct {
		x, y int
	}
)

func NewGrid() *Grid {
	return &Grid{
		cells: make(map[coords]*cell),
		wires: make(map[string]*wire),

		topRight:   coords{0, 0},
		bottomLeft: coords{0, 0},
	}
}

func (g *Grid) AddWire(id, instructions string) {
	g.wires[id] = &wire{
		end: coords{0, 0},
	}

	lines := strings.Split(instructions, ",")
	for _, line := range lines {
		g.addLine(id, line)
	}
}

func (g *Grid) addLine(id, instruction string) {
	wire := g.wires[id]
	dir, steps := parseInstruction(instruction)

	switch dir {
	case 'U':
		for i := 1; i <= steps; i++ {
			g.add(id, coords{wire.end.x, wire.end.y + i}, wire.length+i)
		}
		wire.end = coords{wire.end.x, wire.end.y + steps}
	case 'D':
		for i := 1; i <= steps; i++ {
			g.add(id, coords{wire.end.x, wire.end.y - i}, wire.length+i)
		}
		wire.end = coords{wire.end.x, wire.end.y - steps}
	case 'R':
		for i := 1; i <= steps; i++ {
			g.add(id, coords{wire.end.x + i, wire.end.y}, wire.length+i)
		}
		wire.end = coords{wire.end.x + steps, wire.end.y}
	case 'L':
		for i := 1; i <= steps; i++ {
			g.add(id, coords{wire.end.x - i, wire.end.y}, wire.length+i)
		}
		wire.end = coords{wire.end.x - steps, wire.end.y}
	}
	wire.length += steps
}

func (g *Grid) add(wire string, pos coords, steps int) {
	_, ok := g.cells[pos]
	if !ok {
		g.cells[pos] = newCell()

		g.bottomLeft = coords{
			int(math.Min(float64(g.bottomLeft.x), float64(pos.x))),
			int(math.Min(float64(g.bottomLeft.y), float64(pos.y))),
		}
		g.topRight = coords{
			int(math.Max(float64(g.topRight.x), float64(pos.x))),
			int(math.Max(float64(g.topRight.y), float64(pos.y))),
		}
	}

	g.cells[pos].add(wire, steps)
}

func (g *Grid) FindClosestIntersectionDistance() int {
	var r coords
	for pos, cell := range g.cells {
		if !cell.isIntersection() {
			continue
		}
		if r.zero() || r.distance() > pos.distance() {
			r = pos
		}
	}
	return r.distance()
}

func (g *Grid) FindShortestIntersectingPath() int {
	var r *cell
	for _, cell := range g.cells {
		if !cell.isIntersection() {
			continue
		}
		if r == nil || r.sumPathLengths() > cell.sumPathLengths() {
			r = cell
		}
	}
	return r.sumPathLengths()
}

func newCell() *cell {
	return &cell{
		wires: make(map[string]int),
	}
}

func (c *cell) add(wire string, steps int) {
	if _, ok := c.wires[wire]; !ok {
		c.wires[wire] = steps
	}
}

func (c *cell) isIntersection() bool {
	return len(c.wires) == 2
}

func (c *cell) sumPathLengths() int {
	var sum int
	for _, l := range c.wires {
		sum += l
	}
	return sum
}

func (p coords) distance() int {
	return int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
}

func (p coords) zero() bool {
	return p.x == 0 && p.y == 0
}

func parseInstruction(instruction string) (byte, int) {
	steps, err := strconv.Atoi(instruction[1:])
	if err != nil {
		lib.Abort("invalid instruction %q: %w", instruction, err)
	}
	return instruction[0], steps
}
