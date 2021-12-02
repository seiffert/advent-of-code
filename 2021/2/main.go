package main

import (
	"fmt"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	in := lib.MustReadFile("input.txt")

	p := Navigate(in)
	fmt.Printf("horizontal: %d, depth: %d, dist: %d\n", p.H, p.D, p.H*p.D)

	p = NavigateCorrectly(in)
	fmt.Printf("correct: horizontal: %d, depth: %d, dist: %d\n", p.H, p.D, p.H*p.D)
}

func Navigate(in string) (out Pos) {
	for _, i := range strings.Split(in, "\n") {
		p := strings.Split(i, " ")
		switch p[0] {
		case "forward":
			out.H += lib.MustInt(p[1])
		case "up":
			out.D -= lib.MustInt(p[1])
		case "down":
			out.D += lib.MustInt(p[1])
		}
	}
	return
}

func NavigateCorrectly(in string) (out Pos) {
	for _, i := range strings.Split(in, "\n") {
		p := strings.Split(i, " ")
		switch p[0] {
		case "forward":
			x := lib.MustInt(p[1])
			out.H += x
			out.D += x * out.A
		case "up":
			out.A -= lib.MustInt(p[1])
		case "down":
			out.A += lib.MustInt(p[1])
		}
	}
	return
}

type Pos struct{ H, D, A int }
