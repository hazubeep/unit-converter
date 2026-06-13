package converter

type LengthUnit string

const (
	Kilometer  LengthUnit = "Kilometer"
	Meter      LengthUnit = "Meter"
	Centimeter LengthUnit = "Centimeter"
	Millimeter LengthUnit = "Millimeter"
)

var lengthFactors = map[LengthUnit]float64{
	Kilometer:  1000,
	Meter:      1,
	Centimeter: 0.01,
	Millimeter: 0.001,
}

func ConvertLength(value float64, from, to LengthUnit) float64 {
	base := value * lengthFactors[from]
	return base / lengthFactors[to]
}
