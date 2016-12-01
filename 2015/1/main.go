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

	var floor int
	for _, c := range input {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}
	}

	fmt.Fprintf(os.Stdout, "Result: %d\n", floor)
}
