package main

import "fmt"

func ExampleOne() {
	fmt.Printf("%d", CalculateDistance("R2, L3"))
	// Output:
	// 5
}

func ExampleTwo() {
	fmt.Printf("%d", CalculateDistance("R2, R2, R2"))
	// Output:
	// 2
}

func ExampleThree() {
	fmt.Printf("%d", CalculateDistance("R5, L5, R5, R3"))
	// Output:
	// 12
}

func ExampleFour() {
	fmt.Printf("%d", CalculateDistance("R8, R4, R4, R8"))
	// Output:
	// 4
}