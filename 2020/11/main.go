package main

import (
	"fmt"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	sm := NewSeatMap(lib.MustReadFile("input.txt"))
	occupiedSeats := sm.RunSimulation(4, true)
	fmt.Printf("%d are occupied after the simulations\n", occupiedSeats)

	sm = NewSeatMap(lib.MustReadFile("input.txt"))
	occupiedSeats = sm.RunSimulation(5, false)
	fmt.Printf("%d are occupied after the second simulations\n", occupiedSeats)
}

type SeatMap struct {
	occupied   map[int]map[int]bool
	maxX, maxY int
}

func NewSeatMap(input string) *SeatMap {
	occupied := make(map[int]map[int]bool)
	var maxX, maxY int
	for y, row := range strings.Split(input, "\n") {
		occupied[y] = make(map[int]bool, len(row))
		for x, seat := range row {
			if seat == 'L' {
				occupied[y][x] = false
			}
		}
		maxX = len(row)
		maxY++
	}
	return &SeatMap{occupied, maxX, maxY}
}

func (sm *SeatMap) RunSimulation(minOccupiedToChange int, adjacent bool) int {
	for changes := sm.changes(minOccupiedToChange, adjacent); len(changes) > 0; changes = sm.changes(minOccupiedToChange, adjacent) {
		for _, c := range changes {
			sm.occupied[c.y][c.x] = !sm.occupied[c.y][c.x]
		}
	}
	return sm.countOccupied()
}

type change struct{ x, y int }

func (sm *SeatMap) changes(minOccupiedToChange int, adjacent bool) []change {
	var changes []change
	for y := 0; y < sm.maxY; y++ {
		for x := 0; x < sm.maxX; x++ {
			if _, ok := sm.occupied[y][x]; !ok {
				continue
			}
			numNeighbors := sm.numNeighbors(x, y, adjacent)
			if numNeighbors == 0 && !sm.occupied[y][x] {
				changes = append(changes, change{x, y})
			} else if numNeighbors >= minOccupiedToChange && sm.occupied[y][x] {
				changes = append(changes, change{x, y})
			}
		}
	}
	return changes
}

func (sm *SeatMap) numNeighbors(x, y int, adjacent bool) int {
	var result int
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			for n := 1; ; n++ {
				outOfBounds := x+i*n < 0 || x+i*n >= sm.maxX || y+j*n < 0 || y+j*n >= sm.maxY
				if outOfBounds {
					break
				}
				occupied, isSeat := sm.occupied[y+j*n][x+i*n]
				if occupied {
					result++
				}
				if adjacent || isSeat {
					break
				}
			}
		}
	}
	return result
}

func (sm *SeatMap) countOccupied() int {
	var count int
	for _, row := range sm.occupied {
		for _, seat := range row {
			if seat {
				count++
			}
		}
	}
	return count
}
