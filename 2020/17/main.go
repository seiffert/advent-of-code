package main

import (
	"fmt"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	g3 := New3DGrid(lib.MustReadFile("input.txt"))
	g3.Play(6)

	fmt.Printf("3d after six cycles: %d\n", g3.CountActive())

	g4 := New4DGrid(lib.MustReadFile("input.txt"))
	g4.Play(6)

	fmt.Printf("4d after six cycles: %d\n", g4.CountActive())
}
