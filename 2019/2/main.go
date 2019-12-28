package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		abort("expects the input as argument, got none")
	}

	p := Parse(os.Args[1])
	c := NewComputer(p, 12, 2)

	result, err := c.Calculate()
	if err != nil {
		abort(err.Error())
	}

	fmt.Printf("Result for noun 12 and verb 2: %d\n", result)

	noun, verb, err := IterMatrix(func(noun, verb int) error {
		c := NewComputer(p, noun, verb)

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
		abort(err.Error())
	}

	fmt.Printf("To produce result '19690720', use noun %d and verb %d.\n",
		noun, verb,
	)
}

func abort(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
