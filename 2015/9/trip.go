package main

import (
	"fmt"
	"strings"
)

type Trip struct {
	Cities        []*City
	visitedCities map[*City]bool
}

func (t *Trip) AddCity(c *City) {
	if t.visitedCities == nil {
		t.reset()
	}

	if _, ok := t.visitedCities[c]; !ok {
		t.Cities = append(t.Cities, c)
		t.visitedCities[c] = true
	}
}

func (t *Trip) VisitedCity(c *City) bool {
	if t.visitedCities == nil {
		t.reset()
	}

	visited, ok := t.visitedCities[c]

	return ok && visited
}

func (t *Trip) Distance() int {
	var prev *City
	var distance int
	for _, c := range t.Cities {
		if prev != nil {
			distance += dist(prev, c)
		}
		prev = c
	}
	return distance
}

func (t *Trip) IsShorterThan(other *Trip) bool {
	return t.Distance() < other.Distance()
}

func (t *Trip) Copy() *Trip {
	new := &Trip{}
	new.CopyFrom(t)
	return new
}

func (t *Trip) CopyFrom(other *Trip) {
	t.reset()

	for _, c := range other.Cities {
		t.AddCity(c)
	}
}

func (t *Trip) reset() {
	t.Cities = []*City{}
	t.visitedCities = make(map[*City]bool)
}

func (t *Trip) String() string {
	var names []string
	for _, c := range t.Cities {
		names = append(names, c.Name)
	}

	return fmt.Sprintf("%s = %d", strings.Join(names, " -> "), t.Distance())
}
