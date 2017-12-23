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
	rawInput, err := ioutil.ReadAll(reader)

	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(rawInput))

	lengths := []int{}

	for _, inputLength := range strings.Split(input, ",") {
		length, _ := strconv.Atoi(inputLength)
		lengths = append(lengths, length)
	}

	list := NewList(0, 255)
	fmt.Println(Hash(list, lengths))
}

// NewList generates an int list starting at `start` and ending at `end`
func NewList(start, end int) []int {
	digits := []int{}

	for i := start; i <= end; i++ {
		digits = append(digits, i)
	}

	return digits
}

// Hash hashes the digit list with the lengths list provided
func Hash(digits, lengths []int) []int {
	skip := 0
	pos := 0

	for _, length := range lengths {
		start := pos
		end := pos + length

		sublist := []int{}
		for i := start; i < end; i++ {
			sublist = append(sublist, digits[i%len(digits)])
		}

		// Reverse the list
		for i, j := 0, len(sublist)-1; i < j; i, j = i+1, j-1 {
			sublist[i], sublist[j] = sublist[j], sublist[i]
		}

		// Put it back in place
		for i, item := range sublist {
			digits[(start+i)%len(digits)] = item
		}

		pos += skip + (end - start)
		skip++
	}

	return digits
}
