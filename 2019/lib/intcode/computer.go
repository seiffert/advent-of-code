package intcode

import (
	"fmt"

	"github.com/seiffert/advent-of-code/2019/lib"
)

const (
	opcodeAdd       = 1
	opcodeMultiply  = 2
	opcodeInput     = 3
	opcodeOutput    = 4
	opcodeTerminate = 99
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
		switch opcode := c.memory[c.ic]; opcode {
		case opcodeAdd:
			var (
				addrOp1    = c.memory[c.ic+1]
				addrOp2    = c.memory[c.ic+2]
				addrResult = c.memory[c.ic+3]
			)

			c.Set(addrResult, c.Get(addrOp1)+c.Get(addrOp2))
			c.ic += 4
		case opcodeMultiply:
			var (
				addrOp1    = c.memory[c.ic+1]
				addrOp2    = c.memory[c.ic+2]
				addrResult = c.memory[c.ic+3]
			)

			c.Set(addrResult, c.Get(addrOp1)*c.Get(addrOp2))
			c.ic += 4
		case opcodeInput:
			var (
				addrResult = c.memory[c.ic+1]

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

			c.Set(addrResult, input)
			c.ic += 2
		case opcodeOutput:
			addrOp := c.memory[c.ic+1]

			fmt.Printf("Output from addr %d: %d\n", addrOp, c.Get(addrOp))
			c.ic += 2
		case opcodeTerminate:
			return nil
		default:
			return fmt.Errorf("unknown opcode %d", opcode)
		}
	}
}

func (c *Computer) Get(addr int) int {
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
