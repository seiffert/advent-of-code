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

	i, j := 0, 0
	count := 1
	grid := make(map[int]map[int]bool)
	grid[i] = map[int]bool{j: true}
	for _, c := range input {
		switch c {
		case 'v':
			j--
		case '^':
			j++
		case '>':
			i++
		case '<':
			i--
		}
		if _, ok := grid[i]; !ok {
			grid[i] = map[int]bool{}
		}
		if _, ok := grid[i][j]; !ok {
			grid[i][j] = true
			count++
		}
	}

	fmt.Fprintf(os.Stdout, "Result: %d\n", count)
}
