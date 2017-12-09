package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMatrixCoordinates(t *testing.T) {
	tests := []struct {
		position int
		expected Point
	}{
		{1, Point{0, 0}},
		{12, Point{2, -1}},
		{23, Point{0, 2}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Matrix coordinates for %d", test.position), func(t *testing.T) {
			got := MatrixCoordinates(test.position)

			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("Expected %+v, got %+v\n", test.expected, got)
			}
		})
	}
}

func TestManhattanDistance(t *testing.T) {
	tests := []struct {
		from, to Point
		expected int
	}{
		{Point{6, 4}, Point{1, 2}, 7},
		{Point{1, 2}, Point{6, 4}, 7},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Manhattan Distance between %+v and %+v", test.from, test.to), func(t *testing.T) {
			got := ManhattanDistance(test.from, test.to)

			if got != test.expected {
				t.Errorf("Expected %d, got %d\n", test.expected, got)
			}
		})
	}
}
