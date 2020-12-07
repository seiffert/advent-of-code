package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSumGroups(t *testing.T) {
	l := NewAnswerList(`abc

a
b
c

ab
ac

a
a
a
a

b`)

	require.Equal(t, 11, l.SumGroups(func(g GroupAnswers) int {
		return g.PositivelyAnsweredQuestions()
	}))
	require.Equal(t, 6, l.SumGroups(func(g GroupAnswers) int {
		return g.UnanimousPositivelyAnsweredQuestions()
	}))
}
