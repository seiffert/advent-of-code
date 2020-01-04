package main

import (
	"fmt"
	"math"
	"strings"
)

func NewAsteroidField(in string) *AsteroidField {
	f := &AsteroidField{
		asteroids: make(map[Coord]*Asteroid),
	}
	rows := strings.Split(in, "\n")
	for y, row := range rows {
		if len(strings.TrimSpace(row)) == 0 {
			continue
		}

		for x, location := range row {
			if location == '#' {
				f.asteroids[Coord{x, y}] = &Asteroid{Coord{x, y}}
				f.maxX = int(math.Max(float64(x), float64(f.maxX)))
				f.maxY = int(math.Max(float64(y), float64(f.maxY)))
			}
		}
	}
	return f
}

type AsteroidField struct {
	asteroids  map[Coord]*Asteroid
	maxX, maxY int
}

func (f *AsteroidField) Find(cmp func(a, b *Asteroid) bool) *Asteroid {
	var max *Asteroid
	for _, a := range f.asteroids {
		if max == nil || cmp(a, max) {
			max = a
		}
	}
	return max
}

type Asteroid struct {
	Coord
}

func (a *Asteroid) String() string {
	return a.Coord.String()
}

func (a *Asteroid) AsteroidsInSight(field *AsteroidField) int {
	var (
		blockedDirs = map[Coord][]float32{}
		visited     = map[Coord]bool{
			a.Coord: true,
		}
		inSight int
		maxDist = int(math.Max(
			float64(a.x+a.y),
			float64(field.maxX+field.maxY),
		))
	)

	for dist := 1; dist <= maxDist; dist++ {
		for _, xDir := range []int{-1, 0, 1} {
			for _, yDir := range []int{-1, 0, 1} {
				for xDist := 0; xDist <= dist; xDist++ {
					yDist := dist - xDist

					coords := Coord{a.x + xDir*xDist, a.y + yDir*yDist}
					if visited[coords] {
						continue
					}
					visited[coords] = true

					if _, ok := field.asteroids[coords]; ok {
						dir := float32(yDist) / float32(xDist)

						var blocked bool
						for _, blockedDir := range blockedDirs[Coord{xDir, yDir}] {
							if blockedDir == dir {
								blocked = true
							}
						}

						if !blocked {
							blockedDirs[Coord{xDir, yDir}] = append(blockedDirs[Coord{xDir, yDir}], dir)
							inSight++
						}
					}
				}
			}
		}
	}

	return inSight
}

type Coord struct {
	x, y int
}

func (c Coord) String() string {
	return fmt.Sprintf("(%d/%d)", c.x, c.y)
}
