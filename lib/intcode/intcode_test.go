package intcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntCode(t *testing.T) {
	require.Equal(t, 2, New("1,0,0,0,99").Run().Get(0))
	require.Equal(t, 9801, New("2,4,4,5,99,0").Run().Get(5))
	require.Equal(t, 30, New("1,1,1,4,99,5,6,0,99").Run().Get(0))
}
