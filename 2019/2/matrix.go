package main

import (
	"errors"
)

var ErrCantorTemporaryErr = errors.New("temporary error")

func IterMatrix(f func(x, y int) error) (int, int, error) {
	i := 0
	for {
		for x := i; x >= 0; x-- {
			y := i - x

			if err := f(x, y); err != nil {
				if !errors.Is(err, ErrCantorTemporaryErr) {
					return x, y, err
				}
			} else {
				return x, y, nil
			}
		}
		i += 1
	}
}
