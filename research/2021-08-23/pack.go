package main

import (
   "encoding/base64"
   "fmt"
   p "google.golang.org/protobuf/testing/protopack"
)

func main() {
   m := p.Message{
      p.Tag{2, p.BytesType}, p.LengthPrefix{
         p.Tag{2, p.BytesType}, p.String("q5UnT4Ik6KU"),
      },
      p.Tag{3, p.VarintType}, p.Varint(6),
      p.Tag{6, p.BytesType}, p.LengthPrefix{
         p.Tag{4, p.BytesType}, p.LengthPrefix{
            p.Tag{4, p.BytesType}, p.String("q5UnT4Ik6KU"),
         },
      },
   }
   s := base64.StdEncoding.EncodeToString(m.Marshal())
   fmt.Println(s)
}
