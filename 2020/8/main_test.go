package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBootSequence(t *testing.T) {
	b := NewBootSequence(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`)

	require.Equal(t, errInfiniteLoopDetected, b.Run())
	require.Equal(t, 5, b.Acc)
}
