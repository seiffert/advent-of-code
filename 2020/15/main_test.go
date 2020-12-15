package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGame(t *testing.T) {
	require.Equal(t, 436, NewGame(`0,3,6`).Play(2020))
	require.Equal(t, 1, NewGame(`1,3,2`).Play(2020))
	require.Equal(t, 10, NewGame(`2,1,3`).Play(2020))
	require.Equal(t, 27, NewGame(`1,2,3`).Play(2020))
	require.Equal(t, 78, NewGame(`2,3,1`).Play(2020))
	require.Equal(t, 438, NewGame(`3,2,1`).Play(2020))
	require.Equal(t, 1836, NewGame(`3,1,2`).Play(2020))

	require.Equal(t, 175594, NewGame(`0,3,6`).Play(30000000))
	require.Equal(t, 2578, NewGame(`1,3,2`).Play(30000000))
	require.Equal(t, 3544142, NewGame(`2,1,3`).Play(30000000))
	require.Equal(t, 261214, NewGame(`1,2,3`).Play(30000000))
	require.Equal(t, 6895259, NewGame(`2,3,1`).Play(30000000))
	require.Equal(t, 18, NewGame(`3,2,1`).Play(30000000))
	require.Equal(t, 362, NewGame(`3,1,2`).Play(30000000))
}
