package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	s := NewShip(lib.MustReadFile("input.txt"))
	s.Drive()
	fmt.Printf("1) manhatten dist: %d\n", s.ManhattenDist())

	s = NewShip(lib.MustReadFile("input.txt"))
	s.DriveWithWaypoint()
	fmt.Printf("2) manhatten dist: %d\n", s.ManhattenDist())
}

type (
	Ship struct {
		Pos   Coords
		Angle int
		instr []instruction
		wayp  Coords
	}
	Coords      struct{ x, y int }
	instruction struct {
		action byte
		value  int
	}
)

func NewShip(input string) *Ship {
	s := &Ship{wayp: Coords{10, -1}}
	for _, line := range strings.Split(input, "\n") {
		val, _ := strconv.Atoi(line[1:])
		s.instr = append(s.instr, instruction{line[0], val})
	}
	return s
}

func (s *Ship) Drive() {
	for _, i := range s.instr {
		switch i.action {
		case 'N':
			s.Pos.y -= i.value
		case 'S':
			s.Pos.y += i.value
		case 'E':
			s.Pos.x += i.value
		case 'W':
			s.Pos.x -= i.value
		case 'L':
			s.Angle += i.value
		case 'R':
			s.Angle -= i.value
		case 'F':
			sin, cos := math.Sincos(float64(s.Angle) / 180 * math.Pi)
			s.Pos.x += int(cos) * i.value
			s.Pos.y -= int(sin) * i.value
		}
	}
}

func (s *Ship) DriveWithWaypoint() {
	for _, i := range s.instr {
		switch i.action {
		case 'N':
			s.wayp.y -= i.value
		case 'S':
			s.wayp.y += i.value
		case 'E':
			s.wayp.x += i.value
		case 'W':
			s.wayp.x -= i.value
		case 'L':
			s.wayp.Rotate(-i.value)
		case 'R':
			s.wayp.Rotate(i.value)
		case 'F':
			s.Pos.x += i.value * s.wayp.x
			s.Pos.y += i.value * s.wayp.y
		}
	}
}

func (s *Ship) ManhattenDist() int {
	return int(math.Abs(float64(s.Pos.x)) + math.Abs(float64(s.Pos.y)))
}

func (p *Coords) Rotate(deg int) {
	sin, cos := math.Sincos(float64(deg) / 180 * math.Pi)
	x, y := p.x, p.y
	p.x = int(cos)*x - int(sin)*y
	p.y = int(sin)*x + int(cos)*y
}
