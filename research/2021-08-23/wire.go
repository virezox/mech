package main

import (
   "fmt"
   "google.golang.org/protobuf/encoding/protowire"
)

type result struct {
   PageParams struct {
      VideoID string `protobuf:"2"`
   } `protobuf:"2"`
   Six uint32 `protobuf:"3"`
   OffsetInformation struct {
      PageInfo struct {
         VideoID string `protobuf:"4"`
      } `protobuf:"4"`
   } `protobuf:"6"`
}

func main() {
   // \x12\r\x12\vq5UnT4Ik6KU\x18\x062\x0f\"\r\"\vq5UnT4Ik6KU
   var buf []byte
   buf = protowire.AppendTag(buf, 2, protowire.BytesType)
   buf = protowire.AppendString(buf, "q5UnT4Ik6KU")
   fmt.Printf("%q\n", buf)
}
