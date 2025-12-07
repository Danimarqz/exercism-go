package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) (totalTime int) {
	totalTime = timePerLayer * len(layers)
	return
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for i := range layers {
		switch layers[i] {
		case "sauce":
			sauce += 0.2
		case "noodles":
			noodles += 50
		}
	}
	return
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, ownList []string) {
	ownList[len(ownList) - 1] = friendList [len(friendList) - 1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, wantedPortions int) []float64 {
	scaled := make([]float64, len(quantities))
	copy(scaled, quantities)

	factor := float64(wantedPortions) / 2
	for i := range scaled {
		scaled[i] *= factor
	}
	return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
