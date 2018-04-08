package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Row contains an array of columns
type Row []int

// Spreadsheet carries an array of rows
type Spreadsheet []Row

func main() {
	reader := bufio.NewReader(os.Stdin)
	b, _ := ioutil.ReadAll(reader)

	spreadsheet := NewSpreadsheet(string(b))
	fmt.Printf("First Checksum: %d\n", spreadsheet.Checksum1())
	fmt.Printf("Second Checksum: %d\n", spreadsheet.Checksum2())
}

// NewSpreadsheet initializes a Spreadsheet from a string
func NewSpreadsheet(contents string) Spreadsheet {
	spreadsheet := make(Spreadsheet, 0)

	reader := csv.NewReader(strings.NewReader(contents))
	reader.Comma = '\t'

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		row := make([]int, 0, len(record))
		for _, r := range record {
			i, _ := strconv.Atoi(r)
			row = append(row, i)
		}

		spreadsheet = append(spreadsheet, row)
	}

	return spreadsheet
}

// Checksum1 calculates the first checksum
func (s Spreadsheet) Checksum1() (checksum int) {
	for _, row := range s {
		min, max := row.MinMax()
		checksum += max - min
	}
	return
}

// Checksum2 calculates the second checksum
func (s Spreadsheet) Checksum2() (checksum int) {
	for _, row := range s {
		for j, a := range row[:len(row)-1] {
			for _, b := range row[j+1:] {
				switch {
				case a > b && a%b == 0:
					checksum += a / b
				case b > a && b%a == 0:
					checksum += b / a
				}
			}
		}
	}
	return
}

// MinMax returns the minimum and maximum values of a row
func (r Row) MinMax() (min, max int) {
	if len(r) == 0 {
		return
	}

	min, max = r[0], r[0]

	for _, x := range r[1:] {
		switch {
		case x > max:
			max = x
		case x < min:
			min = x
		}
	}

	return
}
