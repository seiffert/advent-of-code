package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalcFuelSum(t *testing.T) {
	require.Equal(t, 34241, CalcFuelSum(`12
14
1969
100756`, false))
}

func TestCalcFuel(t *testing.T) {
	require.Equal(t, 2, CalcFuel(12, false))
	require.Equal(t, 2, CalcFuel(12, true))
	require.Equal(t, 2, CalcFuel(14, false))
	require.Equal(t, 2, CalcFuel(14, true))
	require.Equal(t, 654, CalcFuel(1969, false))
	require.Equal(t, 966, CalcFuel(1969, true))
	require.Equal(t, 33583, CalcFuel(100756, false))
	require.Equal(t, 50346, CalcFuel(100756, true))
}
