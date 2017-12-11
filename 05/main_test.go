package main

import (
	"testing"
)

func TestStepsToExit(t *testing.T) {
	tests := []struct {
		name     string
		jumps    []int
		expected int
	}{
		{"simple", []int{0, 3, 0, 1, -3}, 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := StepsToExit(test.jumps)

			if got != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, got)
			}
		})
	}
}

func TestWeirdStepsToExit(t *testing.T) {
	tests := []struct {
		name     string
		jumps    []int
		expected int
	}{
		{"simple", []int{0, 3, 0, 1, -3}, 10},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := WeirdStepsToExit(test.jumps)

			if got != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, got)
			}
		})
	}
}
