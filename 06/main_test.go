package main

import (
	"testing"
)

func TestCalculateRedistributions(t *testing.T) {
	tests := []struct {
		name                                 string
		banks                                []int
		expectedRedistributions, expectedAge int
	}{
		{"simple", []int{0, 2, 7, 0}, 5, 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotRedistributions, gotAge := CalculateRedistributions(test.banks)

			if gotRedistributions != test.expectedRedistributions {
				t.Errorf("Expected redistributions to be %d, got %d\n", test.expectedRedistributions, gotRedistributions)
			}

			if gotAge != test.expectedAge {
				t.Errorf("Expected age to be %d, got %d\n", test.expectedAge, gotAge)
			}
		})
	}
}
