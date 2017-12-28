
//line parser.go.rl:1
package main

import (
  "errors"
  "strconv"

	"github.com/yourbasic/graph"
)


//line parser.go:14
var _programs_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 
}

var _programs_key_offsets []byte = []byte{
	0, 0, 2, 5, 6, 7, 8, 9, 
	11, 
}

var _programs_trans_keys []byte = []byte{
	48, 57, 32, 48, 57, 60, 45, 62, 
	32, 48, 57, 10, 44, 48, 57, 
}

var _programs_single_lengths []byte = []byte{
	0, 0, 1, 1, 1, 1, 1, 0, 
	2, 
}

var _programs_range_lengths []byte = []byte{
	0, 1, 1, 0, 0, 0, 0, 1, 
	1, 
}

var _programs_index_offsets []byte = []byte{
	0, 0, 2, 5, 7, 9, 11, 13, 
	15, 
}

var _programs_trans_targs []byte = []byte{
	2, 0, 3, 2, 0, 4, 0, 5, 
	0, 6, 0, 7, 0, 8, 0, 1, 
	6, 8, 0, 
}

var _programs_trans_actions []byte = []byte{
	1, 0, 3, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 1, 0, 5, 
	5, 0, 0, 
}

var _programs_eof_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	5, 
}

const programs_start int = 1
const programs_first_final int = 8
const programs_error int = 0

const programs_en_main int = 1


//line parser.go.rl:13


// Parse parses a list of programs and returns a graph
func Parse(data string) (*graph.Mutable, error) {
  graph := graph.New(2000)

  cs, p, pe, eof := 0, 0, len(data), len(data)
  mark := 0

  id := 0

  
//line parser.go:81
	{
	cs = programs_start
	}

//line parser.go:86
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
//line parser.go.rl:25
 mark = p 
		case 1:
//line parser.go.rl:28
 id, _ = strconv.Atoi(data[mark:p]) 
		case 2:
//line parser.go.rl:29

      child, _ := strconv.Atoi(data[mark:p])
      graph.AddBoth(id, child)
    
//line parser.go:176
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
			case 2:
//line parser.go.rl:29

      child, _ := strconv.Atoi(data[mark:p])
      graph.AddBoth(id, child)
    
//line parser.go:201
			}
		}
	}

	_out: {}
	}

//line parser.go.rl:46


  if eof != p {
    return nil, errors.New("parse error")
  }

  return graph, nil
}

