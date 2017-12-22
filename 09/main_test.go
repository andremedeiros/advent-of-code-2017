package main

import (
	"testing"
)

func TestScore(t *testing.T) {
	tests := []struct {
		blocks []byte
		score  int
	}{
		{[]byte("{}"), 1},
		{[]byte("{{{}}}"), 6},
		{[]byte("{{},{}}"), 5},
		{[]byte("{{{},{},{{}}}}"), 16},
		{[]byte("{<a>,<a>,<a>,<a>}"), 1},
		{[]byte("{{<ab>},{<ab>},{<ab>},{<ab>}}"), 9},
		{[]byte("{{<!!>},{<!!>},{<!!>},{<!!>}}"), 9},
		{[]byte("{{<a!>},{<a!>},{<a!>},{<ab>}}"), 3},
	}

	for _, test := range tests {
		t.Run(string(test.blocks), func(t *testing.T) {
			score, _ := ProcessStream(test.blocks)

			if score != test.score {
				t.Errorf("Expected score of %d, got %d\n", test.score, score)
			}
		})
	}
}

func TestRemoveGarbage(t *testing.T) {
	tests := []struct {
		blocks  []byte
		removed int
	}{
		{[]byte("<>"), 0},
		{[]byte("<random characters>"), 17},
		{[]byte("<<<<>"), 3},
		{[]byte("<{!>}>"), 2},
		{[]byte("<!!>"), 0},
		{[]byte("<!!!>>"), 0},
		{[]byte("<{o\"i!a,<{i<a>"), 10},
	}

	for _, test := range tests {
		t.Run(string(test.blocks), func(t *testing.T) {
			_, removed := ProcessStream(test.blocks)

			if removed != test.removed {
				t.Errorf("Expected removed to be %d, got %d\n", test.removed, removed)
			}
		})
	}
}
