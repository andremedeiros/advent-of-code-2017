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
	fmt.Printf("New weight: %d\n", program.Balance())
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

// IsBalanced checks whether a program is balanced or not
func (p *Program) IsBalanced() bool {
	weights := make(map[int]bool)

	for _, child := range p.Children {
		weights[child.TotalWeight()] = true
	}

	return len(weights) == 1
}

// TotalWeight recursively calculates total weight of a program's children
func (p *Program) TotalWeight() int {
	sum := p.Weight

	for _, child := range p.Children {
		sum += child.TotalWeight()
	}

	return sum
}

// Balance will find the program that needs balancing and balances it. It works under the assumption that the root program is not balanced.
func (p *Program) Balance() int {
	program := p

	for {
		previousProgram := program

		for _, child := range program.Children {
			if !child.IsBalanced() {
				program = child
				break
			}
		}

		if previousProgram == program {
			break
		}
	}

	// Divide weights by groups
	weightGroups := make(map[int][]*Program)

	fmt.Println(program.Name)

	for _, child := range program.Children {
		totalWeight := child.TotalWeight()
		if _, ok := weightGroups[totalWeight]; ok {
			weightGroups[totalWeight] = append(weightGroups[totalWeight], child)
		} else {
			weightGroups[totalWeight] = []*Program{child}
		}
	}

	badChild := p
	correctTotalWeight := 0

	for weight, group := range weightGroups {
		if len(group) > 1 {
			correctTotalWeight = weight
		} else {
			badChild = group[0]
		}
	}

	fmt.Println(badChild.Name, badChild.Weight, badChild.TotalWeight(), correctTotalWeight)
	diff := badChild.TotalWeight() - correctTotalWeight
	if diff < 0 {
		diff = diff * -1
	}

	return badChild.Weight - diff
}
