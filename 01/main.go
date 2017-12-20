package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	digits := bytes.Runes([]byte(strings.TrimSpace(scanner.Text())))

	firstCaptcha := solveFirstCaptcha(digits)
	fmt.Printf("First Captcha: %d\n", firstCaptcha)
	secondCaptcha := solveSecondCaptcha(digits)
	fmt.Printf("Second Captcha: %d\n", secondCaptcha)
}

func solveCaptcha(digits []rune, f func(int) int) (captcha rune) {
	for i, c := range digits {
		if next := digits[f(i)%len(digits)]; c == next {
			captcha += c - '0'
		}
	}
	return
}

func solveFirstCaptcha(digits []rune) rune {
	return solveCaptcha(digits, func(i int) int { return i + 1 })
}

func solveSecondCaptcha(digits []rune) rune {
	return solveCaptcha(digits, func(i int) int { return i + len(digits)/2 })
}
