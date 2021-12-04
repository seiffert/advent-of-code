package main

import (
	"testing"

	"github.com/seiffert/advent-of-code/lib"
	"github.com/stretchr/testify/require"
)

func TestGamePart1(t *testing.T) {
	g := NewGame(lib.MustReadFile("sample.txt"))
	g.Play()

	require.Equal(t, 4512, g.WinnerScore())
}

func TestGamePart2(t *testing.T) {
	g := NewGame(lib.MustReadFile("sample.txt"))
	g.Play()

	require.Equal(t, 1924, g.LastWinnerScore())
}
