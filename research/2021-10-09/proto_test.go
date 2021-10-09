package proto

import (
   "fmt"
   "testing"
)

func TestProto(t *testing.T) {
   data := []byte("\x12\x02\b\x01")
   toks := parseUnknown(data)
   fmt.Printf("%+v\n", toks)
}
