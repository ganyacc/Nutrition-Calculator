package main

import "fmt"

func main() {

	ns := GetNutrionalScore(NutrionalData{

		Energy:              EnergyFromKcal(150),
		Suger:               SugerGram(5),
		SaturatedFattyAcids: SaturatedFattyAcids(5),
		Sodium:              SodiumMilligram(20),
		Fruits:              FruitsPercent(40),
		Fibre:               FibreGram(30),
		Protein:             ProteinGram(135),
	}, Beverages)

	fmt.Printf("Nutritional Score is %d\n", ns.Negative)
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())
}
