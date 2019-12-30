package main

import (
	"strings"

	"github.com/seiffert/advent-of-code/2019/lib"
)

func NewMap() *Map {
	return &Map{
		orbits: make(map[string]string),
	}
}

type Map struct {
	orbits map[string]string
}

func (m *Map) Add(orbit string) {
	comps := strings.Split(orbit, ")")
	if len(comps) != 2 {
		lib.Abort("invalid orbit, instruction %q had %d components", orbit, len(comps))
	}

	m.orbits[comps[1]] = comps[0]
}

func (m *Map) CountOrbits() int {
	var sum int
	for inOrbit, _ := range m.orbits {
		d, ok := m.distanceToParent(inOrbit, "COM")
		if !ok {
			return 0
		}
		sum += d
	}
	return sum
}

func (m *Map) Distance(from, to string) int {
	a, b := m.orbits[from], m.orbits[to]

	var distance int
	for {
		if d, ok := m.distanceToParent(b, a); ok {
			return distance + d
		}
		distance++
		a = m.orbits[a]
	}
}

func (m *Map) distanceToParent(from, parent string) (int, bool) {
	var distance int
	for from != parent {
		distance++
		p, ok := m.orbits[from]
		if !ok {
			return 0, false
		}
		from = p
	}
	return distance, true
}
