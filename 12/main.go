package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/yourbasic/graph"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	b, err := ioutil.ReadAll(reader)

	if err != nil {
		panic(err)
	}

	input := string(b)
	g, err := Parse(input)

	if err != nil {
		panic(err)
	}

	seen := 1 // The program talks to itself so we initialize at 1
	graph.BFS(g, 0, func(v, w int, c int64) {
		seen++
	})

	groups := graph.StrongComponents(g)

	fmt.Println("Programs on the same group as program 0:", seen)
	fmt.Println("Different groups:", len(groups))
}
