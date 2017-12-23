package main

import (
	"bufio"
	"bytes"
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

	rawInput = bytes.TrimSpace(rawInput)
	input := string(rawInput)

	lengths := []int{}

	for _, inputLength := range strings.Split(input, ",") {
		length, _ := strconv.Atoi(inputLength)
		lengths = append(lengths, length)
	}

	// Part 1
	list := NewList(0, 255)
	fmt.Println(Hash(list, lengths, 1))

	// Part 2
	list = NewList(0, 255)
	newLengths := []int{}
	for _, b := range rawInput {
		newLengths = append(newLengths, int(b))
	}
	newLengths = append(newLengths, []int{17, 31, 73, 47, 23}...)
	sparseHash := Hash(list, newLengths, 64)

	hash := ""
	for block := 0; block < 16; block++ {
		elem := sparseHash[block*16]
		for i := (block * 16) + 1; i < (block*16)+16; i++ {
			elem = elem ^ sparseHash[i]
		}
		hash = fmt.Sprintf("%s%x", hash, elem)
	}

	fmt.Println(hash)
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
func Hash(digits, lengths []int, runs int) []int {
	skip := 0
	pos := 0

	for x := 0; x < runs; x++ {
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
	}

	return digits
}
