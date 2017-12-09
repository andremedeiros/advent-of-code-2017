package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	simpleValidPasswords := 0

	for scanner.Scan() {
		if SimpleCheckValidPassword(scanner.Text()) {
			simpleValidPasswords++
		}
	}

	fmt.Printf("Simple valid passwords: %d\n", simpleValidPasswords)
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
