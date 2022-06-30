package mech
// github.com/89z

import (
   "bytes"
   "encoding/json"
)

func Encode(value any) (*bytes.Buffer, error) {
   buf := new(bytes.Buffer)
   enc := json.NewEncoder(buf)
   enc.SetIndent("", " ")
   err := enc.Encode(value)
   if err != nil {
      return nil, err
   }
   return buf, nil
}
