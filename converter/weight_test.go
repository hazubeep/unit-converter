package converter

import (
	"math"
	"testing"
)

func TestConvertWeight(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		from, to WeightUnit
		expected float64
	}{
		{"kilogram to gram", 1, Kilogram, Gram, 1000},
		{"gram  to kilogram", 1000, Gram, Kilogram, 1},
		{"gram to milligram", 1, Gram, Milligram, 1000},
		{"milligram to gram", 1000, Milligram, Gram, 1},

		{"kilogram to milligram", 1, Kilogram, Milligram, 1000000},
		{"hectorgram  to decagram", 1, Hectogram, Decagram, 10},
		{"decigram to centigram", 1, Decigram, Centigram, 10},

		{"gram to gram", 100, Gram, Gram, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertWeight(tt.value, tt.from, tt.to)
			if math.Abs(result-tt.expected) > 0.001 {
				t.Errorf("ConvertWeight(%.2f, %s, %s) = %.4f, want %.4f",
					tt.value, tt.from, tt.to, result, tt.expected)
			}
		})
	}
}
