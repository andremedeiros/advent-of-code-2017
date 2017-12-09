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
		if record, err := reader.Read(); err == io.EOF {
			break
		} else {
			row := make([]int, len(record))

			for i := 0; i < len(record); i++ {
				row[i], _ = strconv.Atoi(record[i])
			}

			spreadsheet = append(spreadsheet, row)
		}
	}

	return spreadsheet
}

// Checksum1 calculates the first checksum
func (s Spreadsheet) Checksum1() int {
	checksum := 0

	for i := 0; i < len(s); i++ {
		row := s[i]

		min, max := row.MinMax()
		checksum += max - min
	}

	return checksum
}

// Checksum2 calculates the second checksum
func (s Spreadsheet) Checksum2() int {
	checksum := 0

	for i := 0; i < len(s); i++ {
		row := s[i]

		for j := 0; j < len(row)-1; j++ {
			for k := j + 1; k < len(row); k++ {
				// TODO: I hate this, we should immediately get the min and max out of this thing. Maybe sort the row before iterating?
				a := row[j]
				b := row[k]

				if a > b && a%b == 0 {
					checksum += a / b
				} else if b > a && b%a == 0 {
					checksum += b / a
				}
			}
		}
	}

	return checksum
}

// MinMax returns the minimum and maximum values of a row
func (r Row) MinMax() (int, int) {
	if len(r) == 0 {
		return 0, 0
	} else if len(r) == 1 {
		return r[0], r[0]
	} else {
		min := r[0]
		max := r[0]

		for i := 1; i < len(r); i++ {
			if r[i] > max {
				max = r[i]
			} else if r[i] < min {
				min = r[i]
			}
		}

		return min, max
	}
}
