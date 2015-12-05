package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var sum int
	var ribbon int
	for scanner.Scan() {
		sides := strings.Split(scanner.Text(), "x")
		intSlice, err := parseIntSlice(sides)
		if err != nil {
			abort(err)
		}
		if len(intSlice) != 3 {
			abort(fmt.Errorf("invalid input: %q", sides))
		}
		sort.Ints(intSlice)
		sum += calcWrappingPaper(intSlice)
		ribbon += calcRibbon(intSlice)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Printf("Paper: %d\n", sum)
	fmt.Printf("Ribbon: %d\n", ribbon)
}

func calcWrappingPaper(sides []int) int {
	return sides[0]*sides[1]*3 + sides[0]*sides[2]*2 + sides[1]*sides[2]*2
}

func calcRibbon(sides []int) int {
	return sides[0]*2 + sides[1]*2 + sides[0]*sides[1]*sides[2]
}

func parseIntSlice(slice []string) ([]int, error) {
	var result []int
	for _, s := range slice {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}

func abort(err error) {
	fmt.Fprintln(os.Stderr, "Error: %s", err)
	os.Exit(1)
}
