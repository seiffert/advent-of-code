package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/2020/lib"
)

func main() {
	g := NewGame(lib.MustReadFile("input.txt"))
	score := g.PlayA()

	fmt.Println("winner's score regular game:", score)

	g = NewGame(lib.MustReadFile("input.txt"))
	score = g.PlayB()

	fmt.Println("winner's score recursive game:", score)
}

type (
	Game struct {
		playerA, playerB stack
		roundsA, roundsB map[string]bool
	}
	stack []int
)

func NewGame(input string) *Game {
	players := strings.Split(input, "\n\n")
	return &Game{
		playerA: parseStack(players[0]),
		playerB: parseStack(players[1]),
		roundsA: make(map[string]bool),
		roundsB: make(map[string]bool),
	}
}

func parseStack(input string) stack {
	var s stack
	for i, line := range strings.Split(input, "\n") {
		if i == 0 {
			continue
		}
		n, _ := strconv.Atoi(line)
		s = append(s, n)
	}
	return s
}

func (g *Game) PlayA() int {
	for len(g.playerA) != 0 && len(g.playerB) != 0 {
		if g.playerA[0] > g.playerB[0] {
			g.winRoundA()
		} else {
			g.winRoundB()
		}
	}

	if len(g.playerA) == 0 {
		return calcScore(g.playerB)
	}
	return calcScore(g.playerA)
}

func (g *Game) winRoundA() {
	g.playerA = append(g.playerA[1:], g.playerA[0], g.playerB[0])
	g.playerB = g.playerB[1:]
}

func (g *Game) winRoundB() {
	g.playerB = append(g.playerB[1:], g.playerB[0], g.playerA[0])
	g.playerA = g.playerA[1:]
}

func (g *Game) PlayB() int {
	if 'a' == g.playB() {
		return calcScore(g.playerA)
	}
	return calcScore(g.playerB)
}

func (g *Game) playB() byte {
	for len(g.playerA) != 0 && len(g.playerB) != 0 {
		keyA, keyB := fmt.Sprint(g.playerA), fmt.Sprint(g.playerB)
		if g.roundsA[keyA] || g.roundsB[keyB] {
			return 'a'
		}
		g.roundsA[keyA] = true
		g.roundsB[keyB] = true

		a, b := g.playerA[0], g.playerB[0]
		if a >= len(g.playerA) || b >= len(g.playerB) {
			if a > b {
				g.winRoundA()
				continue
			}
			g.winRoundB()
			continue
		}

		sub := &Game{
			roundsA: make(map[string]bool),
			roundsB: make(map[string]bool),
			playerA: make(stack, a),
			playerB: make(stack, b),
		}
		copy(sub.playerA, g.playerA[1:a+1])
		copy(sub.playerB, g.playerB[1:b+1])
		if 'a' == sub.playB() {
			g.winRoundA()
			continue
		}
		g.winRoundB()
	}
	if len(g.playerA) == 0 {
		return 'b'
	}
	return 'a'
}

func calcScore(s stack) int {
	var score int
	for i := 1; i <= len(s); i++ {
		score += i * s[len(s)-i]
	}
	return score
}
