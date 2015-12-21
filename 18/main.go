package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strconv"
)

const (
	LightOn  = '#'
	LightOff = '.'
)

type Grid []GridRow
type GridRow []byte

func main() {
	s := bufio.NewScanner(os.Stdin)
	var (
		iterations int
		grid       Grid
	)
	iterations, _ = strconv.Atoi(os.Args[1])

	i := 0
	for s.Scan() {
		row := s.Text()
		if grid == nil {
			grid = make(Grid, len(row))
		}
		grid[i] = GridRow(row)
		i++
	}

	for i = 0; i < iterations; i++ {
		grid[0][0], grid[0][len(grid[0])-1], grid[len(grid)-1][len(grid[len(grid)-1])-1], grid[len(grid)-1][0] = LightOn, LightOn, LightOn, LightOn
		grid = iterate(grid)
		grid[0][0], grid[0][len(grid[0])-1], grid[len(grid)-1][len(grid[len(grid)-1])-1], grid[len(grid)-1][0] = LightOn, LightOn, LightOn, LightOn
		log.Printf("After %d iterations, %d lights are switched on.", i+1, grid.SumOn())
	}
}

func iterate(grid Grid) Grid {
	newGrid := make(Grid, len(grid))
	for i := range grid {
		newGrid[i] = make(GridRow, len(grid[i]))
		for j := range grid[i] {
			surrounding := countSurroundingLights(grid, i, j)
			newGrid[i][j] = grid[i][j]
			if grid.On(i, j) && surrounding != 2 && surrounding != 3 {
				newGrid[i][j] = LightOff
			} else if grid.Off(i, j) && surrounding == 3 {
				newGrid[i][j] = LightOn
			}
		}
	}

	return newGrid
}

func countSurroundingLights(grid Grid, i, j int) int {
	c := 0
XLoop:
	for x := i - 1; x <= i+1; x++ {
	YLoop:
		for y := j - 1; y <= j+1; y++ {
			if x < 0 || x >= len(grid) {
				continue XLoop
			}
			if y < 0 || y >= len(grid[x]) {
				continue YLoop
			}
			if x == i && y == j {
				continue
			}
			if grid.On(x, y) {
				c++
			}
		}
	}
	return c
}

func (g Grid) SumOn() int {
	c := 0
	for _, row := range g {
		rc := row.SumOn()
		c += rc
	}
	return c
}
func (g Grid) On(x, y int) bool {
	return g[x][y] == LightOn
}
func (g Grid) Off(x, y int) bool {
	return g[x][y] == LightOff
}
func (g Grid) String() string {
	buf := bytes.Buffer{}
	for _, row := range g {
		buf.Write(append([]byte(row), '\n'))
	}
	return buf.String()
}

func (r GridRow) SumOn() int {
	c := 0
	for _, light := range r {
		if light == LightOn {
			c++
		}
	}
	return c
}
