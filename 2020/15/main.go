package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/2020/lib"
)

func main() {
	g := NewGame(lib.MustReadFile("input.txt"))

	fmt.Printf("the 2020th number is %d\n", g.Play(2020))
	fmt.Printf("the 30000000th number is %d\n", g.Play(30000000))
}

type Game struct {
	turn, lastNum int
	lastSpoken    map[int]twoLast
}
type twoLast struct{ first, second int }

func NewGame(input string) *Game {
	g := &Game{
		lastSpoken: make(map[int]twoLast),
	}
	for _, s := range strings.Split(input, ",") {
		n, _ := strconv.Atoi(s)
		g.speak(n)
	}
	return g
}

func (g *Game) Play(turns int) int {
	for g.turn < turns {
		tl := g.lastSpoken[g.lastNum]
		if tl.first == 0 {
			g.speak(0)
		} else {
			g.speak(tl.second - tl.first)
		}
	}
	return g.lastNum
}

func (g *Game) speak(n int) {
	g.turn++
	g.lastSpoken[n] = twoLast{
		first:  g.lastSpoken[n].second,
		second: g.turn,
	}
	g.lastNum = n
}
