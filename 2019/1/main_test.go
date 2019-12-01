package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFuelForMass(t *testing.T) {
	testCases := map[int]int{
		12:     2,
		14:     2,
		1969:   654,
		100756: 33583,
	}
	for mass, expected := range testCases {
		t.Run(fmt.Sprintf("fuel for mass of %d", mass), func(t *testing.T) {
			require.Equal(t, expected, FuelForMass(mass))
		})
	}
}

func TestFuelForMassAndFuel(t *testing.T) {
	testCases := map[int]int{
		14:     2,
		1969:   966,
		100756: 50346,
	}
	for mass, expected := range testCases {
		t.Run(fmt.Sprintf("fuel for mass and fuel of %d", mass), func(t *testing.T) {
			require.Equal(t, expected, FuelForMassAndFuel(mass))
		})
	}
}
