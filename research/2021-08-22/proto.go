package main

import (
   "fmt"
   "google.golang.org/protobuf/encoding/protowire"
)

func main() {
   b := protowire.AppendString(nil, "q5UnT4Ik6KU")
   fmt.Printf("%q\n", b)
}
