package intcode

import (
	"fmt"
)

const (
	ModePosition  = 0
	ModeImmediate = 1
	ModeRelBase   = 2

	opcodeAdd         = 1
	opcodeMultiply    = 2
	opcodeInput       = 3
	opcodeOutput      = 4
	opcodeJumpIfTrue  = 5
	opcodeJumpIfFalse = 6
	opcodeLessThan    = 7
	opcodeEquals      = 8
	opcodeSetRelBase  = 9
	opcodeTerminate   = 99
)

func NewComputer(p []int) *Computer {
	mem := append([]int(nil), p...)
	return &Computer{
		memory: mem,
		input:  StdinReceiver,
		output: StdoutSender,
	}
}

type Computer struct {
	memory  []int
	ic      int
	relBase int

	input  InputReceiver
	output OutputSender
}

func (c *Computer) SetInputReceiver(in InputReceiver) {
	c.input = in
}

func (c *Computer) SetOutputSender(out OutputSender) {
	c.output = out
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
			c.Set(c.Get(c.ic+1, ModeImmediate), c.input())
			c.ic += 2
		case opcodeOutput:
			c.output(c.Get(c.ic+1, param1Mode))
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
		case opcodeSetRelBase:
			c.relBase += c.Get(c.ic+1, param1Mode)
			c.ic += 2
		case opcodeTerminate:
			return nil
		default:
			return fmt.Errorf("unknown opcode %d", opcode)
		}
	}
}

func (c *Computer) Get(addr, mode int) int {
	switch mode {
	case ModePosition:
		addr = c.Get(addr, ModeImmediate)
	case ModeRelBase:
		addr = c.Get(addr, ModeImmediate) + c.relBase
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
