
//line parser.go.rl:1
package main

import (
  "errors"
  "strconv"
)


//line parser.go:12
var _programs_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 2, 3, 4, 
}

var _programs_key_offsets []byte = []byte{
	0, 0, 2, 5, 6, 8, 11, 12, 
	13, 14, 16, 18, 
}

var _programs_trans_keys []byte = []byte{
	97, 122, 32, 97, 122, 40, 48, 57, 
	41, 48, 57, 45, 62, 32, 97, 122, 
	10, 32, 10, 44, 97, 122, 
}

var _programs_single_lengths []byte = []byte{
	0, 0, 1, 1, 0, 1, 1, 1, 
	1, 0, 2, 2, 
}

var _programs_range_lengths []byte = []byte{
	0, 1, 1, 0, 1, 1, 0, 0, 
	0, 1, 0, 1, 
}

var _programs_index_offsets []byte = []byte{
	0, 0, 2, 5, 7, 9, 12, 14, 
	16, 18, 20, 23, 
}

var _programs_trans_targs []byte = []byte{
	2, 0, 3, 2, 0, 4, 0, 5, 
	0, 10, 5, 0, 7, 0, 8, 0, 
	9, 0, 11, 0, 1, 6, 0, 1, 
	8, 11, 0, 
}

var _programs_trans_actions []byte = []byte{
	1, 0, 3, 0, 0, 0, 0, 1, 
	0, 5, 0, 0, 0, 0, 0, 0, 
	0, 0, 1, 0, 9, 0, 0, 11, 
	7, 0, 0, 
}

var _programs_eof_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 9, 11, 
}

const programs_start int = 1
const programs_first_final int = 10
const programs_error int = 0

const programs_en_main int = 1


//line parser.go.rl:11


// Parse parses a list of programs and returns the list of node hints
func Parse(data string) ([]NodeHint, error) {
  hints := []NodeHint{}

  cs, p, pe, eof := 0, 0, len(data), len(data)
  mark := 0

  currentHint := NodeHint{}

  
//line parser.go:83
	{
	cs = programs_start
	}

//line parser.go:88
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
//line parser.go.rl:23
 mark = p 
		case 1:
//line parser.go.rl:26
 currentHint.Name = data[mark:p] 
		case 2:
//line parser.go.rl:27
 currentHint.Weight, _ = strconv.Atoi(data[mark:p]) 
		case 3:
//line parser.go.rl:28
 currentHint.Children = append(currentHint.Children, data[mark:p]) 
		case 4:
//line parser.go.rl:29

      hints = append(hints, currentHint)
      currentHint = NodeHint{}
    
//line parser.go:184
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
			case 3:
//line parser.go.rl:28
 currentHint.Children = append(currentHint.Children, data[mark:p]) 
			case 4:
//line parser.go.rl:29

      hints = append(hints, currentHint)
      currentHint = NodeHint{}
    
//line parser.go:212
			}
		}
	}

	_out: {}
	}

//line parser.go.rl:45


  if eof != p {
    return []NodeHint{}, errors.New("parse error")
  }

  return hints, nil
}

