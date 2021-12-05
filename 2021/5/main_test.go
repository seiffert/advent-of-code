package main

import (
	"testing"

	"github.com/seiffert/advent-of-code/lib"
	"github.com/stretchr/testify/require"
)

func TestCountOverlaps(t *testing.T) {
	g := NewGame(lib.MustReadFile("sample.txt"), false)

	require.Equal(t, 5, g.CountOverlaps())
}

func TestCountOverlapsWithDiagonals(t *testing.T) {
	g := NewGame(lib.MustReadFile("sample.txt"), true)

	require.Equal(t, 12, g.CountOverlaps())
}
