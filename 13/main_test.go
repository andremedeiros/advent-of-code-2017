package main

import (
	"fmt"
	"reflect"
	"testing"
)

const ranges = `0: 3
1: 2
4: 4
6: 4`

func TestParseDepths(t *testing.T) {
	ranges, depths := ParseDepths(ranges)

	if ranges != 7 {
		t.Errorf("Expected 6 ranges, got %d\n", ranges)
	}

	depthsMap := map[int]int{0: 3, 1: 2, 4: 4, 6: 4}

	if !reflect.DeepEqual(depthsMap, depths) {
		t.Errorf("Expected depths to be %+v, got %+v\n", depthsMap, depths)
	}
}

func TestCalculateSeverity(t *testing.T) {
	ranges, depths := ParseDepths(ranges)
	severity := CalculateSeverity(ranges, depths)

	if severity != 24 {
		t.Errorf("Expected severity to be 24, got %d\n", severity)
	}
}

func TestScannerPosition(t *testing.T) {
	tests := []struct {
		depth, time, expected int
	}{
		{7, 0, 1},
		{7, 6, 7},
		{7, 8, 5},
		{7, 12, 1},
		{7, 16, 5},
	}

	for _, test := range tests {
		got := scannerPosition(test.depth, test.time)

		t.Run(fmt.Sprintf("depth: %d, time: %d", test.depth, test.time), func(t *testing.T) {
			if got != test.expected {
				t.Errorf("Expected position to be %d, got %d\n", test.expected, got)
			}
		})
	}
}
