package main

import (
	"fmt"
	"os"

	"github.com/seiffert/advent-of-code/2019/lib"
	"github.com/seiffert/advent-of-code/2019/lib/intcode"
)

func main() {
	if len(os.Args) != 2 {
		lib.Abort("expects the input as argument, got none")
	}

	p := intcode.Parse(os.Args[1])
	c := intcode.NewComputer(p, 12, 2)

	result, err := c.Calculate()
	if err != nil {
		lib.Abort("Error calculating result: %w", err)
	}

	fmt.Printf("Result for noun 12 and verb 2: %d\n", result)

	noun, verb, err := IterMatrix(func(noun, verb int) error {
		c := intcode.NewComputer(p, noun, verb)

		result, err := c.Calculate()
		if err != nil {
			return fmt.Errorf("%s: %w",
				err.Error(), ErrCantorTemporaryErr,
			)
		}

		if result != 19690720 {
			return fmt.Errorf("not the expected result: %w",
				ErrCantorTemporaryErr,
			)
		}
		return nil
	})
	if err != nil {
		lib.Abort(err.Error())
	}

	fmt.Printf("To produce result '19690720', use noun %d and verb %d.\n",
		noun, verb,
	)
}
