package main

import (
	"io/ioutil"
	"os"

	"github.com/seiffert/advent-of-code/2019/lib"
	"github.com/seiffert/advent-of-code/2019/lib/intcode"
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

	p := intcode.Parse(input)
	computer := intcode.NewComputer(p)

	if err := computer.Calculate(); err != nil {
		lib.Abort("error calculating result: %w", err)
	}
}
