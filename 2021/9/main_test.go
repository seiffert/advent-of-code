package main

import (
	"testing"

	"github.com/seiffert/advent-of-code/lib"
	"github.com/stretchr/testify/require"
)

func TestSumLowPointsRiskLevels(t *testing.T) {
	hm := NewHeightMap(lib.MustReadFile("sample.txt"))
	lps := hm.LowPoints()

	require.Equal(t, 15, hm.SumRiskLevels(lps))
}

func TestMultiplyBiggestBasinSizes(t *testing.T) {
	hm := NewHeightMap(lib.MustReadFile("sample.txt"))

	require.Equal(t, 1134, hm.MultiplyBiggestBasinSizes(3))
}
