package main

import (
	"reflect"
	"testing"
)

func TestHash(t *testing.T) {
	digits := NewList(0, 4)
	hashed := Hash(digits, []int{3, 4, 1, 5})

	wanted := []int{3, 4, 2, 1, 0}
	if !reflect.DeepEqual(wanted, hashed) {
		t.Errorf("Expected list to be %+v, got %+v\n", wanted, hashed)
	}
}
