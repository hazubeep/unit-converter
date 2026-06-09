package main

type Unit string

const (
	Kilometer  Unit = "Kilometer"
	Meter      Unit = "Meter"
	Centimeter Unit = "Centimeter"
	Millimeter Unit = "Milimeter"
)

var lengthFactors = map[Unit]float64{
	Kilometer:  1000,
	Meter:      1,
	Centimeter: 0.01,
	Millimeter: 0.001,
}

func ConvertLength(value float64, from, to Unit) float64 {
	base := value * lengthFactors[from]
	return base / lengthFactors[to]
}
