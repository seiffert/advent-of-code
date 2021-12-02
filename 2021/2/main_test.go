package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testInput = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

func TestNavigate(t *testing.T) {
	require.Equal(t,
		Position{15, 10, 0},
		Navigate(testInput),
	)
}

func TestNavigateCorrectly(t *testing.T) {
	require.Equal(t,
		Position{15, 60, 10},
		NavigateCorrectly(testInput),
	)
}
