package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountFish(t *testing.T) {
	require.Equal(t, 5934, CountFishAfterDays(80, "3,4,3,1,2"))
	require.Equal(t, 26984457539, CountFishAfterDays(256, "3,4,3,1,2"))
}
