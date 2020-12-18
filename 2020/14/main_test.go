package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProgram(t *testing.T) {
	p := NewProgram(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`)

	require.EqualValues(t, 165, p.SumMemory())
}

func TestProgramB(t *testing.T) {
	p := NewProgramB(`mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`)

	require.EqualValues(t, 208, p.SumMemory())
}