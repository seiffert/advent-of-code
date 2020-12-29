package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGameA(t *testing.T) {
	g := NewGame(`Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`)

	require.Equal(t, 306, g.PlayA())
}

func TestGameB(t *testing.T) {
	g := NewGame(`Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`)

	require.Equal(t, 291, g.PlayB())
}

func TestExitConditionB(t *testing.T) {
	g := NewGame(`Player 1:
43
19

Player 2:
2
29
14`)
	g.PlayB()
}
