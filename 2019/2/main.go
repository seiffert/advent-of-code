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
	c := intcode.NewComputer(p)
	c.Set(1, 12)
	c.Set(2, 2)

	if err := c.Calculate(); err != nil {
		lib.Abort("Error calculating result: %w", err)
	}
	fmt.Printf("Result for noun 12 and verb 2: %d\n", c.Get(0, intcode.ModeImmediate))

	noun, verb, err := IterMatrix(func(noun, verb int) error {
		c := intcode.NewComputer(p)
		c.Set(1, noun)
		c.Set(2, verb)

		if err := c.Calculate(); err != nil {
			return fmt.Errorf("%s: %w",
				err.Error(), ErrCantorTemporaryErr,
			)
		}

		result := c.Get(0, intcode.ModeImmediate)
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
