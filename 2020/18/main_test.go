package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculate(t *testing.T) {
	require.EqualValues(t, 51, CalculateWithoutPrecedence(`1 + (2 * 3) + (4 * (5 + 6))`))
	require.EqualValues(t, 51, CalculateWithPrecedence(`1 + (2 * 3) + (4 * (5 + 6))`))
	require.EqualValues(t, 26, CalculateWithoutPrecedence(`2 * 3 + (4 * 5)`))
	require.EqualValues(t, 46, CalculateWithPrecedence(`2 * 3 + (4 * 5)`))
	require.EqualValues(t, 437, CalculateWithoutPrecedence(`5 + (8 * 3 + 9 + 3 * 4 * 3)`))
	require.EqualValues(t, 1445, CalculateWithPrecedence(`5 + (8 * 3 + 9 + 3 * 4 * 3)`))
	require.EqualValues(t, 12240, CalculateWithoutPrecedence(`5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))`))
	require.EqualValues(t, 669060, CalculateWithPrecedence(`5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))`))
	require.EqualValues(t, 13632, CalculateWithoutPrecedence(`((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`))
	require.EqualValues(t, 23340, CalculateWithPrecedence(`((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`))
}
