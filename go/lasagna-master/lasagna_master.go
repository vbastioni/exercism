package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		switch layer {
		case "sauce":
			sauce += .2
		case "noodles":
			noodles += 50
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friends, mine []string) {
	mine[len(mine)-1] = friends[len(friends)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(inputs []float64, portions int) (out []float64) {
	out = make([]float64, len(inputs))
	for i, e := range inputs {
		out[i] = e * float64(portions) / 2
	}
	return
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
