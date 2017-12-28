package main

import (
  "errors"
  "strconv"

  "github.com/yourbasic/graph"
)

%%{
  machine programs;
  write data;
}%%

// Parse parses a list of programs and returns the list of node hints
func Parse(data string) (graph.Graph, error) {
  cs, p, pe, eof := 0, 0, len(data), len(data)
  mark := 0

  graph := graph.New(len(hints))
  currentHint := NodeHint{}

  %%{
    action mark                        { mark = p }
    action debug                       { fmt.Println(data[mark:p]) }

    action register_program_name       { currentHint.Name = data[mark:p] }
    action register_weight             { currentHint.Weight, _ = strconv.Atoi(data[mark:p]) }
    action register_child_program_name { currentHint.Children = append(currentHint.Children, data[mark:p]) }
    action new_entry {
      hints = append(hints, currentHint)
      currentHint = NodeHint{}
    }

    program_name = ( lower )+ >mark;
    weight       = ( digit )+ >mark;

    entry            = ( program_name %register_program_name ' (' weight %register_weight ')' );
    entry_with_hints = ( entry ' -> ' ( program_name %register_child_program_name ', ' )* program_name %register_child_program_name );
    program_entry    = ( entry_with_hints | entry ) %new_entry;

    main := ( program_entry '\n' )* program_entry;

    write init;
    write exec;
  }%%

  if eof != p {
    return []NodeHint{}, errors.New("parse error")
  }




  return hints, nil
}

