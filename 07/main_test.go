package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	programs := `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`

	hints, err := Parse(programs)

	if err != nil {
		t.Errorf("Got error: %+v\n", err)
	}

	if len(hints) != 13 {
		t.Errorf("Expected 13 programs, got %d\n", len(hints))
	}

	hint := hints[5]
	wanted := NodeHint{
		Name:     "fwft",
		ParentID: 72,
		Children: []string{"ktlj", "cntj", "xhth"},
	}

	if !reflect.DeepEqual(hint, wanted) {
		t.Errorf("Expected %+v, got %+v\n", wanted, hint)
	}
}
