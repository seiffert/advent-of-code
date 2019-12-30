package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValidPassword1(t *testing.T) {
	testCases := []struct {
		input int
		valid bool
	}{{
		111111, true,
	}, {
		223450, false,
	}, {
		123789, false,
	}}

	for _, testCase := range testCases {
		t.Run(strconv.Itoa(testCase.input), func(t *testing.T) {
			require.Equal(t, testCase.valid, IsValidPassword1(testCase.input))
		})
	}
}

func TestIsValidPassword2(t *testing.T) {
	testCases := []struct {
		input int
		valid bool
	}{{
		112233, true,
	}, {
		123444, false,
	}, {
		111122, true,
	}}

	for _, testCase := range testCases {
		t.Run(strconv.Itoa(testCase.input), func(t *testing.T) {
			require.Equal(t, testCase.valid, IsValidPassword2(testCase.input))
		})
	}
}
