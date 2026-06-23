package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPreparationTime int) int {
    if avgPreparationTime == 0 {
        avgPreparationTime = 2
    }

    return len(layers) * avgPreparationTime
}

// TODO: define the 'Quantities()' function
func Quantities(ingredients []string) (noodles int, sauce float64) {
	for _, layer := range ingredients {
		switch layer {
		case "sauce":
			sauce += 0.2

		case "noodles":
			noodles += 50
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
    friendLastItem := friendsList[len(friendsList)-1]
    myList[len(myList)-1] = friendLastItem
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portionsNum int) (scaledQuantities []float64) {
	for _, q := range quantities {
		quantityForOne := q / 2.0
		scaledQuantities = append(scaledQuantities, quantityForOne*float64(portionsNum))
	}
	return
}
