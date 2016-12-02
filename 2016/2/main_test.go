package main

import "fmt"

func ExampleOne() {
	input := []string{"ULL", "RRDDD", "LURDL", "UUUUD"}
	fmt.Printf("%s", CalculateCode(
		input,
		NewKeyPad([]Row{
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		}),
	))
	// Output: 1985
}
