package main

import (
	"fmt"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	fmt.Printf("fish after 80 days: %d\n", CountFishAfterDays(80, lib.MustReadFile("input.txt")))
	fmt.Printf("fish after 256 days: %d\n", CountFishAfterDays(256, lib.MustReadFile("input.txt")))
}

type memory map[dayFish]int
type dayFish struct{ day, offset int }

func CountFishAfterDays(days int, in string) (out int) {
	mem := memory{}
	for _, f := range lib.MustAllInts(strings.Split(in, ",")) {
		out += countAfterDays(days, f, mem)
	}
	return
}

func countAfterDays(days, offset int, mem memory) (out int) {
	defer func() { mem[dayFish{days, offset}] = out }()

	if days == 0 {
		return 1
	}

	if res, ok := mem[dayFish{days, offset}]; ok {
		out = res
		return
	}

	if offset == 0 {
		out = countAfterDays(days-1, 8, mem) + countAfterDays(days-1, 6, mem)
		return
	}
	out = countAfterDays(days-1, offset-1, mem)
	return
}
