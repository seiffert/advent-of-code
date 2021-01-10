package intcode

import (
	"strconv"
	"strings"
)

const (
	opCodeAdd = 1
	opCodeMul = 2
)

func New(input string) *IntCode {
	var (
		as   = strings.Split(input, ",")
		ints = make([]int, len(as))
	)
	for i, a := range strings.Split(input, ",") {
		n, _ := strconv.Atoi(a)
		ints[i] = n
	}
	return &IntCode{
		mem: ints,
	}
}

type IntCode struct {
	mem []int
	ip  int
}

func (c *IntCode) Run() *IntCode {
	for c.Get(c.ip) != 99 {
		c.step()
	}
	return c
}

func (c *IntCode) step() {
	switch c.mem[c.ip] {
	case opCodeAdd:
		c.Set(c.GetRel(3), c.Get(c.GetRel(1))+c.Get(c.GetRel(2)))
		c.Advance(4)
	case opCodeMul:
		c.Set(c.GetRel(3), c.Get(c.GetRel(1))*c.Get(c.GetRel(2)))
		c.Advance(4)
	}
}

func (c *IntCode) Get(ip int) int {
	return c.mem[ip]
}

func (c *IntCode) GetRel(offset int) int {
	return c.mem[c.ip+offset]
}

func (c *IntCode) Set(ip, val int) {
	c.mem[ip] = val
}

func (c *IntCode) Advance(offset int) {
	c.ip += offset
}
