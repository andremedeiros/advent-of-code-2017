package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	b, _ := ioutil.ReadAll(reader)
	input := strings.TrimSpace(string(b))

	inputs := strings.Split(input, "\t")
	banks := make([]int, len(inputs))

	for i, input := range inputs {
		banks[i], _ = strconv.Atoi(input)
	}

	redistributions, age := CalculateRedistributions(banks)
	fmt.Printf("Redistribution cycles: %d\n", redistributions)
	fmt.Printf("Age of previous cycle: %d\n", age)
}

// CalculateRedistributions takes a list of blocks in banks and returns how many reditributions it can do before
// the value ends up being the same
func CalculateRedistributions(origBanks []int) (int, int) {
	banks := make([]int, len(origBanks))
	copy(banks, origBanks)

	configurationsSeen := make([][]int, 0)

	for {
		// Redistribute
		maxIdx, remainder := 0, banks[0]
		for i := 1; i < len(banks); i++ {
			if banks[i] <= remainder {
				continue
			}

			maxIdx = i
			remainder = banks[i]
		}

		banks[maxIdx] = 0

		for i := maxIdx + 1; i <= maxIdx+remainder; i++ {
			banks[i%len(banks)]++
		}

		// Check if already seen and, if so, bail out
		for i := 0; i < len(configurationsSeen); i++ {
			if reflect.DeepEqual(banks, configurationsSeen[i]) {
				return len(configurationsSeen) + 1, len(configurationsSeen) - i
			}
		}

		// Haven't seen this one yet, so let's register it
		banksCopy := make([]int, len(banks))
		copy(banksCopy, banks)

		configurationsSeen = append(configurationsSeen, banksCopy)
	}
}
