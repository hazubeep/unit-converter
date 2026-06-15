package converter

import (
	"math"
	"testing"
)

func TestConvertLength(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		from, to LengthUnit
		expected float64
	}{
		{"meter to kilometer", 1000, Meter, Kilometer, 1},
		{"kilometer to meter", 1, Kilometer, Meter, 1000},
		{"meter to centimeter", 1, Meter, Centimeter, 100},
		{"centimeter to meter", 100, Centimeter, Meter, 1},

		{"meter to meter", 100, Meter, Meter, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertLength(tt.value, tt.from, tt.to)
			if math.Abs(result-tt.expected) > 0.001 {
				t.Errorf("ConvertLength(%.2f, %s, %s) = %.4f, want %.4f",
					tt.value, tt.from, tt.to, result, tt.expected)
			}
		})
	}
}
