package main

import (
	"testing"

	"github.com/seiffert/advent-of-code/lib"
	"github.com/stretchr/testify/require"
)

func TestElementCounts(t *testing.T) {
	p := NewPolymer(lib.MustReadFile("sample.txt"))

	// Part 1
	p.Expand(10)
	c := p.ElementCounts()

	require.Equal(t, 1749, c[len(c)-1])
	require.Equal(t, 161, c[0])

	// Part 2
	p.Expand(30)
	c = p.ElementCounts()

	require.Equal(t, 2192039569602, c[len(c)-1])
	require.Equal(t, 3849876073, c[0])
}
