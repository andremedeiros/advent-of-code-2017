package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// ProgramHint holds the raw input from the program list
type ProgramHint struct {
	ID       int
	Children []int
}

// IntSlice holds a slice of ints
type IntSlice []int

// ProgramConnections stores connections between programs
type ProgramConnections map[int]IntSlice

func main() {
	reader := bufio.NewReader(os.Stdin)
	b, err := ioutil.ReadAll(reader)

	if err != nil {
		panic(err)
	}

	input := string(b)
	hints, err := Parse(input)

	if err != nil {
		panic(err)
	}

	lookup := SortHints(hints)
	fmt.Println(lookup.CountConnections(0))
}

// SortHints sorts through program hints and builds the tree accordingly
func SortHints(hints []ProgramHint) *ProgramConnections {
	lookup := &ProgramConnections{}

	for _, hint := range hints {
		for _, child := range hint.Children {
			lookup.Connect(hint.ID, child)
			lookup.Connect(child, hint.ID)
		}
	}

	return lookup
}

// Connect connects 2 programs to each other
func (pc *ProgramConnections) Connect(x, y int) {
	if slice, ok := (*pc)[x]; ok {
		if !slice.Contain(y) {
			(*pc)[x] = append((*pc)[x], y)
		}
	} else {
		(*pc)[x] = []int{y}
	}
}

// CountConnections counts how many programs are connected to a specific one
func (pc *ProgramConnections) CountConnections(id int) int {
	connections := 0
	seen := IntSlice{}
	todo := IntSlice{id}
	seenNew := true

	for {
		seenNew = false

		for _, id := range todo {
			if seen.Contain(id) {
				continue
			}

			seenNew = true

			for _, program := range (*pc)[id] {
				if todo.Contain(program) {
					continue
				}

				todo = append(todo, program)
			}

			seen = append(seen, id)
			connections++
		}

		if !seenNew {
			break
		}
	}

	return connections
}

// CountGroups counts how many distinct groups the connection set has
func (pc *ProgramConnections) CountGroups() int {
	return 0
}

// Contain returns true if the slice contains the specified element
func (slice IntSlice) Contain(x int) bool {
	for _, val := range slice {
		if val == x {
			return true
		}
	}

	return false
}
