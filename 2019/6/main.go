package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/seiffert/advent-of-code/2019/lib"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		lib.Abort("error reading input file: %w", err)
	}

	m := NewMap()
	for _, orbit := range strings.Split(string(input), "\n") {
		if len(strings.TrimSpace(orbit)) == 0 {
			continue
		}
		m.Add(orbit)
	}

	fmt.Println("number of orbits", m.CountOrbits())

	fmt.Println("number of orbits between you and santa", m.Distance("YOU", "SAN"))
}
