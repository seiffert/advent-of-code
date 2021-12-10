package main

import (
	"testing"

	"github.com/seiffert/advent-of-code/lib"
	"github.com/stretchr/testify/require"
)

func TestTotalSyntaxErrorScore(t *testing.T) {
	require.Equal(t, 26397, TotalSyntaxErrorScore(lib.MustReadFile("sample.txt")))
}

func TestMiddleCompletionScore(t *testing.T) {
	require.Equal(t, 288957, MiddleCompletionScore(lib.MustReadFile("sample.txt")))
}
