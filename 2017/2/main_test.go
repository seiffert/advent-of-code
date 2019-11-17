package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Checksum_MinMax(t *testing.T) {
	r := require.New(t)

	sum, err := Checksum(`5 1 9 5
7 5 3
2 4 6 8`, minMaxChecksum)
	r.NoError(err)
	r.Equal(18, sum)
}

func Test_Checksum_EvenlyDevisible(t *testing.T) {
	r := require.New(t)

	sum, err := Checksum(`5 9 2 8
9 4 7 3
3 8 6 5`, evenlyDevisibleChecksum)
	r.NoError(err)
	r.Equal(9, sum)
}
