package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	puzzle := strings.TrimSpace(scanner.Text())
	// TODO: Try to avoid doing all this work ahead of time.
	digits := getIntSlice(puzzle)

	firstCaptcha := solveFirstCaptcha(digits)
	fmt.Printf("First Captcha: %d\n", firstCaptcha)
	secondCaptcha := solveSecondCaptcha(digits)
	fmt.Printf("Second Captcha: %d\n", secondCaptcha)
}

func getIntSlice(digits string) []int {
	ints := make([]int, len(digits))
	for i := 0; i < len(digits); i++ {
		ints[i] = int(digits[i] - byte('0'))
	}

	return ints
}

func solveFirstCaptcha(digits []int) int {
	captcha := 0

	for i := 0; i < len(digits); i++ {
		current := digits[i]
		next := digits[(i+1)%len(digits)]

		if current == next {
			captcha += current
		}
	}

	return captcha
}

func solveSecondCaptcha(digits []int) int {
	captcha := 0

	for i := 0; i < len(digits); i++ {
		current := digits[i]
		next := digits[((len(digits)/2)+i)%len(digits)]

		if current == next {
			captcha += current
		}
	}

	return captcha
}
