package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdapterList1(t *testing.T) {
	a := NewAdapterList(`16
10
15
5
1
11
7
19
6
12
4`)

	d1, _, d3 := a.Distribution()
	require.Equal(t, 7, d1)
	require.Equal(t, 5, d3)

	require.EqualValues(t, 8, a.PossibleCombinations())
}

func TestAdapterList2(t *testing.T) {
	a := NewAdapterList(`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`)

	d1, _, d3 := a.Distribution()
	require.Equal(t, 22, d1)
	require.Equal(t, 10, d3)

	require.EqualValues(t, 19208, a.PossibleCombinations())
}
