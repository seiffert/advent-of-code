package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	RuleRegexp = regexp.MustCompile("^([A-Za-z]+) => ([A-Za-z]+)$")
	cache      = map[string][]string{}
)

type Rule struct {
	In  string
	Out string
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	rules := []Rule{}
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}

		rules = append(rules, parseRule(line))
	}
	s.Scan()
	molecule := s.Text()

	if len(os.Args) > 1 && os.Args[1] == "2" {
		reverseRules := []Rule{}
		for _, r := range rules {
			reverseRules = append(reverseRules, Rule{
				In:  r.Out,
				Out: r.In,
			})
		}

		in := []string{molecule}
		i := 1
	OuterLoop:
		for {
			log.Printf("Round %d with %d elements", i, len(in))
			new := []string{}
			for _, i := range in {
				for _, t := range allTransformations(i, reverseRules) {
					if t == "e" {
						break OuterLoop
					}
					new = append(new, t)
				}
			}
			in = new
			i++
		}

		log.Printf("%d", i)
	} else {
		molecules := map[string]bool{}
		for _, rule := range rules {
			for _, v := range transformations(molecule, rule) {
				molecules[v] = true
			}
		}
		log.Printf("%d", len(molecules))
	}
}

func allTransformations(molecule string, rules []Rule) []string {
	if t, ok := cache[molecule]; ok {
		return t
	}

	t := []string{}
	for _, r := range rules {
		t = append(t, transformations(molecule, r)...)
	}
	cache[molecule] = t
	return t
}

func transformations(molecule string, r Rule) []string {
	offset := 0
	molecules := []string{}
	for index := strings.Index(molecule[offset:], r.In); index != -1; index = strings.Index(molecule[offset:], r.In) {
		molecule := molecule[:offset+index] + r.Out + molecule[offset+index+len(r.In):]

		molecules = append(molecules, molecule)
		offset += index + len(r.In)

		if offset >= len(molecule) {
			break
		}
	}
	return molecules
}

func parseRule(s string) Rule {
	m := RuleRegexp.FindAllStringSubmatch(s, -1)
	return Rule{
		In:  m[0][1],
		Out: m[0][2],
	}
}

func countSteps(in, out string, rules []Rule, step, limit int) int {
	shortest := -1
	for _, r := range rules {
		for _, t := range transformations(in, r) {
			if t == out {
				return step + 1
			}

			s := countSteps(t, out, rules, step+1, limit)
			if shortest == -1 || shortest > s {
				shortest = s
			}
			if shortest < limit {
				log.Printf("New limit: %d", shortest)
				limit = shortest
			}
		}
	}
	return shortest
}
