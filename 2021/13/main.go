package main

import (
	"fmt"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	g, folds := NewPaper(lib.MustReadFile("input.txt"))

	for i, f := range folds {
		g.Fold(f)
		fmt.Printf("after fold %d: %d dots visible\n", i, g.CountDots())
	}

	fmt.Println()
	fmt.Println(g.String())
}

type Paper struct {
	dots map[Coord]bool
	max  Coord
}

type Coord struct{ X, Y int }

func NewPaper(in string) (Paper, []Coord) {
	p, folds := Paper{dots: map[Coord]bool{}}, []Coord{}

	var dotsRead bool
	for _, l := range strings.Split(in, "\n") {
		if strings.TrimSpace(l) == "" {
			dotsRead = true
		} else if !dotsRead {
			c := lib.MustAllInts(strings.Split(l, ","))
			p.dots[Coord{c[0], c[1]}] = true
			if c[0] >= p.max.X {
				p.max.X = c[0]
			}
			if c[1] >= p.max.Y {
				p.max.Y = c[1]
			}
		} else {
			s := strings.Split(strings.Trim(l, "fold ang"), "=")
			axis, xy := s[0], lib.MustInt(s[1])

			var fold Coord
			if axis == "x" {
				fold.X = xy
			} else {
				fold.Y = xy
			}

			folds = append(folds, fold)
		}
	}

	return p, folds
}

func (p Paper) Fold(fold Coord) {
	for c := range p.dots {
		switch {
		case fold.X == 0 && c.Y > fold.Y:
			delete(p.dots, c)
			p.dots[Coord{c.X, fold.Y + fold.Y - c.Y}] = true
		case fold.Y == 0 && c.X > fold.X:
			delete(p.dots, c)
			p.dots[Coord{fold.X + fold.X - c.X, c.Y}] = true
		}
	}
}

func (p Paper) CountDots() int {
	return len(p.dots)
}

func (p Paper) String() (out string) {
	for y := 0; y <= p.max.Y; y++ {
		for x := 0; x <= p.max.X; x++ {
			if p.dots[Coord{x, y}] {
				out += "#"
			} else {
				out += "."
			}
		}
		out += "\n"
	}
	return
}
