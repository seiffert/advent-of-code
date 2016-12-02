package main

import (
	"bufio"
	"fmt"
	"os"
)

func NewKeyPad(rows []Row) *KeyPad {
	var x5, y5 int
loops:
	for y := 0; y < len(rows); y++ {
		for x := 0; x < len(rows[y]); x++ {
			if rows[y][x] == "5" {
				x5 = x
				y5 = y
				break loops
			}
		}
	}
	return &KeyPad{
		Rows: rows,
		Position: Coords{
			X: x5,
			Y: y5,
		},
	}
}

type Coords struct {
	X, Y int
}

type KeyPad struct {
	Rows     []Row
	Position Coords
}
type Row []string

func (k *KeyPad) Move(step rune) {
	switch step {
	case 'U':
		if k.Position.Y-1 >= 0 && k.Rows[k.Position.Y-1][k.Position.X] != "O" {
			k.Position.Y--
		}
	case 'D':
		if k.Position.Y+1 < len(k.Rows) && k.Rows[k.Position.Y+1][k.Position.X] != "O" {
			k.Position.Y++
		}
	case 'L':
		if k.Position.X-1 >= 0 && k.Rows[k.Position.Y][k.Position.X-1] != "O" {
			k.Position.X--
		}
	case 'R':
		if k.Position.X+1 < len(k.Rows[k.Position.Y]) && k.Rows[k.Position.Y][k.Position.X+1] != "O" {
			k.Position.X++
		}
	}
}
func (k *KeyPad) CurrentPosition() string {
	return k.Rows[k.Position.Y][k.Position.X]
}

func CalculateCode(lines []string, k *KeyPad) string {
	result := ""

	for _, line := range lines {
		for _, step := range line {
			k.Move(step)
		}
		result += k.CurrentPosition()
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, string(scanner.Text()))
	}
	fmt.Println(CalculateCode(lines, NewKeyPad([]Row{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	})))
	fmt.Println(CalculateCode(lines, NewKeyPad([]Row{
		{"O", "O", "1", "O", "O"},
		{"O", "2", "3", "4", "O"},
		{"5", "6", "7", "8", "9"},
		{"O", "A", "B", "C", "O"},
		{"O", "O", "D", "O", "O"},
	})))
}
