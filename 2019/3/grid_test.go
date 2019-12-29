package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGrid_FindClosestIntersectionDistance(t *testing.T) {
	testCases := []struct {
		firstWire  string
		secondWire string
		result     int
	}{{
		"R8,U5,L5,D3",
		"U7,R6,D4,L4",
		6,
	}, {
		"R75,D30,R83,U83,L12,D49,R71,U7,L72",
		"U62,R66,U55,R34,D71,R55,D58,R83",
		159,
	}, {
		"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
		"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
		135,
	}}

	for i, testCase := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			g := NewGrid()
			g.AddWire("0", testCase.firstWire)
			g.AddWire("1", testCase.secondWire)

			require.Equal(t,
				testCase.result, g.FindClosestIntersectionDistance(),
			)
		})
	}
}

func TestGrid_FindShortestIntersectingPath(t *testing.T) {
	testCases := []struct {
		firstWire  string
		secondWire string
		result     int
	}{{
		"R8,U5,L5,D3",
		"U7,R6,D4,L4",
		30,
	}, {
		"R75,D30,R83,U83,L12,D49,R71,U7,L72",
		"U62,R66,U55,R34,D71,R55,D58,R83",
		610,
	}, {
		"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
		"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
		410,
	}}

	for i, testCase := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			g := NewGrid()
			g.AddWire("0", testCase.firstWire)
			g.AddWire("1", testCase.secondWire)

			require.Equal(t,
				testCase.result, g.FindShortestIntersectingPath(),
			)
		})
	}

}
