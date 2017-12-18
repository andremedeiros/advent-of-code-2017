// -*-go-*-
package main

import (
  "errors"
  "fmt"
  "strconv"
)

%%{
  machine programs;
  write data;
}%%

// Parse parses a list of programs and returns the list of node hints
func Parse(data string) (*VirtualMachine, error) {
  instructions := []Instruction{}

  cs, p, pe, eof := 0, 0, len(data), len(data)
  mark := 0

  currentInstruction := Instruction{}

  %%{
    action mark  { mark = p }
    action debug { fmt.Println(data[mark:p]) }

    action new_instruction {
      instructions = append(instructions, currentInstruction)
      currentInstruction = Instruction{}
    }

    action register_amount              { currentInstruction.Amount, _ = strconv.Atoi(data[mark:p]) }
    action register_comparison_register { currentInstruction.ComparisonRegister = data[mark:p] }
    action register_comparison_value    { currentInstruction.ComparisonValue, _ = strconv.Atoi(data[mark:p]) }
    action register_register            { currentInstruction.Register = data[mark:p] }

    action register_operation {
      switch op := data[mark:p]; op {
        case "inc":
          currentInstruction.Operation = Increment
        case "dec":
          currentInstruction.Operation = Decrement
        default:
          return nil, fmt.Errorf("invalid operation found: %s", op)
      }
    }

    action register_comparison_type {
      switch comp := data[mark:p]; comp {
        case "==":
          currentInstruction.Comparison = Equal
        case "!=":
          currentInstruction.Comparison = NotEqual
        case ">":
          currentInstruction.Comparison = GreaterThan
        case "<":
          currentInstruction.Comparison = LowerThan
        case ">=":
          currentInstruction.Comparison = GreaterOrEqualThan
        case "<=":
          currentInstruction.Comparison = LowerOrEqualThan
        default:
          return nil, fmt.Errorf("invalid comparison found: %s", comp)
      }
    }

    register   = ( lower )* >mark;
    value      = ( '-'? ( digit )* ) >mark;
    operation  = ( 'inc' | 'dec' ) >mark;
    comparison = ( '>' | '<' | '>=' | '<=' | '==' | '!=' ) >mark;

    entry = (
      register %register_register            space
      operation %register_operation          space
      value %register_amount                 ' if '
      register %register_comparison_register space
      comparison %register_comparison_type   space
      value %register_comparison_value
    ) %new_instruction;

    main := ( entry '\n' )* entry;

    write init;
    write exec;
  }%%

  if eof != p {
    return nil, errors.New("parse error")
  }

  return &VirtualMachine{Instructions: instructions}, nil
}

