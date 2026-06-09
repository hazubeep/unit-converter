package main

type WeightUnit string

const (
	Kilogram  WeightUnit = "Kilogram"
	Hectogram WeightUnit = "Hectogram"
	Decagram  WeightUnit = "Decagram"
	Gram      WeightUnit = "Gram"
	Decigram  WeightUnit = "Decigram"
	Centigram WeightUnit = "Centigram"
	Milligram WeightUnit = "Milligram"
)

var WeightFactors = map[WeightUnit]float64{
	Kilogram:  1000,
	Hectogram: 100,
	Decagram:  10,
	Gram:      1,
	Decigram:  0.1,
	Centigram: 0.01,
	Milligram: 0.001,
}

func ConvertWeight(value float64, from, to WeightUnit) float64 {
	base := value * WeightFactors[from]
	return base / WeightFactors[to]
}
