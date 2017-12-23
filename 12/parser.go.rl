package main

import (
  "errors"
  "strconv"
)

%%{
  machine programs;
  write data;
}%%

// Parse parses a list of programs and returns the list of node hints
func Parse(data string) ([]ProgramHint, error) {
  hints := []ProgramHint{}

  cs, p, pe, eof := 0, 0, len(data), len(data)
  mark := 0

  currentHint := ProgramHint{}

  %%{
    action mark                        { mark = p }
    action debug                       { fmt.Println(data[mark:p]) }

    action register_program_id       { currentHint.ID, _ = strconv.Atoi(data[mark:p]) }
    action register_child_program_id {
      child, _ := strconv.Atoi(data[mark:p])
      currentHint.Children = append(currentHint.Children, child)
    }

    action new_entry {
      hints = append(hints, currentHint)
      currentHint = ProgramHint{}
    }

    program_id = ( digit )+ >mark;

    entry = (
      program_id %register_program_id ' <-> '
      ( program_id %register_child_program_id ', ' )*
      program_id %register_child_program_id
    ) %new_entry;

    main := ( entry '\n' )* entry;

    write init;
    write exec;
  }%%

  if eof != p {
    return []ProgramHint{}, errors.New("parse error")
  }

  return hints, nil
}

