package main

import (
   "bytes"
   "fmt"
   pp "google.golang.org/protobuf/testing/protopack"
)

func main() {
   a := []byte("\x12\r\x12\vq5UnT4Ik6KU\x18\x062\x0f\"\r\"\vq5UnT4Ik6KU")
   m := pp.Message{
      pp.Tag{2, pp.BytesType},
      pp.LengthPrefix{
         pp.Tag{2, pp.BytesType},
         pp.String("q5UnT4Ik6KU"),
      },
      pp.Tag{3, pp.VarintType},
      pp.Uvarint(6),
      pp.Tag{6, pp.BytesType},
      pp.LengthPrefix{
         pp.Tag{4, pp.BytesType},
         pp.LengthPrefix{
            pp.Tag{4, pp.BytesType},
            pp.String("q5UnT4Ik6KU"),
         },
      },
   }
   b := m.Marshal()
   fmt.Println(bytes.Equal(a, b))
}
