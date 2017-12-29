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
	layers, depths := ParseDepths(ranges)

	if layers != 7 {
		t.Errorf("Expected 7 layers, got %d\n", layers)
	}

	depthsMap := map[int]int{0: 3, 1: 2, 4: 4, 6: 4}

	if !reflect.DeepEqual(depthsMap, depths) {
		t.Errorf("Expected depths to be %+v, got %+v\n", depthsMap, depths)
	}
}

func TestCalculateSeverity(t *testing.T) {
	ranges, depths := ParseDepths(ranges)
	severity := CalculateSeverity(ranges, depths, CalculateSeverityOptions{})

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

func TestMinDelay(t *testing.T) {
	layers, depths := ParseDepths(ranges)
	drawScannerMovements(depths, 20)
	minDelay := MinDelay(layers, depths)

	if minDelay != 10 {
		t.Errorf("Expected delay to be 10, got %d\n", minDelay)
	}
}
