package main

import (
	"testing"
)

func TestCalculateSteps(t *testing.T) {
	tests := []struct {
		name  string
		moves []string
		steps int
	}{
		{"3 steps away", []string{"ne", "ne", "ne"}, 3},
		{"back where we started", []string{"ne", "ne", "sw", "sw"}, 0},
		{"2 steps away", []string{"ne", "ne", "s", "s"}, 2},
		{"3 steps away", []string{"se", "sw", "se", "sw", "sw"}, 3},
	}

	for _, test := range tests {

		t.Run(string(test.name), func(t *testing.T) {
			steps, _ := CalculateSteps(test.moves)
			if steps != test.steps {
				t.Errorf("Expected %d steps, got %d", test.steps, steps)
			}
		})
	}
}
