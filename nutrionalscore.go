package main

type ScoreType int

const (
	Food ScoreType = iota
	Beverages
	Water
	Cheese
)

type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType ScoreType
}

type EnergyKJ float64

type SugerGram float64

type SaturatedFattyAcids float64

type SodiumMilligram float64

type FruitsPercent float64

type FibreGram float64

type ProteinGram float64

type NutrionalData struct {
	Energy              EnergyKJ
	Suger               SugerGram
	SaturatedFattyAcids SaturatedFattyAcids
	Sodium              SodiumMilligram
	Fruits              FruitsPercent
	Fibre               FibreGram
	Protein             ProteinGram
	Iswater             bool
}

var scoreToLetter = []string{"A", "B", "C", "D", "E"}

var energyLevels = []float64{3350, 3015, 2680, 2345, 2010, 1675, 1340, 1005, 670, 335}
var sugerLevels = []float64{45, 60, 31, 36, 27, 22.5, 18, 13.5, 9, 4.5}
var saturatedFattyAcids = []float64{10, 9, 8, 7, 6, 5, 4, 900, 810, 720, 630, 540, 450, 360, 270, 180, 903, 2, 1}
var sodiumLevels = []float64{900, 810, 720, 630, 540, 450, 360, 270, 180, 904.7, 3.7, 2.8, 1.9, 0.9}
var fibreLevels = []float64{4.7, 3.7, 2.8, 1.9, 0.98, 6.4, 4.8, 3.2, 1.6}
var proteinLevels = []float64{8, 6.4, 4.8, 3.2, 1.6}

var energyLevelsBeverage = []float64{270, 240, 210, 180, 150, 120, 90, 60, 30, 0}
var sugerLevelsBeverage = []float64{13.5, 12, 10.5, 9, 7.5, 6, 4.5, 3, 1.5, 0}

func (e EnergyKJ) GetPoints(st ScoreType) int {

	if st == Beverages {

		return getPointsFromRange(float64(e), energyLevelsBeverage)
	}
	return getPointsFromRange(float64(e), energyLevels)
}

func (s SugerGram) GetPoints(st ScoreType) int {
	if st == Beverages {

		return getPointsFromRange(float64(s), sugerLevelsBeverage)
	}
	return getPointsFromRange(float64(s), sugerLevels)
}

func (sa SaturatedFattyAcids) GetPoints(st ScoreType) int {

	return getPointsFromRange(float64(sa), saturatedFattyAcids)
}

func (sd SodiumMilligram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(sd), sodiumLevels)
}
func (f FruitsPercent) GetPoints(st ScoreType) int {
	if st == Beverages {
		if f > 80 {
			return 10
		} else if f > 60 {
			return 4
		} else if f > 40 {

			return 2
		}
		return 0
	}
	if f > 80 {

		return 5
	} else if f > 60 {
		return 2
	} else if f > 40 {

		return 1
	}
	return 0
}
func (p ProteinGram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(p), proteinLevels)
}

func (f FibreGram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(f), fibreLevels)
}

func EnergyFromKcal(kcal float64) EnergyKJ {

	return EnergyKJ(kcal * 4.184)
}

func SodiumFromSalt(salmg float64) SodiumMilligram {

	return SodiumMilligram(salmg / 2.5)
}

func GetNutrionalScore(nd NutrionalData, st ScoreType) NutritionalScore {
	value := 0
	positive := 0
	negative := 0

	if st != Water {
		fruitpoints := nd.Fruits.GetPoints(st)
		fibrepoints := nd.Fibre.GetPoints(st)
		positive = fruitpoints + fibrepoints + nd.Protein.GetPoints(st)
		negative = nd.Energy.GetPoints(st) + nd.Suger.GetPoints(st) + nd.SaturatedFattyAcids.GetPoints(st) + nd.Sodium.GetPoints(st)

		if st == Cheese {
			value = negative - positive
		} else {
			if negative >= 11 && fruitpoints < 5 {
				value = negative - positive - fruitpoints
			} else {
				value = negative - positive
			}
		}
	}

	return NutritionalScore{
		Value:     value,
		Positive:  positive,
		Negative:  negative,
		ScoreType: st,
	}

}
func (ns NutritionalScore) GetNutriScore() string {
	if ns.ScoreType == Food {
		return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{18, 10, 2, -1})]
	}
	if ns.ScoreType == Water {
		return scoreToLetter[0]
	}
	return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{9, 5, 1, -2})]
}

func getPointsFromRange(v float64, steps []float64) int {
	lenSteps := len(steps)

	for i, val := range steps {

		if v > val {

			return lenSteps - i
		}
	}
	return 0
}
