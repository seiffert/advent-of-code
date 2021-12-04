package main

import (
	"testing"

	"github.com/seiffert/advent-of-code/lib"
	"github.com/stretchr/testify/require"
)

func TestPowerConsumption(t *testing.T) {
	in := lib.MustReadFile("sample.txt")

	require.Equal(t, 198, PowerConsumption(in))
}

func TestLifeSupportRating(t *testing.T) {
	in := lib.MustReadFile("sample.txt")

	require.Equal(t, 230, LifeSupportRating(in))
}
