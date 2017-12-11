package main

import (
	"testing"
)

func TestSimplePasswordValidation(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{"aa bb cc dd ee", true},
		{"aa bb cc dd aa", false},
		{"aa bb cc dd aaa", true},
	}

	for _, test := range tests {
		t.Run(test.password, func(t *testing.T) {
			got := SimpleCheckValidPassword(test.password)
			if got != test.expected {
				t.Errorf("Expected %t, got %t", test.expected, got)
			}
		})
	}
}

func TestAnagramPasswordValidation(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{"aa bb cc dd ee", true},
		{"abc bb cc dd cba", false},
		{"aa bb cc dd aaa", true},
	}

	for _, test := range tests {
		t.Run(test.password, func(t *testing.T) {
			got := AnagramCheckValidPassword(test.password)
			if got != test.expected {
				t.Errorf("Expected %t, got %t", test.expected, got)
			}
		})
	}
}
