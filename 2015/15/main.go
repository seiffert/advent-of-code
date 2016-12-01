package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	IngredientsRegexp = regexp.MustCompile("^([a-zA-Z]+): capacity ([-0-9]+), durability ([-0-9]+), flavor ([-0-9]+), texture ([-0-9]+), calories ([-0-9]+)$")
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	ingredients := []Ingredient{}
	for s.Scan() {
		l := s.Text()
		m := IngredientsRegexp.FindAllStringSubmatch(l, -1)

		ingredients = append(ingredients, Ingredient{
			Name:       m[0][1],
			Capacity:   i(m[0][2]),
			Durability: i(m[0][3]),
			Flavor:     i(m[0][4]),
			Texture:    i(m[0][5]),
			Calories:   i(m[0][6]),
		})
	}

	bestRecipe := bestRecipe(ingredients, &Recipe{})
	log.Printf("The best cookies are made like this:\n%s\nScore: %d\n", bestRecipe, bestRecipe.Score())
}

type Ingredient struct {
	Name       string
	Capacity   int
	Durability int
	Flavor     int
	Texture    int
	Calories   int
}

type Recipe struct {
	Ingredients map[Ingredient]int
}

func (r Recipe) Has(i Ingredient) bool {
	amount, ok := r.Ingredients[i]
	return ok && amount > 0
}

func (r Recipe) TotalAmount() int {
	sum := 0
	for _, amount := range r.Ingredients {
		sum += amount
	}
	return sum
}

func (r *Recipe) CopyFrom(other *Recipe) {
	r.reset()
	for i, a := range other.Ingredients {
		r.Add(i, a)
	}
}

func (r *Recipe) Add(i Ingredient, amount int) {
	if r.Ingredients == nil {
		r.reset()
	}
	r.Ingredients[i] = amount
}

func (r *Recipe) reset() {
	r.Ingredients = make(map[Ingredient]int)
}

func (r Recipe) String() string {
	ingredients := []string{}
	for i, amount := range r.Ingredients {
		ingredients = append(ingredients, fmt.Sprintf("%s: %d", i.Name, amount))
	}
	sort.Strings(ingredients)

	return strings.Join(ingredients, ", ")
}

func (r Recipe) Score() int64 {
	fakeIngredient := Ingredient{}
	for i, a := range r.Ingredients {
		fakeIngredient.Capacity = a*i.Capacity + fakeIngredient.Capacity
		fakeIngredient.Durability = a*i.Durability + fakeIngredient.Durability
		fakeIngredient.Flavor = a*i.Flavor + fakeIngredient.Flavor
		fakeIngredient.Texture = a*i.Texture + fakeIngredient.Texture
	}

	return int64(math.Max(0, float64(fakeIngredient.Capacity)) *
		math.Max(0, float64(fakeIngredient.Durability)) *
		math.Max(0, float64(fakeIngredient.Flavor)) *
		math.Max(0, float64(fakeIngredient.Texture)))
}

func (r Recipe) TotalCalories() int64 {
	var sum int64
	for i, a := range r.Ingredients {
		sum += int64(a) * int64(i.Calories)
	}
	return sum
}

func i(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func bestRecipe(ingredients []Ingredient, recipe *Recipe) *Recipe {
	var best *Recipe
	for _, i := range ingredients {
		if !recipe.Has(i) {
			for amount := 1; amount <= 100-recipe.TotalAmount(); amount++ {
				newRecipe := &Recipe{}
				newRecipe.CopyFrom(recipe)

				newRecipe.Add(i, amount)
				newRecipe = bestRecipe(ingredients, newRecipe)

				if best == nil || newRecipe.TotalAmount() == 100 && newRecipe.TotalCalories() == 500 && newRecipe.Score() > best.Score() {
					best = newRecipe
				}
			}
		}
	}
	if best != nil {
		return best
	}

	return recipe
}
