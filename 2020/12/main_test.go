package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShip_Drive(t *testing.T) {
	s := NewShip(`F10
N3
F7
R90
F11`)

	s.Drive()

	require.Equal(t, 25, s.ManhattenDist())
}

func TestShip_DriveWithWaypoint(t *testing.T) {
	s := NewShip(`F10
N3
F7
R90
F11`)

	s.DriveWithWaypoint()

	require.Equal(t, 286, s.ManhattenDist())
}
