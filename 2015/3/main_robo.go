package main

import (
	"fmt"
	"os"
)

func main() {
	var input string
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not read input: %s", err)
		os.Exit(1)
	}

	is, js, ir, jr := 0, 0, 0, 0
	var i, j *int
	count := 1
	round := 0
	grid := make(map[int]map[int]bool)
	grid[0] = map[int]bool{0: true}
	for _, c := range input {
		if round%2 == 0 {
			i, j = &is, &js
		} else {
			i, j = &ir, &jr
		}
		switch c {
		case 'v':
			*j--
		case '^':
			*j++
		case '>':
			*i++
		case '<':
			*i--
		}
		if _, ok := grid[*i]; !ok {
			grid[*i] = map[int]bool{}
		}
		if _, ok := grid[*i][*j]; !ok {
			grid[*i][*j] = true
			count++
		}
		round++
	}

	fmt.Fprintf(os.Stdout, "Result: %d\n", count)
}
