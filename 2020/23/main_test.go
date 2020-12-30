package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRing_Play(t *testing.T) {
	r := NewRing("389125467").Play(10)

	one := r.Find[1]
	require.Equal(t, "92658374", one.Next.StringUntil(1))

	r = r.Play(90)
	require.Equal(t, "67384529", one.Next.StringUntil(1))

	r = NewRing("389125467").AddUntil(1000000).Play(10000000)

	require.Equal(t, 934001, r.Find[1].Next.v)
	require.Equal(t, 159792, r.Find[1].Next.Next.v)
	require.Equal(t, 149245887792, r.Find[1].Next.v*r.Find[1].Next.Next.v)
}
