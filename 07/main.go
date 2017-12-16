package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
}

// Program holds the program's name and its list of children
type Program struct {
	Name     string
	Children []string
}

// NodeHint holds some hints about a program -- its parent ID and a list of siblings
type NodeHint struct {
	Name     string
	ParentID int
	Children []string
}
