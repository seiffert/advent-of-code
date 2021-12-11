package main

import (
	"testing"

	"github.com/seiffert/advent-of-code/lib"
	"github.com/stretchr/testify/require"
)

func TestCountFlashes(t *testing.T) {
	require.Equal(t, 1656, CountFlahes(100, lib.MustReadFile("sample.txt")))
}

func TestFirstSyncedFlash(t *testing.T) {
	require.Equal(t, 195, FirstSyncedFlash(lib.MustReadFile("sample.txt")))
}
