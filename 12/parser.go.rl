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

// Parse parses a list of programs and returns a graph
func Parse(data string) (*graph.Mutable, error) {
  graph := graph.New(2000)

  cs, p, pe, eof := 0, 0, len(data), len(data)
  mark := 0

  id := 0

  %%{
    action mark                        { mark = p }
    action debug                       { fmt.Println(data[mark:p]) }

    action register_program_id       { id, _ = strconv.Atoi(data[mark:p]) }
    action register_child_program_id {
      child, _ := strconv.Atoi(data[mark:p])
      graph.AddBoth(id, child)
    }

    program_id = ( digit )+ >mark;

    entry = (
      program_id %register_program_id ' <-> '
      ( program_id %register_child_program_id ', ' )*
      program_id %register_child_program_id
    );

    main := ( entry '\n' )* entry;

    write init;
    write exec;
  }%%

  if eof != p {
    return nil, errors.New("parse error")
  }

  return graph, nil
}

