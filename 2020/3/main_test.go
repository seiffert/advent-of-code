package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountTrees(t *testing.T) {
	input := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

	require.Equal(t, 7, NewMap(input).CountTrees(Slope{3, 1}))
}
