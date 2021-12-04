package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	g := NewGame(lib.MustReadFile("input.txt"))

	g.Play()

	fmt.Printf("final score of winning board: %d\n", g.WinnerScore())
	fmt.Printf("final score of last-winning board: %d\n", g.LastWinnerScore())
}

type (
	Game struct {
		DrawnNumbers []int
		Boards       []*Board
		Winners      []Winner
	}
	Winner struct {
		WinningNumber int
		Board         *Board
	}
	Board struct {
		Fields map[Coords]*Field
	}
	Coords struct {
		X, Y int
	}
	Field struct {
		Number int
		Drawn  bool
	}
)

func NewGame(in string) *Game {
	// Replace all duplicate whitespaces
	in = regexp.MustCompile(` +`).ReplaceAllString(in, " ")

	splitIn := strings.Split(in, "\n\n")

	boards := make([]*Board, 0, len(splitIn)-1)
	for _, boardIn := range splitIn[1:] {
		b := Board{
			Fields: make(map[Coords]*Field),
		}
		for y, r := range strings.Split(boardIn, "\n") {
			for x, n := range lib.MustAllInts(strings.Split(strings.TrimSpace(r), " ")) {
				b.Fields[Coords{x, y}] = &Field{Number: n}
			}
		}
		boards = append(boards, &b)
	}

	return &Game{
		DrawnNumbers: lib.MustAllInts(strings.Split(splitIn[0], ",")),
		Boards:       boards,
	}
}

func (g *Game) Play() {
	for _, n := range g.DrawnNumbers {
		for _, b := range g.Boards {
			if g.HasWon(b) {
				continue
			}
			for c, f := range b.Fields {
				if f.Number == n {
					f.Drawn = true

					if b.IsWinningAfterDrawing(c) {
						g.Winners = append(g.Winners, Winner{WinningNumber: n, Board: b})
					}
				}
			}
		}
	}
	fmt.Printf("%d winners\n", len(g.Winners))
}

func (g *Game) WinnerScore() int {
	return g.winnerScore(0)
}

func (g *Game) LastWinnerScore() int {
	return g.winnerScore(len(g.Winners) - 1)
}

func (g *Game) winnerScore(i int) int {
	if len(g.Winners) < i+1 {
		lib.Abort("must play first")
	}
	w := g.Winners[i]

	var sumUnplayed int
	for _, f := range w.Board.Fields {
		if !f.Drawn {
			sumUnplayed += f.Number
		}
	}

	return sumUnplayed * w.WinningNumber
}

func (g *Game) HasWon(b *Board) bool {
	for _, w := range g.Winners {
		if w.Board == b {
			return true
		}
	}

	return false
}

func (b Board) IsWinningAfterDrawing(c Coords) bool {
	hasRow, hasCol := true, true
	for x := 0; x < 5; x++ {
		if !b.Fields[Coords{x, c.Y}].Drawn {
			hasRow = false
			break
		}
	}
	if hasRow {
		return true
	}

	for y := 0; y < 5; y++ {
		if !b.Fields[Coords{c.X, y}].Drawn {
			hasCol = false
			break
		}
	}

	return hasCol
}
