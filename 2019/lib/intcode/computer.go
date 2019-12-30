package intcode

import (
	"fmt"

	"github.com/seiffert/advent-of-code/2019/lib"
)

const (
	ModePosition  = 0
	ModeImmediate = 1

	opcodeAdd         = 1
	opcodeMultiply    = 2
	opcodeInput       = 3
	opcodeOutput      = 4
	opcodeJumpIfTrue  = 5
	opcodeJumpIfFalse = 6
	opcodeLessThan    = 7
	opcodeEquals      = 8
	opcodeTerminate   = 99
)

func NewComputer(p []int) *Computer {
	mem := append([]int(nil), p...)
	return &Computer{
		memory: mem,
	}
}

type Computer struct {
	memory []int
	ic     int
}

func (c *Computer) Calculate() error {
	for {
		var (
			instruction = c.memory[c.ic]
			opcode      = instruction % 100
			param1Mode  = instruction % 1000 / 100
			param2Mode  = instruction % 10000 / 1000
			//param3Mode  = instruction % 100000 / 10000
		)

		switch opcode {
		case opcodeAdd:
			c.Set(
				c.Get(c.ic+3, ModeImmediate),
				c.Get(c.ic+1, param1Mode)+c.Get(c.ic+2, param2Mode),
			)
			c.ic += 4
		case opcodeMultiply:
			c.Set(
				c.Get(c.ic+3, ModeImmediate),
				c.Get(c.ic+1, param1Mode)*c.Get(c.ic+2, param2Mode),
			)
			c.ic += 4
		case opcodeInput:
			var (
				valid bool
				input int
			)

			for !valid {
				fmt.Println("Input value (integer):")
				if _, err := fmt.Scanf("%d", &input); err != nil {
					lib.LogError("invalid input: %w", err)
				} else {
					valid = true
				}
			}

			c.Set(c.Get(c.ic+1, ModeImmediate), input)
			c.ic += 2
		case opcodeOutput:
			fmt.Printf("Output: %d\n", c.Get(c.ic+1, param1Mode))
			c.ic += 2
		case opcodeJumpIfTrue:
			if c.Get(c.ic+1, param1Mode) != 0 {
				c.ic = c.Get(c.ic+2, param2Mode)
			} else {
				c.ic += 3
			}
		case opcodeJumpIfFalse:
			if c.Get(c.ic+1, param1Mode) == 0 {
				c.ic = c.Get(c.ic+2, param2Mode)
			} else {
				c.ic += 3
			}
		case opcodeLessThan:
			if c.Get(c.ic+1, param1Mode) < c.Get(c.ic+2, param2Mode) {
				c.Set(c.Get(c.ic+3, ModeImmediate), 1)
			} else {
				c.Set(c.Get(c.ic+3, ModeImmediate), 0)
			}
			c.ic += 4
		case opcodeEquals:
			if c.Get(c.ic+1, param1Mode) == c.Get(c.ic+2, param2Mode) {
				c.Set(c.Get(c.ic+3, ModeImmediate), 1)
			} else {
				c.Set(c.Get(c.ic+3, ModeImmediate), 0)
			}
			c.ic += 4
		case opcodeTerminate:
			return nil
		default:
			return fmt.Errorf("unknown opcode %d", opcode)
		}
	}
}

func (c *Computer) Get(addr, mode int) int {
	if mode == ModePosition {
		addr = c.Get(addr, ModeImmediate)
	}

	if addr < len(c.memory) {
		return c.memory[addr]
	}
	return 0
}

func (c *Computer) Set(addr, val int) {
	if addr >= len(c.memory) {
		c.growMemory(addr - len(c.memory) + 1)
	}

	c.memory[addr] = val
}

func (c *Computer) growMemory(diff int) {
	c.memory = append(c.memory, make([]int, diff)...)
}
