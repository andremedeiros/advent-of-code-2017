package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	simpleValidPasswords := 0
	anagramValidPasswords := 0

	for scanner.Scan() {
		if SimpleCheckValidPassword(scanner.Text()) {
			simpleValidPasswords++
		}

		if AnagramCheckValidPassword(scanner.Text()) {
			anagramValidPasswords++
		}
	}

	fmt.Printf("Simple valid passwords: %d\n", simpleValidPasswords)
	fmt.Printf("Anagram valid passwords: %d\n", anagramValidPasswords)
}

// SimpleCheckValidPassword checks whether a password has duplicate words and
// returns true if all the words are unique
func SimpleCheckValidPassword(password string) bool {
	usedWords := make(map[string]bool)
	words := strings.Split(password, " ")

	for _, word := range words {
		if _, ok := usedWords[word]; ok {
			return false
		}

		usedWords[word] = true
	}

	return true
}

type sortedByteSlice []byte

func (s sortedByteSlice) Len() int {
	return len(s)
}

func (s sortedByteSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortedByteSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// SortString sorts the characters in a string
func SortString(s string) string {
	b := sortedByteSlice(s)
	sort.Sort(b)
	return string(b)
}

// AnagramCheckValidPassword checks whether any of the words are the same, but also if any of the words
// are anagrams of one another
func AnagramCheckValidPassword(password string) bool {
	usedWords := make(map[string]bool)
	words := strings.Split(password, " ")

	for _, word := range words {
		sortedWord := SortString(word)
		if _, ok := usedWords[sortedWord]; ok {
			return false
		}

		usedWords[sortedWord] = true
	}

	return true
}
