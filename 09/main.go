package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	blocks, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	score, removed := ProcessStream(blocks)
	fmt.Println("Score:", score)
	fmt.Println("Removed:", removed)
}

// ProcessStream processes an incoming stream and returns the amount of groups it contains
func ProcessStream(blocks []byte) (int, int) {
	ignoreNext := false
	insideGarbage := false

	depth := 0
	score := 0
	removed := 0

	for _, block := range blocks {
		if ignoreNext {
			ignoreNext = false
		} else {
			if block == '!' {
				ignoreNext = true
			} else if insideGarbage {
				if block != '>' {
					removed++
					continue
				}

				insideGarbage = false
			} else {
				if block == '<' {
					insideGarbage = true
				} else if block == '{' {
					depth++
				} else if block == '}' {
					score += depth
					depth--
				}
			}
		}
	}

	return score, removed
}
