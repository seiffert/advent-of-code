package main

import (
	"strings"
)

type (
	Grid4D struct {
		WBounds Bounds
		XBounds Bounds
		YBounds Bounds
		ZBounds Bounds
		Grids   map[int]Grid
	}
	Grid map[int]Layer
)

func New4DGrid(input string) *Grid4D {
	l := Layer{}
	for y, row := range strings.Split(input, "\n") {
		r := Row{}
		for x, cube := range row {
			r[x] = cube == '#'
		}
		l[y] = r
	}
	return &Grid4D{
		Grids:   map[int]Grid{0: Grid{0: l}},
		WBounds: Bounds{0, 0},
		XBounds: Bounds{0, len(l[0]) - 1},
		YBounds: Bounds{0, len(l) - 1},
		ZBounds: Bounds{0, 0},
	}
}

func (g *Grid4D) Play(cycles int) {
	for i := 0; i < cycles; i++ {
		g.play()
	}
}

func (g *Grid4D) play() {
	for w := g.WBounds.Min - 1; w <= g.WBounds.Max+1; w++ {
		for z := g.ZBounds.Min - 1; z <= g.ZBounds.Max+1; z++ {
			for y := g.YBounds.Min - 1; y <= g.YBounds.Max+1; y++ {
				for x := g.XBounds.Min - 1; x <= g.XBounds.Max+1; x++ {
					numActiveNeighbors := 0
					for wi := -1; wi <= 1; wi++ {
						for zi := -1; zi <= 1; zi++ {
							for yi := -1; yi <= 1; yi++ {
								for xi := -1; xi <= 1; xi++ {
									if wi == 0 && zi == 0 && yi == 0 && xi == 0 {
										continue
									}
									if g.Grids[w+wi][z+zi][y+yi][x+xi] {
										numActiveNeighbors++
									}
								}
							}
						}
					}

					active := g.Grids[w][z][y][x]
					switch {
					case numActiveNeighbors == 2 && bool(active):
					case numActiveNeighbors == 3 && bool(active):
					case numActiveNeighbors == 3 && !bool(active):
						defer g.update(w, z, y, x, true)
					default:
						defer g.update(w, z, y, x, false)
					}
				}
			}
		}
	}
}

func (g *Grid4D) update(w, z, y, x int, state Cube) {
	if _, ok := g.Grids[w]; !ok {
		g.Grids[w] = Grid{}
	}
	if _, ok := g.Grids[w][z]; !ok {
		g.Grids[w][z] = Layer{}
	}
	if _, ok := g.Grids[w][z][y]; !ok {
		g.Grids[w][z][y] = Row{}
	}

	g.Grids[w][z][y][x] = state

	g.WBounds.update(w)
	g.XBounds.update(x)
	g.YBounds.update(y)
	g.ZBounds.update(z)
}

func (g *Grid4D) CountActive() int {
	var count int
	for _, gr := range g.Grids {
		for _, l := range gr {
			for _, r := range l {
				for _, c := range r {
					if c {
						count++
					}
				}
			}
		}
	}
	return count
}
