package converter

type TemperatureUnit string

const (
	Celsius    TemperatureUnit = "Celsius"
	Fahrenheit TemperatureUnit = "Fahrenheit"
	Kelvin     TemperatureUnit = "Kelvin"
)

func toCelsius(value float64, from TemperatureUnit) float64 {
	switch from {
	case Celsius:
		return value
	case Fahrenheit:
		return (value - 32) * 5 / 9
	case Kelvin:
		return value - 273.15
	default:
		return value
	}
}

func fromCelsius(celsius float64, to TemperatureUnit) float64 {
	switch to {
	case Celsius:
		return celsius
	case Fahrenheit:
		return (celsius * 9 / 5) + 32
	case Kelvin:
		return celsius + 273.15
	default:
		return celsius
	}
}
func ConvertTemperature(value float64, from, to TemperatureUnit) float64 {
	celsius := toCelsius(value, from)
	return fromCelsius(celsius, to)
}
