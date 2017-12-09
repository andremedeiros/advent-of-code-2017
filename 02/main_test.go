package main

import (
	"regexp"
	"strings"
	"testing"
)

func sanitizeSpreadsheet(contents string) string {
	re := regexp.MustCompile(" ")

	result := re.ReplaceAllString(contents, "\t")
	return strings.TrimSpace(result)
}

func TestChecksum1(t *testing.T) {
	puzzle := sanitizeSpreadsheet(`
5 1 9 5
7 5 3
2 4 6 8
	`)

	spreadsheet := NewSpreadsheet(puzzle)
	checksum := spreadsheet.Checksum1()

	if checksum != 18 {
		t.Errorf("Expected checksum to be 18, got %d\n", checksum)
	}
}

func TestChecksum2(t *testing.T) {
	puzzle := sanitizeSpreadsheet(`
5 9 2 8
9 4 7 3
3 8 6 5
	`)

	spreadsheet := NewSpreadsheet(puzzle)
	checksum := spreadsheet.Checksum2()

	if checksum != 9 {
		t.Errorf("Expected checksum to be 9, got %d\n", checksum)
	}
}
