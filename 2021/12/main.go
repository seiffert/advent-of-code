package main

import (
	"fmt"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	fmt.Printf("number of paths part 1: %d\n", CountPaths(lib.MustReadFile("input.txt"), false))
	fmt.Printf("number of paths part 2: %d\n", CountPaths(lib.MustReadFile("input.txt"), true))
}

type Path []string

func CountPaths(in string, smallCaveTwice bool) int {
	conns := make(map[string][]string)
	for _, l := range strings.Split(in, "\n") {
		conn := strings.Split(l, "-")
		conns[conn[0]] = append(conns[conn[0]], conn[1])
		conns[conn[1]] = append(conns[conn[1]], conn[0])
	}

	unfinished, finished := []Path{{"start"}}, []Path{}
	for len(unfinished) > 0 {
		p := unfinished[0]
		unfinished = unfinished[1:]

		last := p[len(p)-1]
		for _, c := range conns[last] {
			switch {
			case c == "start":
				continue
			case c == "end":
				finished = append(finished, append(append(Path{}, p...), c))
			case c[0] >= 'a' && c[0] <= 'z' && (!p.contains(c) || smallCaveTwice && !p.containsSmallCaveTwice()):
				fallthrough
			case c[0] >= 'A' && c[0] <= 'Z':
				unfinished = append(unfinished, append(append(Path{}, p...), c))
			}
		}
	}

	return len(finished)
}

func (p Path) contains(c string) bool {
	for _, s := range p {
		if s == c {
			return true
		}
	}
	return false
}

func (p Path) containsSmallCaveTwice() bool {
	visited := map[string]bool{}
	for _, s := range p {
		if visited[s] && s[0] >= 'a' && s[0] <= 'z' {
			return true
		}
		visited[s] = true
	}
	return false
}
