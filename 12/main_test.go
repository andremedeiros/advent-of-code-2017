package main

import (
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
	g, err := Parse(programs)

	if err != nil {
		t.Errorf("Got error: %+v", err)
	} else {
		edges := []struct {
			a, b int
		}{
			{2, 0}, {2, 3}, {2, 4},
			{3, 2}, {3, 4}, {4, 2},
			{4, 3}, {4, 6}, {5, 6},
			{6, 4}, {6, 5},
		}

		for _, edge := range edges {
			if !g.Edge(edge.a, edge.b) {
				t.Errorf("Expected %+v to be connected but wasnt\n", edge)
			}
		}
	}
}
