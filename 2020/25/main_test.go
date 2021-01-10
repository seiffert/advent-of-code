package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalcEncryptionKey(t *testing.T) {
	require.Equal(t, 14897079, CalcEncryptionKey(5764801, 17807724))
}

func TestCalcLoopSize(t *testing.T) {
	require.Equal(t, 11, CalcLoopSize(7, 17807724))
	require.Equal(t, 8, CalcLoopSize(7, 5764801))
}

func TestTransform(t *testing.T) {
	require.Equal(t, 14897079, transform(8, 17807724))
	require.Equal(t, 14897079, transform(11, 5764801))
}
