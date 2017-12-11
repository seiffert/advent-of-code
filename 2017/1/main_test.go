package main

import "testing"

func TestSolveA(t *testing.T) {
	tc := map[string]int{
		"1122":     3,
		"1111":     4,
		"1234":     0,
		"91212129": 9,
	}

	for c, sum := range tc {
		if res := SolveA(c); res != sum {
			t.Errorf("SolveA(%q) == %d != (expected) %d", c, res, sum)
		}
	}
}

func TestSolveB(t *testing.T) {
	tc := map[string]int{
		"1212":     6,
		"1221":     0,
		"123123":   12,
		"12131415": 4,
	}

	for c, sum := range tc {
		if res := SolveB(c); res != sum {
			t.Errorf("SolveB(%q) == %d != (expected) %d", c, res, sum)
		}
	}
}
