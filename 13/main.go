package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	raw, _ := ioutil.ReadAll(reader)
	input := strings.TrimSpace(string(raw))

	layers, depths := ParseDepths(input)
	fmt.Println("Severity:", CalculateSeverity(layers, depths, CalculateSeverityOptions{}))
	fmt.Println("Min delay:", MinDelay(layers, depths))
}

// CalculateSeverityOptions passes options to the CalculateSeverity method
type CalculateSeverityOptions struct {
	BailWhenCaught bool
	Delay          int
}

// ParseDepths returns depths for the layers that hve them and how many layers there are
func ParseDepths(input string) (int, map[int]int) {
	entries := strings.Split(input, "\n")
	depthMap := map[int]int{}
	maxLayer := 0

	for _, entry := range entries {
		parts := strings.Split(entry, ": ")

		layer, _ := strconv.Atoi(parts[0])
		depth, _ := strconv.Atoi(parts[1])

		if maxLayer < layer {
			maxLayer = layer
		}

		depthMap[layer] = depth
	}

	return maxLayer + 1, depthMap
}

// CalculateSeverity calculates the severity of a trip
func CalculateSeverity(layers int, depths map[int]int, options CalculateSeverityOptions) int {
	severity := 0
	clock := options.Delay

	for i := 0; i < layers; i++ {
		depth, ok := depths[i]

		if ok {
			// Depth found. Let's see if the scanner hits us.
			pos := scannerPosition(depth, clock)

			if pos == 1 {
				severity += (i * depth)
				if options.BailWhenCaught {
					return severity
				}
			}
		}

		clock++
	}

	return severity
}

// MinDelay calculates the minimum delay necessary to avoid detection
func MinDelay(layers int, depths map[int]int) int {
	for i := 1; ; i++ {
		severity := CalculateSeverity(layers, depths, CalculateSeverityOptions{BailWhenCaught: true, Delay: i})
		if severity == 0 {
			return i
		}
	}
}

func scannerPosition(depth, time int) int {
	pos := 1
	direction := 1

	for i := 0; i < time; i++ {
		if pos == depth && direction == 1 {
			direction = -1
		} else if pos == 1 && direction == -1 {
			direction = 1
		}
		pos += direction
	}

	return pos
}
