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
func Parse(data string) ([]NodeHint, error) {
  hints := []NodeHint{}

  cs, p, pe, eof := 0, 0, len(data), len(data)
  mark := 0

  currentHint := NodeHint{}

  %%{
    action mark                        { mark = p }
    action debug                       { fmt.Println(data[mark:p]) }

    action register_program_name       { currentHint.Name = data[mark:p] }
    action register_parent_id          { currentHint.ParentID, _ = strconv.Atoi(data[mark:p]) }
    action register_child_program_name { currentHint.Children = append(currentHint.Children, data[mark:p]) }
    action new_entry {
      hints = append(hints, currentHint)
      currentHint = NodeHint{}
    }

    program_name = ( lower )+ >mark %debug;
    id           = ( digit )+ >mark;

    entry            = ( program_name %register_program_name ' (' id %register_parent_id ')' );
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

