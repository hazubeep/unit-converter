package converter

import (
	"math"
	"testing"
)

func TestConvertTemperature(t *testing.T) {
	tests := []struct {
		value    float64
		from, to TemperatureUnit
		expected float64
	}{
		{100, Celsius, Fahrenheit, 212},
		{32, Fahrenheit, Celsius, 0},
		{0, Celsius, Kelvin, 273.15},
		{300, Kelvin, Fahrenheit, 80.33},
	}

	for _, tt := range tests {
		result := ConvertTemperature(tt.value, tt.from, tt.to)
		if math.Abs(result-tt.expected) > 0.01 {
			t.Errorf("got %.2f, want %.2f", result, tt.expected)
		}
	}

}
