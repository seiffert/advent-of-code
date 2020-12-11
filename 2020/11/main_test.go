package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimulation1(t *testing.T) {
	sm := NewSeatMap(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)

	require.Equal(t, 37, sm.RunSimulation(4, true))
}
func TestSimulation2(t *testing.T) {
	sm := NewSeatMap(`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`)

	require.Equal(t, 26, sm.RunSimulation(5, false))
}
