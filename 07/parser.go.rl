// -*-go-*-
package main

import (
  "errors"
  "fmt"
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
  amt := 0

  %%{
    action tok_start              { amt, mark = 0, p }
    action tok_char               { amt++ }

    action program_name_end       { fmt.Println("Program name", data[mark:p]) }
    action program_id_end         { fmt.Println("Program ID", data[mark:p]) }
    action child_program_name_end { fmt.Println("Child program name", data[mark:p]) }

    action debug                  { fmt.Printf("mark: %d\np: %d\npe: %d\n", mark, p, pe) }

    str_token = ( lower* ) >tok_start @tok_char;
    int_token = ( digit* ) >tok_start @tok_char;

    program_id = ( int_token ) %program_id_end;
    child_list = ( str_token ', '? )* %child_program_name_end;

    entry = ( str_token ) %program_name_end ' (' program_id ')' ' -> '? child_list? '\n'?;

    main := ( entry '\n' )*;

    write init;
    write exec;
  }%%

  if eof != p {
    return []NodeHint{}, errors.New("parse error")
  }

  return hints, nil
}

