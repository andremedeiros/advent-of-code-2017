package main

import (
	"bytes"
	"testing"
)

func TestFirstCaptcha(t *testing.T) {
	tests := []struct {
		name, puzzle string
		expected     rune
	}{
		{"simple match", "1122", 3},
		{"multiple matches", "1111", 4},
		{"no matches", "1234", 0},
		{"circular match", "91212129", 9},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// digits := getIntSlice(test.puzzle)
			got := solveFirstCaptcha(bytes.Runes([]byte(test.puzzle)))
			if got != test.expected {
				t.Errorf("Expected: %d, got: %d\n", test.expected, got)
			}
		})
	}
}

func TestSecondCaptcha(t *testing.T) {
	tests := []struct {
		name, puzzle string
		expected     rune
	}{
		{"all matches", "1212", 6},
		{"no matches", "1221", 0},
		{"single match", "123425", 4},
		{"all matches (2)", "123123", 12},
		{"some matches", "12131415", 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := solveSecondCaptcha(bytes.Runes([]byte(test.puzzle)))
			if got != test.expected {
				t.Errorf("Expected: %d, got: %d\n", test.expected, got)
			}
		})
	}
}
