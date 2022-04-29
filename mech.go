package mech

import (
   "bytes"
   "encoding/json"
   "strings"
)

func Clean(path string) string {
   fn := func(r rune) rune {
      switch r {
      case
      '"',
      '*',
      '/',
      ':',
      '<',
      '>',
      '?',
      '\\',
      '|',
      '’': // github.com/PowerShell/PowerShell/issues/16084
         return -1
      }
      return r
   }
   return strings.Map(fn, path)
}

func Encode[T any](value T) (*bytes.Buffer, error) {
   buf := new(bytes.Buffer)
   enc := json.NewEncoder(buf)
   enc.SetIndent("", " ")
   err := enc.Encode(value)
   if err != nil {
      return nil, err
   }
   return buf, nil
}
