package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	p := NewProgram(lib.MustReadFile("input.txt"))
	fmt.Printf("sum: %d\n", p.SumMemory())

	p = NewProgramB(lib.MustReadFile("input.txt"))
	fmt.Printf("sum: %d\n", p.SumMemory())
}

type Program struct {
	mem map[uint64]uint64
}

func NewProgram(input string) *Program {
	p := &Program{
		mem: make(map[uint64]uint64),
	}
	var and, or uint64
	for _, line := range strings.Split(input, "\n") {
		switch {
		case strings.HasPrefix(line, "mask"):
			var newAnd uint64
			var newOr uint64
			for i, c := range strings.TrimPrefix(line, "mask = ") {
				switch c {
				case '0':
					newAnd = newAnd | 1<<(35-i)
				case '1':
					newOr = newOr | 1<<(35-i)
				}
			}
			and, or = ^newAnd, newOr
		case strings.HasPrefix(line, "mem"):
			a := strings.Split(strings.TrimPrefix(line, "mem["), "] = ")
			addr, _ := strconv.Atoi(a[0])
			val, _ := strconv.ParseInt(a[1], 10, 64)
			p.mem[uint64(addr)] = uint64(val)&and | or
		}
	}
	return p
}

func NewProgramB(input string) *Program {
	p := &Program{
		mem: make(map[uint64]uint64),
	}
	var mask string
	for _, line := range strings.Split(input, "\n") {
		switch {
		case strings.HasPrefix(line, "mask"):
			mask = strings.TrimPrefix(line, "mask = ")
		case strings.HasPrefix(line, "mem"):
			a := strings.Split(strings.TrimPrefix(line, "mem["), "] = ")
			addr, _ := strconv.Atoi(a[0])
			val, _ := strconv.ParseInt(a[1], 10, 64)
			for _, addr := range applyBMask(addr, mask) {
				p.mem[addr] = uint64(val)
			}
		}
	}
	return p
}

func (p Program) SumMemory() (result int64) {
	for _, v := range p.mem {
		result += int64(v)
	}
	return
}

func applyBMask(addr int, mask string) []uint64 {
	s := []uint64{uint64(addr)}

	for i, c := range mask {
		var n []uint64
		switch c {
		case '0':
			continue
		case '1':
			for _, a := range s {
				n = append(n, a|(1<<(35-i)))
			}
		case 'X':
			for _, a := range s {
				n = append(n,
					a|(1<<(35-i)),
					a & ^(1<<(35-i)),
				)
			}
		}
		s = n
	}
	return s
}
