package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMenu(t *testing.T) {
	m := NewMenu(`mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`)
	m.DetermineAllergens()

	require.Equal(t, 5, m.CountAllergenicIngredients())
	require.Equal(t, "mxmxvkd,sqjhc,fvjkl", m.CanonicalDangerousIngredients())
}
