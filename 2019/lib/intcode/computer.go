package intcode

import "fmt"

const (
	opcodeAdd       = 1
	opcodeMultiply  = 2
	opcodeTerminate = 99
)

func NewComputer(p []int, noun, verb int) *Computer {
	mem := append([]int(nil), p...)
	mem[1] = noun
	mem[2] = verb
	return &Computer{
		memory: mem,
	}
}

type Computer struct {
	memory []int
	ic     int
}

func (c *Computer) Calculate() (int, error) {
	for {
		switch opcode := c.memory[c.ic]; opcode {
		case opcodeAdd:
			var (
				addrOp1    = c.memory[c.ic+1]
				addrOp2    = c.memory[c.ic+2]
				addrResult = c.memory[c.ic+3]
			)

			c.set(addrResult, c.get(addrOp1)+c.get(addrOp2))
			c.ic += 4
		case opcodeMultiply:
			var (
				addrOp1    = c.memory[c.ic+1]
				addrOp2    = c.memory[c.ic+2]
				addrResult = c.memory[c.ic+3]
			)

			c.set(addrResult, c.get(addrOp1)*c.get(addrOp2))
			c.ic += 4
		case opcodeTerminate:
			return c.memory[0], nil
		default:
			return 0, fmt.Errorf("unknown opcode %d", opcode)
		}
	}
}

func (c *Computer) get(addr int) int {
	if addr < len(c.memory) {
		return c.memory[addr]
	}
	return 0
}

func (c *Computer) set(addr, val int) {
	if addr >= len(c.memory) {
		c.growMemory(addr - len(c.memory) + 1)
	}

	c.memory[addr] = val
}

func (c *Computer) growMemory(diff int) {
	c.memory = append(c.memory, make([]int, diff)...)
}
