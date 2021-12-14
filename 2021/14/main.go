package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	p := NewPolymer(lib.MustReadFile("input.txt"))
	p.Expand(10)

	c := p.ElementCounts()
	fmt.Printf("%d - %d = %d\n", c[len(c)-1], c[0], c[len(c)-1]-c[0])

	p.Expand(30)
	c = p.ElementCounts()
	fmt.Printf("%d - %d = %d\n", c[len(c)-1], c[0], c[len(c)-1]-c[0])
}

type Polymer struct {
	Pairs map[string]int
	Chars map[string]int
	Rules map[string]string
}

func NewPolymer(in string) *Polymer {
	var (
		lines = strings.Split(in, "\n")
		rules = map[string]string{}
		pairs = map[string]int{}
		chars = map[string]int{}
	)

	for _, l := range lines[2:] {
		r := strings.Split(l, " -> ")
		rules[r[0]] = r[1]
	}

	for i := 0; i < len(lines[0])-1; i++ {
		pairs[lines[0][i:i+2]]++
		chars[lines[0][i:i+1]]++
	}
	chars[lines[0][len(lines[0])-1:]]++

	return &Polymer{
		Pairs: pairs,
		Chars: chars,
		Rules: rules,
	}
}

func (p *Polymer) Expand(n int) {
	for ; n > 0; n-- {
		np := map[string]int{}
		for pair, c := range p.Pairs {
			if sub, ok := p.Rules[pair]; ok {
				np[pair[:1]+sub] += c
				np[sub+pair[1:]] += c
				p.Chars[sub] += c
			}
		}
		p.Pairs = np
	}
}

func (p Polymer) ElementCounts() []int {
	nums := make([]int, 0, len(p.Chars))
	for _, c := range p.Chars {
		nums = append(nums, c)
	}

	sort.Ints(nums)
	return nums
}
