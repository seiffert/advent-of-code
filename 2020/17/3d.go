package main

import "strings"

type (
	Grid3D struct {
		XBounds Bounds
		YBounds Bounds
		ZBounds Bounds
		Layers  map[int]Layer
	}
	Layer  map[int]Row
	Row    map[int]Cube
	Cube   bool
	Bounds struct{ Min, Max int }
)

func New3DGrid(input string) *Grid3D {
	l := Layer{}
	for y, row := range strings.Split(input, "\n") {
		r := Row{}
		for x, cube := range row {
			r[x] = cube == '#'
		}
		l[y] = r
	}
	return &Grid3D{
		Layers:  map[int]Layer{0: l},
		XBounds: Bounds{0, len(l[0]) - 1},
		YBounds: Bounds{0, len(l) - 1},
		ZBounds: Bounds{0, 0},
	}
}

func (g *Grid3D) Play(cycles int) {
	for i := 0; i < cycles; i++ {
		g.play()
	}
}

func (g *Grid3D) play() {
	for z := g.ZBounds.Min - 1; z <= g.ZBounds.Max+1; z++ {
		for y := g.YBounds.Min - 1; y <= g.YBounds.Max+1; y++ {
			for x := g.XBounds.Min - 1; x <= g.XBounds.Max+1; x++ {
				numActiveNeighbors := 0
				for zi := -1; zi <= 1; zi++ {
					for yi := -1; yi <= 1; yi++ {
						for xi := -1; xi <= 1; xi++ {
							if zi == 0 && yi == 0 && xi == 0 {
								continue
							}
							if g.Layers[z+zi][y+yi][x+xi] {
								numActiveNeighbors++
							}
						}
					}
				}

				active := g.Layers[z][y][x]
				switch {
				case numActiveNeighbors == 2 && bool(active):
				case numActiveNeighbors == 3 && bool(active):
				case numActiveNeighbors == 3 && !bool(active):
					defer g.update(z, y, x, true)
				default:
					defer g.update(z, y, x, false)
				}

			}
		}
	}
}

func (g *Grid3D) update(z, y, x int, state Cube) {
	if _, ok := g.Layers[z]; !ok {
		g.Layers[z] = Layer{}
	}
	if _, ok := g.Layers[z][y]; !ok {
		g.Layers[z][y] = Row{}
	}

	g.Layers[z][y][x] = state

	g.XBounds.update(x)
	g.YBounds.update(y)
	g.ZBounds.update(z)
}

func (g *Grid3D) CountActive() int {
	var count int
	for z := g.ZBounds.Min; z <= g.ZBounds.Max; z++ {
		for y := g.YBounds.Min; y <= g.YBounds.Max; y++ {
			for x := g.XBounds.Min; x <= g.XBounds.Max; x++ {
				if g.Layers[z][y][x] {
					count++
				}
			}
		}
	}
	return count
}

func (b *Bounds) update(n int) {
	if n > b.Max {
		b.Max = n
	}
	if n < b.Min {
		b.Min = n
	}
}
