package main

import (
	"reflect"
	"testing"
)

var (
	program = `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`
)

func TestParse(t *testing.T) {
	vm, err := Parse(program)

	if err != nil {
		t.Errorf("Got error: %+v\n", err)
	}

	if len(vm.Instructions) != 4 {
		t.Errorf("Expected 4 instructions, got %d\n", len(vm.Instructions))
	} else {
		expectedInstructions := []Instruction{
			Instruction{"b", Increment, 5, "a", GreaterThan, 1},
			Instruction{"a", Increment, 1, "b", LowerThan, 5},
			Instruction{"c", Decrement, -10, "a", GreaterOrEqualThan, 1},
			Instruction{"c", Increment, -20, "c", Equal, 10},
		}

		for i, expected := range expectedInstructions {
			got := vm.Instructions[i]

			if !reflect.DeepEqual(got, expected) {
				t.Errorf("Expected %+v, got %+v\n", expected, got)
			}
		}
	}
}

func TestExecute(t *testing.T) {
	vm, _ := Parse(program)
	vm.Execute()

	tests := []struct {
		register string
		value    int
	}{
		{"a", 1},
		{"c", -10},
	}

	for _, test := range tests {
		value, ok := vm.Registers[test.register]
		if !ok {
			t.Errorf("Expected register %+q to be set but it wasn't\n", test.register)
		} else {
			if value != test.value {
				t.Errorf("Expected value for register %+v to be %d, got %d", test.register, test.value, value)
			}
		}
	}
}
