package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

var (
	bagRegexp = regexp.MustCompile(`^([0-9]+) ([a-z ]+) bags?$`)
)

func main() {
	rules := NewRules(lib.MustReadFile("input.txt"))

	fmt.Printf("%d different bags can contain a 'shiny gold' bag\n",
		rules.CountPossibleContainers("shiny gold"))
	fmt.Printf("%d bags required in a 'shiny gold' bag\n",
		rules.CountRequiredContents("shiny gold"))
}

type Rules struct {
	PossibleContainers map[string][]string
	RequiredContents   map[string]map[string]int
}

func NewRules(input string) Rules {
	r := Rules{
		PossibleContainers: make(map[string][]string),
		RequiredContents:   make(map[string]map[string]int),
	}
	for _, line := range strings.Split(input, "\n") {
		var (
			p               = strings.Split(strings.TrimSuffix(line, "."), " bags contain ")
			color, contains = p[0], p[1]
		)
		if contains == "no other bags" {
			continue
		}
		others := strings.Split(contains, ", ")
		for _, other := range others {
			matches := bagRegexp.FindStringSubmatch(other)
			otherCount, otherColor := matches[1], matches[2]
			if r.PossibleContainers[otherColor] == nil {
				r.PossibleContainers[otherColor] = []string{color}
			} else {
				r.PossibleContainers[otherColor] = append(r.PossibleContainers[otherColor], color)
			}
			if r.RequiredContents[color] == nil {
				r.RequiredContents[color] = make(map[string]int)
			}
			r.RequiredContents[color][otherColor], _ = strconv.Atoi(otherCount)
		}
	}
	return r
}

func (r Rules) CountPossibleContainers(c string) int {
	var (
		s      = []string{c}
		unique = map[string]struct{}{}
	)
	for len(s) > 0 {
		c, s = s[0], s[1:]
		if _, ok := unique[c]; !ok {
			unique[c] = struct{}{}
			s = append(s, r.PossibleContainers[c]...)
		}
	}
	return len(unique) - 1
}

func (r Rules) CountRequiredContents(c string) int {
	var sum int
	for color, num := range r.RequiredContents[c] {
		sum += num + num*r.CountRequiredContents(color)
	}
	return sum
}
