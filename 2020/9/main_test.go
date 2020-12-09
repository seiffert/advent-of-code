package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncryptedStream(t *testing.T) {
	es := NewEncryptedStream(`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`)

	require.EqualValues(t, 127, es.FindFirstInvalidNumber(5))

	min, max := es.FindContiguousSetWithSum(127)
	require.EqualValues(t, 62, min+max)
}
