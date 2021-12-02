package main

import (
	"fmt"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	in := lib.MustReadFile("input.txt")

	p := Navigate(in)
	fmt.Printf("horizontal: %d, depth: %d, dist: %d\n", p.Horizontal, p.Depth, p.Horizontal*p.Depth)

	p = NavigateCorrectly(in)
	fmt.Printf("correct: horizontal: %d, depth: %d, dist: %d\n", p.Horizontal, p.Depth, p.Horizontal*p.Depth)
}

func Navigate(in string) Position {
	sub := Submarine{
		Instructions: map[string]func(*Submarine, int){
			"forward": func(s *Submarine, arg int) { s.Horizontal += arg },
			"up":      func(s *Submarine, arg int) { s.Depth -= arg },
			"down":    func(s *Submarine, arg int) { s.Depth += arg },
		},
	}
	sub.Navigate(in)
	return sub.Position
}

func NavigateCorrectly(in string) Position {
	sub := Submarine{
		Instructions: map[string]func(*Submarine, int){
			"forward": func(s *Submarine, arg int) { s.Horizontal, s.Depth = s.Horizontal+arg, s.Depth+s.Aim*arg },
			"up":      func(s *Submarine, arg int) { s.Aim -= arg },
			"down":    func(s *Submarine, arg int) { s.Aim += arg },
		},
	}
	sub.Navigate(in)
	return sub.Position
}

type (
	Submarine struct {
		Instructions map[string]func(*Submarine, int)
		Position
	}
	Position struct {
		Horizontal, Depth, Aim int
	}
)

func (s *Submarine) Navigate(in string) {
	for _, i := range strings.Split(in, "\n") {
		p := strings.Split(i, " ")
		s.Instructions[p[0]](s, lib.MustInt(p[1]))
	}
}
