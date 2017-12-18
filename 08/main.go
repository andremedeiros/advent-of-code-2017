package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	b, _ := ioutil.ReadAll(reader)
	input := strings.TrimSpace(string(b))

	vm, err := Parse(input)
	if err != nil {
		panic(err)
	}

	vm.Execute()
	_, maxValue := vm.MaxRegister()
	fmt.Printf("Max value (end): %d\n", maxValue)
	fmt.Printf("Max value (all): %d\n", vm.MaxValue)
}

// OperationType is the enum that specifies what operation we run on the register
type OperationType int

// These are the possible operations
const (
	Increment OperationType = iota
	Decrement
)

// ComparisonType is the comparison that we do on the subject register
type ComparisonType int

// These are the possible comparisons
const (
	Equal ComparisonType = iota
	NotEqual
	GreaterThan
	LowerThan
	GreaterOrEqualThan
	LowerOrEqualThan
)

// Instruction holds the details of the instructions
type Instruction struct {
	Register           string
	Operation          OperationType
	Amount             int
	ComparisonRegister string
	Comparison         ComparisonType
	ComparisonValue    int
}

// VirtualMachine contains the state of a run
type VirtualMachine struct {
	Instructions []Instruction
	Registers    map[string]int
	MaxValue     int
	maxValueSet  bool
}

// Execute runs through the instructions and executes them
func (vm *VirtualMachine) Execute() {
	vm.Registers = make(map[string]int)

	for _, instruction := range vm.Instructions {
		registerValue := vm.fetchRegister(instruction.Register)
		comparisonRegisterValue := vm.fetchRegister(instruction.ComparisonRegister)

		execute := false
		switch instruction.Comparison {
		case Equal:
			execute = (comparisonRegisterValue == instruction.ComparisonValue)
		case NotEqual:
			execute = (comparisonRegisterValue != instruction.ComparisonValue)
		case GreaterThan:
			execute = (comparisonRegisterValue > instruction.ComparisonValue)
		case LowerThan:
			execute = (comparisonRegisterValue < instruction.ComparisonValue)
		case GreaterOrEqualThan:
			execute = (comparisonRegisterValue >= instruction.ComparisonValue)
		case LowerOrEqualThan:
			execute = (comparisonRegisterValue <= instruction.ComparisonValue)
		}

		if !execute {
			continue
		}

		newValue := 0
		switch instruction.Operation {
		case Increment:
			newValue = registerValue + instruction.Amount
		case Decrement:
			newValue = registerValue - instruction.Amount
		}

		vm.Registers[instruction.Register] = newValue

		if !vm.maxValueSet || newValue > vm.MaxValue {
			vm.MaxValue = newValue
			vm.maxValueSet = true
		}
	}
}

// MaxRegister returns the register and value with the largest value
func (vm *VirtualMachine) MaxRegister() (string, int) {
	maxValue, maxRegister := 0, ""
	for register, value := range vm.Registers {
		if maxRegister == "" || maxValue < value {
			maxValue = value
			maxRegister = register
		}
	}

	return maxRegister, maxValue
}

func (vm *VirtualMachine) fetchRegister(register string) int {
	registerValue := 0

	if value, ok := vm.Registers[register]; ok {
		registerValue = value
	}

	return registerValue
}
