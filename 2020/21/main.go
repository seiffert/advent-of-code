package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/seiffert/advent-of-code/lib"
)

func main() {
	m := NewMenu(lib.MustReadFile("input.txt"))
	m.DetermineAllergens()

	fmt.Println("allergenic ingredients: ", m.CountAllergenicIngredients())
	fmt.Println("canonical dangerous ingredients: ", m.CanonicalDangerousIngredients())
}

type (
	Menu struct {
		dishes    []*dish
		allergens map[string]string
	}
	dish struct {
		ingredients []string
		allergens   []string
	}
)

func NewMenu(input string) *Menu {
	var dishes []*dish
	for _, line := range strings.Split(input, "\n") {
		p := strings.Split(line, " (contains ")
		dishes = append(dishes, &dish{
			ingredients: strings.Split(p[0], " "),
			allergens:   strings.Split(strings.Trim(p[1], ")"), ", "),
		})
	}
	return &Menu{dishes: dishes}
}

func (m *Menu) CountAllergenicIngredients() int {
	var allergenicIngredients int
	for _, d := range m.dishes {
		ingredientsWithoutAllergens := d.ingredients
		for _, ingredient := range m.allergens {
			ingredientsWithoutAllergens = without(ingredientsWithoutAllergens, ingredient)
		}
		allergenicIngredients += len(ingredientsWithoutAllergens)
	}

	return allergenicIngredients
}

func (m *Menu) CanonicalDangerousIngredients() string {
	var dangerous []string
	for allergen, _ := range m.allergens {
		dangerous = append(dangerous, allergen)
	}
	sort.Strings(dangerous)

	sortedIngredients := make([]string, 0, len(dangerous))
	for _, allergen := range dangerous {
		sortedIngredients = append(sortedIngredients, m.allergens[allergen])
	}

	return strings.Join(sortedIngredients, ",")
}

func (m *Menu) DetermineAllergens() {
	m.allergens = m.guessAllergens(m.allAllergens())
}

func (m *Menu) guessAllergens(allergens []string) map[string]string {
	candidates := map[string][]string{}
	for _, allergen := range allergens {
		candidates[allergen] = m.ingredientsPresentInAllDishesWithAllergen(allergen)
	}

	assumptions := map[string]string{}
	for len(candidates) > 0 {
		for allergen, ingredients := range candidates {
			switch len(ingredients) {
			case 1:
				assumptions[allergen] = ingredients[0]
				delete(candidates, allergen)
			case 0:
				// incorrect solution, next attempt (we randomize by iterating
				// over the map of candidates).
				return m.guessAllergens(allergens)
			}
			for o, is := range candidates {
				if o == allergen {
					continue
				}
				candidates[o] = without(is, ingredients[0])
			}
		}
	}
	return assumptions
}

func (m *Menu) ingredientsPresentInAllDishesWithAllergen(allergen string) []string {
	var result []string
	for _, d := range m.dishes {
		if !d.containsAllergen(allergen) {
			continue
		}

		if len(result) == 0 {
			result = d.ingredients
		} else {
			var nr []string
			for _, r := range result {
				if d.containsIngredient(r) {
					nr = append(nr, r)
				}
			}
			result = nr
		}
	}
	return result
}

func (m *Menu) allAllergens() []string {
	var allergens []string
	for _, d := range m.dishes {
		var found bool
		for _, allergen := range d.allergens {
			for _, foundAllergen := range allergens {
				if allergen == foundAllergen {
					found = true
				}
			}
			if !found {
				allergens = append(allergens, allergen)
			}
		}
	}
	return allergens
}

func (d *dish) containsIngredient(ingredient string) bool {
	return contains(d.ingredients, ingredient)
}

func (d *dish) containsAllergen(allergen string) bool {
	return contains(d.allergens, allergen)
}

func contains(list []string, item string) bool {
	for _, li := range list {
		if li == item {
			return true
		}
	}
	return false
}

func without(list []string, item string) []string {
	var nl []string
	for _, li := range list {
		if li != item {
			nl = append(nl, li)
		}
	}
	return nl
}
