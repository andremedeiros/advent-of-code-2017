package main

import ()

// Point represents a cartesian coordinate
type Point struct {
	X, Y int
}

func main() {
}

// MatrixCoordinates returns the cartesian coordinates of a specific value
// in a spiral matrix
func MatrixCoordinates(position int) Point {
	return Point{0, 0}
}

// ManhattanDistance calculates sum of the absolute difference of two points'
// cartesian coordinates
func ManhattanDistance(p1, p2 Point) int {
	xDiff := p1.X - p2.X
	if xDiff < 0 {
		xDiff = xDiff * -1
	}

	yDiff := p1.Y - p2.Y
	if yDiff < 0 {
		yDiff = yDiff * -1
	}

	return xDiff + yDiff
}
