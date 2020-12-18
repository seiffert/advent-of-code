package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test3DGrid(t *testing.T) {
	g := New3DGrid(`.#.
..#
###`)

	require.Equal(t, 5, g.CountActive())

	g.Play(6)

	require.Equal(t, 112, g.CountActive())
}
func Test4DGrid(t *testing.T) {
	g := New4DGrid(`.#.
..#
###`)

	require.Equal(t, 5, g.CountActive())

	g.Play(6)

	require.Equal(t, 848, g.CountActive())
}
