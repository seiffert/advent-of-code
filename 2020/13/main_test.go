package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBusSchedule(t *testing.T) {
	s := NewSchedule(`7,13,x,x,59,x,31,19`)

	busID, at := s.A(939)
	require.Equal(t, 59, busID)
	require.Equal(t, 944, at)

	require.EqualValues(t, 1068781, s.B())
}

func TestOtherBusSchedule(t *testing.T) {
	cases := map[string]int64{
		"17,x,13,19":      3417,
		"67,7,59,61":      754018,
		"67,x,7,59,61":    779210,
		"67,7,x,59,61":    1261476,
		"1789,37,47,1889": 1202161486,
	}

	for in, res := range cases {
		t.Run(in, func(t *testing.T) {
			require.EqualValues(t, res, NewSchedule(in).B())
		})
	}
}
