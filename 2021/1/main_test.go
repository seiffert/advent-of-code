package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = []string{
	"199", "200", "208", "210", "200",
	"207", "240", "269", "260", "263",
}

func TestCountIncreases(t *testing.T) {
	require.Equal(t, 7, CountIncreases(testInput))
}

func TestCountWindowIncreases(t *testing.T) {
	require.Equal(t, 5, CountWindowIncreases(testInput, 3))
}
