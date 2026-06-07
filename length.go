package main

type Unit string

const (
	Millimeter Unit = "Milimeter"
	Centimeter Unit = "Centimeter"
	Meter      Unit = "Meter"
	Kilometer  Unit = "Kilometer"
)

var lengthFactors = map[Unit]float64{
	Millimeter: 0.001,
	Centimeter: 0.01,
	Meter:      1,
	Kilometer:  1000,
}

func ConvertLength(value float64, from, to Unit) float64 {
	base := value * lengthFactors[from]
	return base / lengthFactors[to]
}
