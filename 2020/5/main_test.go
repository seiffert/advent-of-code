package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBoardingPass_SeatID(t *testing.T) {
	require.Equal(t, 567, NewSeat("BFFFBBFRRR").ID())
	require.Equal(t, 119, NewSeat("FFFBBBFRRR").ID())
	require.Equal(t, 820, NewSeat("BBFFBBFRLL").ID())
}
