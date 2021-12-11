package main

import (
	"fmt"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	fmt.Printf("flahes after 100 rounds: %d\n", CountFlahes(100, lib.MustReadFile("input.txt")))
	fmt.Printf("first synced flash after %d rounds\n", FirstSyncedFlash(lib.MustReadFile("input.txt")))
}

type field map[coord]int
type coord struct{ x, y int }

func CountFlahes(steps int, in string) (out int) {
	f := parse(in)

	for i := 0; i < steps; i++ {
		out += f.step()
	}
	return
}

func FirstSyncedFlash(in string) int {
	f := parse(in)

	for i := 1; true; i++ {
		if 100 == f.step() {
			return i
		}
	}
	return -1
}

func parse(in string) field {
	f := field{}
	for y, l := range strings.Split(in, "\n") {
		for x, o := range l {
			f[coord{x, y}] = lib.MustInt(string(o))
		}
	}
	return f
}

func (f field) step() (flahes int) {
	var todo, flashed []coord
	for c := range f {
		f[c]++
		if f[c] == 10 {
			todo = append(todo, c)
		}
	}

	for len(todo) > 0 {
		c := todo[0]
		todo = todo[1:]

		flahes++
		flashed = append(flashed, c)
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dx == 0 && dy == 0 {
					continue
				}

				next := coord{c.x + dx, c.y + dy}
				if _, ok := f[next]; ok {
					f[next]++
					if f[next] == 10 {
						todo = append(todo, next)
					}
				}
			}
		}
	}

	for _, c := range flashed {
		f[c] = 0
	}

	return
}
