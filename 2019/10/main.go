package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/seiffert/advent-of-code/2019/lib"
)

func main() {
	var input string
	if len(os.Args) == 2 {
		input = os.Args[1]
	} else {
		fileInput, err := ioutil.ReadFile("input.txt")
		if err != nil {
			lib.Abort("error reading input file: %w", err)
		}
		input = string(fileInput)
	}

	field := NewAsteroidField(input)
	best := field.Find(func(a, b *Asteroid) bool {
		return a.AsteroidsInSight(field) > b.AsteroidsInSight(field)
	})

	fmt.Printf("Best asteroid: %s, with %d others in sight\n",
		best, best.AsteroidsInSight(field),
	)
}
