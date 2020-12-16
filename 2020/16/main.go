package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/seiffert/advent-of-code/2020/lib"
)

func main() {
	tl := NewTicketList(lib.MustReadFile("input.txt"))

	fmt.Printf("ticket scanning error rate: %d\n", tl.ErrorRate())

	t := tl.MyTicket()
	fmt.Printf("my ticket: %#v\n", t)

	p := 1
	for f, n := range t {
		if strings.HasPrefix(f, "departure") {
			p *= n
		}
	}
	fmt.Printf("my departure: %d\n", p)
}

type (
	TicketList struct {
		rules         []rule
		myTicket      ticket
		nearbyTickets []ticket
	}
	rule struct {
		Name       string
		MinA, MaxA int
		MinB, MaxB int
	}
	ticket []int
)

var (
	ruleRegexp = regexp.MustCompile(`^([a-z ]+): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)$`)
)

func NewTicketList(input string) *TicketList {
	c := strings.Split(input, "\n\n")

	var tl TicketList
	for _, rawRule := range strings.Split(c[0], "\n") {
		matches := ruleRegexp.FindStringSubmatch(rawRule)
		tl.rules = append(tl.rules, rule{
			Name: matches[1],
			MinA: lib.MustInt(matches[2]),
			MaxA: lib.MustInt(matches[3]),
			MinB: lib.MustInt(matches[4]),
			MaxB: lib.MustInt(matches[5]),
		})
	}

	tl.myTicket = newTicket(strings.Split(c[1], "\n")[1])
	for _, rawTicket := range strings.Split(c[2], "\n")[1:] {
		tl.nearbyTickets = append(tl.nearbyTickets, newTicket(rawTicket))
	}

	return &tl
}

func (tl *TicketList) ErrorRate() int {
	var result int
	for _, t := range tl.nearbyTickets {
	number:
		for _, n := range t {
			for _, r := range tl.rules {
				if r.Matches(n) {
					continue number
				}
			}
			result += n
		}
	}
	return result
}

func (tl *TicketList) MyTicket() map[string]int {
	rulesByPos := [][]rule{}

nextTicket:
	for _, t := range tl.nearbyTickets {
		if len(rulesByPos) == 0 {
			for _, n := range t {
				rules := []rule{}
				for _, r := range tl.rules {
					if r.Matches(n) {
						rules = append(rules, r)
					}
				}
				if len(rules) != 0 {
					rulesByPos = append(rulesByPos, rules)
				}
			}
		} else {
			newRulesByPos := [][]rule{}
			for i, n := range t {
				rules := []rule{}
				for _, r := range rulesByPos[i] {
					if r.Matches(n) {
						rules = append(rules, r)
					}
				}
				if len(rules) == 0 {
					continue nextTicket
				}
				newRulesByPos = append(newRulesByPos, rules)
			}
			rulesByPos = newRulesByPos
		}
	}

	mt := make(map[string]int, len(tl.myTicket))
	for len(mt) < len(rulesByPos) {
		for i := 0; i < len(rulesByPos); i++ {
			if len(rulesByPos[i]) == 1 {
				mt[rulesByPos[i][0].Name] = tl.myTicket[i]

				for j := 0; j < len(rulesByPos); j++ {
					if i == j {
						continue
					}
					n := []rule{}
					for _, r := range rulesByPos[j] {
						if r != rulesByPos[i][0] {
							n = append(n, r)
						}
					}
					rulesByPos[j] = n
				}
			}
		}
	}
	return mt
}

func newTicket(input string) ticket {
	var t ticket
	for _, n := range strings.Split(input, ",") {
		t = append(t, lib.MustInt(n))
	}
	return t
}

func (r rule) Matches(n int) bool {
	return r.MinA <= n && r.MaxA >= n || r.MinB <= n && r.MaxB >= n
}
