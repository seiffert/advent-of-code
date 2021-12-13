package main

import (
	"testing"

	"github.com/seiffert/advent-of-code/lib"
	"github.com/stretchr/testify/require"
)

func TestFoldCount(t *testing.T) {
	g, folds := NewPaper(lib.MustReadFile("sample.txt"))

	g.Fold(folds[0])
	require.Equal(t, 17, g.CountDots())
}
