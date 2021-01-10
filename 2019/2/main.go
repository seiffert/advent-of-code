package main

import (
	"fmt"

	"github.com/seiffert/advent-of-code/lib"
	"github.com/seiffert/advent-of-code/lib/intcode"
)

func main() {
	c := intcode.New(lib.MustReadFile("input.txt"))
	c.Set(1, 12)
	c.Set(2, 2)
	fmt.Println("Pos 0:", c.Run().Get(0))

	for n := 0; n <= 99; n++ {
		for v := 0; v <= n; v++ {
			c = intcode.New(lib.MustReadFile("input.txt"))
			c.Set(1, n)
			c.Set(2, v)
			if c.Run().Get(0) == 19690720 {
				fmt.Println("Noun", n, "verb", v)
				return
			}
		}
	}
}
