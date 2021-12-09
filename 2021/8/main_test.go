package main

import (
	"testing"

	"github.com/seiffert/advent-of-code/lib"
	"github.com/stretchr/testify/require"
)

func TestCountUniqueDigitNumbers(t *testing.T) {
	require.Equal(t, 26, CountUniqueDigitNumbers(lib.MustReadFile("sample.txt")))
}

func TestSumOutputValues(t *testing.T) {
	require.Equal(t, 61229, SumOutputValues(lib.MustReadFile("sample.txt")))
}
