package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Ingredient struct {
	capacity   int
	durability int
	flavour    int
	texture    int
	calories   int
}

type Recipe map[string]int

var inputFile = flag.String("inputFile", "inputs/day15.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	lines := strings.Split(contents, "\n")
	rx := regexp.MustCompile("(\\w+): capacity (-?\\d+), durability (-?\\d+), flavor (-?\\d+), texture (-?\\d+), calories (-?\\d+)")
	pantry := make(map[string]Ingredient)
	for l := range lines {
		submatches := rx.FindStringSubmatch(lines[l])
		capacity, _ := strconv.Atoi(submatches[2])
		durability, _ := strconv.Atoi(submatches[3])
		flavor, _ := strconv.Atoi(submatches[4])
		texture, _ := strconv.Atoi(submatches[5])
		calories, _ := strconv.Atoi(submatches[6])
		pantry[submatches[1]] = Ingredient{capacity: capacity, durability: durability, flavour: flavor, texture: texture, calories: calories}
	}

	space := 100

	ingredients := make([]string, len(pantry))
	i := 0
	for k, _ := range pantry {
		ingredients[i] = k
		i++
	}
	recipe := make(Recipe)
	fmt.Println(buildRecipe(recipe, ingredients, pantry, space))

}

func buildRecipe(recipe Recipe, ingredients []string, pantry map[string]Ingredient, spaceLeft int) int {
	largest := 0

	ingredient := ingredients[0]
	if len(ingredients) == 1 {
		recipe[ingredient] = spaceLeft
		score, calories := evaluateRecipe(recipe, pantry)
		if calories == 500 {
			return score
		} else {
			return -1
		}

	} else {
		remainingIngredients := make([]string, len(ingredients)-1)

		copy(remainingIngredients[:], ingredients[1:])

		for i := 0; i <= spaceLeft; i++ {
			newRecipe := Recipe{}
			for k, v := range recipe {
				newRecipe[k] = v
			}
			newRecipe[ingredient] = i
			value := buildRecipe(newRecipe, remainingIngredients, pantry, spaceLeft-i)
			if value > largest {
				largest = value
			}

		}
	}
	return largest
}

func evaluateRecipe(recipe Recipe, ingredients map[string]Ingredient) (int, int) {
	capacity := 0
	durability := 0
	flavour := 0
	texture := 0
	calories := 0

	for r := range recipe {
		ing := ingredients[r]

		capacity += recipe[r] * ing.capacity
		durability += recipe[r] * ing.durability
		flavour += recipe[r] * ing.flavour
		texture += recipe[r] * ing.texture
		calories += recipe[r] * ing.calories
	}

	if capacity <= 0 || durability <= 0 || flavour <= 0 || texture <= 0 {
		return 0, 0
	}

	return capacity * texture * durability * flavour, calories
}
