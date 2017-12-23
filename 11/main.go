package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := ioutil.ReadAll(reader)
	input = bytes.TrimSpace(input)

	steps := strings.Split(string(input), ",")
	fmt.Println(CalculateSteps(steps))
}

// CalculateSteps takes in a list of steps and calculates the ideal amount of steps
// to get to the destination.
func CalculateSteps(steps []string) (int, int) {
	x := 0
	y := 0
	z := 0

	currentDistance, maxDistance := 0, 0

	for _, step := range steps {
		switch step {
		case "n":
			y++
			z--
		case "s":
			y--
			z++
		case "ne":
			x++
			z--
		case "nw":
			x--
			y++
		case "se":
			y--
			x++
		case "sw":
			x--
			z++
		}

		currentDistance = CalculateDistance(x, y, z)
		if currentDistance > maxDistance {
			maxDistance = currentDistance
		}
	}

	return currentDistance, maxDistance
}

// CalculateDistance takes 3 coordinates and calculates the distance
func CalculateDistance(x, y, z int) int {
	if x < 0 {
		x *= -1
	}

	if y < 0 {
		y *= -1
	}

	if z < 0 {
		z *= -1
	}

	return (x + y + z) / 2
}
