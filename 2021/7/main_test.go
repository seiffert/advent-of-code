package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRequiredFuel(t *testing.T) {
	require.Equal(t, 37, RequiredFuel("16,1,2,0,4,2,7,1,2,14", true))
	require.Equal(t, 168, RequiredFuel("16,1,2,0,4,2,7,1,2,14", false))
}
