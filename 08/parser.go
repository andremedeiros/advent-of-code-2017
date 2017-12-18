
//line parser.go.rl:1
// -*-go-*-
package main

import (
  "errors"
  "fmt"
  "strconv"
)


//line parser.go:14
var _programs_actions []byte = []byte{
	0, 1, 0, 1, 2, 1, 3, 1, 5, 
	1, 6, 1, 7, 2, 0, 2, 2, 
	0, 3, 2, 0, 5, 2, 4, 1, 
	3, 0, 4, 1, 
}

var _programs_key_offsets []byte = []byte{
	0, 0, 5, 7, 8, 9, 12, 16, 
	17, 18, 19, 24, 28, 29, 32, 36, 
	41, 44, 45, 50, 54, 
}

var _programs_trans_keys []byte = []byte{
	32, 9, 13, 97, 122, 100, 105, 101, 
	99, 32, 9, 13, 32, 45, 48, 57, 
	105, 102, 32, 32, 9, 13, 97, 122, 
	33, 61, 60, 62, 61, 32, 9, 13, 
	32, 61, 9, 13, 32, 9, 13, 97, 
	122, 32, 48, 57, 110, 32, 9, 13, 
	97, 122, 10, 45, 48, 57, 10, 48, 
	57, 
}

var _programs_single_lengths []byte = []byte{
	0, 1, 2, 1, 1, 1, 2, 1, 
	1, 1, 1, 2, 1, 1, 2, 1, 
	1, 1, 1, 2, 1, 
}

var _programs_range_lengths []byte = []byte{
	0, 2, 0, 0, 0, 1, 1, 0, 
	0, 0, 2, 1, 0, 1, 1, 2, 
	1, 0, 2, 1, 1, 
}

var _programs_index_offsets []byte = []byte{
	0, 0, 4, 7, 9, 11, 14, 18, 
	20, 22, 24, 28, 32, 34, 37, 41, 
	45, 48, 50, 54, 58, 
}

var _programs_indicies []byte = []byte{
	0, 0, 2, 1, 3, 4, 1, 5, 
	1, 6, 1, 7, 7, 1, 8, 9, 
	9, 1, 10, 1, 11, 1, 12, 1, 
	13, 13, 14, 1, 15, 15, 16, 1, 
	17, 1, 18, 18, 1, 18, 17, 18, 
	1, 19, 19, 20, 1, 21, 22, 1, 
	5, 1, 23, 23, 24, 1, 25, 26, 
	26, 1, 27, 28, 1, 
}

var _programs_trans_targs []byte = []byte{
	2, 0, 18, 3, 17, 4, 5, 6, 
	7, 16, 8, 9, 10, 11, 15, 12, 
	14, 13, 19, 11, 15, 7, 16, 2, 
	18, 1, 20, 1, 20, 
}

var _programs_trans_actions []byte = []byte{
	19, 0, 1, 1, 1, 0, 0, 9, 
	13, 1, 0, 0, 0, 16, 1, 1, 
	1, 0, 11, 5, 0, 3, 0, 7, 
	0, 25, 1, 22, 0, 
}

var _programs_eof_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 25, 22, 
}

const programs_start int = 1
const programs_first_final int = 19
const programs_error int = 0

const programs_en_main int = 1


//line parser.go.rl:13


// Parse parses a list of programs and returns the list of node hints
func Parse(data string) (*VirtualMachine, error) {
  instructions := []Instruction{}

  cs, p, pe, eof := 0, 0, len(data), len(data)
  mark := 0

  currentInstruction := Instruction{}

  
//line parser.go:108
	{
	cs = programs_start
	}

//line parser.go:113
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	_keys = int(_programs_key_offsets[cs])
	_trans = int(_programs_index_offsets[cs])

	_klen = int(_programs_single_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case data[p] < _programs_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _programs_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_programs_range_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case data[p] < _programs_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _programs_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	_trans = int(_programs_indicies[_trans])
	cs = int(_programs_trans_targs[_trans])

	if _programs_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_programs_trans_actions[_trans])
	_nacts = uint(_programs_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _programs_actions[_acts-1] {
		case 0:
//line parser.go.rl:25
 mark = p 
		case 1:
//line parser.go.rl:28

      instructions = append(instructions, currentInstruction)
      currentInstruction = Instruction{}
    
		case 2:
//line parser.go.rl:33
 currentInstruction.Amount, _ = strconv.Atoi(data[mark:p]) 
		case 3:
//line parser.go.rl:34
 currentInstruction.ComparisonRegister = data[mark:p] 
		case 4:
//line parser.go.rl:35
 currentInstruction.ComparisonValue, _ = strconv.Atoi(data[mark:p]) 
		case 5:
//line parser.go.rl:36
 currentInstruction.Register = data[mark:p] 
		case 6:
//line parser.go.rl:38

      switch op := data[mark:p]; op {
        case "inc":
          currentInstruction.Operation = Increment
        case "dec":
          currentInstruction.Operation = Decrement
        default:
          return nil, fmt.Errorf("invalid operation found: %s", op)
      }
    
		case 7:
//line parser.go.rl:49

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
    
//line parser.go:245
		}
	}

_again:
	if cs == 0 {
		goto _out
	}
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		__acts := _programs_eof_actions[cs]
		__nacts := uint(_programs_actions[__acts]); __acts++
		for ; __nacts > 0; __nacts-- {
			__acts++
			switch _programs_actions[__acts-1] {
			case 0:
//line parser.go.rl:25
 mark = p 
			case 1:
//line parser.go.rl:28

      instructions = append(instructions, currentInstruction)
      currentInstruction = Instruction{}
    
			case 4:
//line parser.go.rl:35
 currentInstruction.ComparisonValue, _ = strconv.Atoi(data[mark:p]) 
//line parser.go:276
			}
		}
	}

	_out: {}
	}

//line parser.go.rl:86


  if eof != p {
    fmt.Println(currentInstruction)
    return nil, errors.New("parse error")
  }

  return &VirtualMachine{Instructions: instructions}, nil
}

