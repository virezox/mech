package main

import (
   "encoding/base64"
   "fmt"
   "go.dedis.ch/protobuf"
)

var _ = base64.StdEncoding

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
   var r result
   r.PageParams.VideoID = "q5UnT4Ik6KU"
   r.Six = 6
   r.OffsetInformation.PageInfo.VideoID = "q5UnT4Ik6KU"
   buf, err := protobuf.Encode(&r)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%q\n", buf)
   s := base64.StdEncoding.EncodeToString(buf)
   fmt.Println(s)
}
