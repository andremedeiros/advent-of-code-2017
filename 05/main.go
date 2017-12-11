package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	jumps := make([]int, 0)

	for scanner.Scan() {
		jump, _ := strconv.Atoi(scanner.Text())
		jumps = append(jumps, jump)
	}

	fmt.Printf("Steps to exit: %d\n", StepsToExit(jumps))
	fmt.Printf("Weird steps to exit: %d\n", WeirdStepsToExit(jumps))
}

// StepsToExit calculates how many steps it takes to exit the maze
func StepsToExit(origJumps []int) int {
	jumps := make([]int, len(origJumps))
	copy(jumps, origJumps)

	currentOffset := 0
	steps := 0

	for {
		previousOffset := currentOffset
		currentOffset = currentOffset + jumps[previousOffset]
		jumps[previousOffset]++
		steps++

		if currentOffset < 0 || currentOffset >= len(jumps) {
			break
		}

	}

	return steps
}

// WeirdStepsToExit calculates how many steps it takes to exit the maze
func WeirdStepsToExit(origJumps []int) int {
	jumps := make([]int, len(origJumps))
	copy(jumps, origJumps)

	currentOffset := 0
	steps := 0

	for {
		previousOffset := currentOffset
		currentOffset = currentOffset + jumps[previousOffset]

		if jumps[previousOffset] >= 3 {
			jumps[previousOffset]--
		} else {
			jumps[previousOffset]++
		}

		steps++

		if currentOffset < 0 || currentOffset >= len(jumps) {
			break
		}

	}

	return steps
}
