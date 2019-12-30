package intcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestComputer_Calculate(t *testing.T) {
	testCases := map[string]string{
		"1,0,0,0,99":          "2,0,0,0,99",
		"2,3,0,3,99":          "2,3,0,6,99",
		"2,4,4,5,99,0":        "2,4,4,5,99,9801",
		"1,1,1,4,99,5,6,0,99": "30,1,1,4,2,5,6,0,99",
	}

	for in, out := range testCases {
		t.Run(in, func(t *testing.T) {
			r := require.New(t)

			p := NewComputer(Parse(in))
			_, err := p.Calculate()
			r.NoError(err)

			r.Equal(Parse(out), p.memory)
		})
	}
}
