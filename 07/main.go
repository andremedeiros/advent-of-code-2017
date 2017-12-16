package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	b, _ := ioutil.ReadAll(reader)
	input := strings.TrimSpace(string(b))

	hints, err := Parse(input)
	if err != nil {
		panic(err)
	}

	program := SortNodes(hints)

	fmt.Printf("Root: %s\n", program.Name)
}

// Program holds the program's name and its list of children
type Program struct {
	Name     string
	Weight   int
	Children []*Program
	Parent   *Program
}

// NodeHint holds some hints about a program -- its parent ID and a list of siblings
type NodeHint struct {
	Name     string
	Weight   int
	Children []string
}

// SortNodes gets an array of NodeHint and returns the structured program tree
func SortNodes(hints []NodeHint) *Program {
	// Do a first pass where we put the ones with children inside their parents
	lookup := make(map[string]*Program)

	// Convert hints to programs
	for _, hint := range hints {
		lookup[hint.Name] = &Program{
			Name:   hint.Name,
			Weight: hint.Weight,
		}
	}

	// Populate children and not children
	for _, hint := range hints {
		if len(hint.Children) > 0 {
			parent, _ := lookup[hint.Name]

			for _, childName := range hint.Children {
				child, _ := lookup[childName]
				child.Parent = parent
				parent.Children = append(parent.Children, child)
			}
		}
	}

	for _, program := range lookup {
		if program.Parent == nil {
			return program
		}
	}

	return nil
}
