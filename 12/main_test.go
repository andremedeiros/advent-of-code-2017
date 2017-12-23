package main

import (
	"reflect"
	"testing"
)

const programs = `0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5`

func TestParse(t *testing.T) {
	hints, err := Parse(programs)

	if err != nil {
		t.Errorf("Got error: %+v", err)
	} else {
		expectedHints := []ProgramHint{
			ProgramHint{0, []int{2}},
			ProgramHint{1, []int{1}},
			ProgramHint{2, []int{0, 3, 4}},
			ProgramHint{3, []int{2, 4}},
			ProgramHint{4, []int{2, 3, 6}},
			ProgramHint{5, []int{6}},
			ProgramHint{6, []int{4, 5}},
		}

		if !reflect.DeepEqual(expectedHints, hints) {
			t.Errorf("Expected: %+v\nGot: %+v\n", expectedHints, hints)
		}
	}
}

func TestCounts(t *testing.T) {
	hints, _ := Parse(programs)
	lookup := SortHints(hints)

	tests := []struct {
		name        string
		id          int
		connections int
	}{
		{"self contained group", 1, 1},
		{"the other group", 6, 6},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			connections := lookup.CountConnections(test.id)

			if connections != test.connections {
				t.Errorf("Expected %d connections, got %d\n", test.connections, connections)
			}
		})
	}

	if groups := lookup.CountGroups(); groups != 2 {
		t.Errorf("Expected 2 groups, got %d\n", groups)
	}
}
