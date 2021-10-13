package main

import (
   "encoding/base64"
   "fmt"
   "github.com/segmentio/encoding/proto"
)

func main() {
   // CAESAA==
   p := param{SortBy: 1}
   b, err := proto.Marshal(p)
   if err != nil {
      panic(err)
   }
   s := base64.StdEncoding.EncodeToString(b)
   fmt.Println(s)
}

type param struct {
   SortBy uint `protobuf:"varint,1"`
   Filter struct {
      UploadDate uint `protobuf:"varint,1"`
      Type uint `protobuf:"varint,2"`
      Duration uint `protobuf:"varint,3"`
      HD uint `protobuf:"varint,4"`
      Subtitles uint `protobuf:"varint,5"`
      CreativeCommons uint `protobuf:"varint,6"`
      ThreeD uint `protobuf:"varint,7"`
      Live uint `protobuf:"varint,8"`
      Purchased uint `protobuf:"varint,9"`
      FourK uint `protobuf:"varint,14"`
      ThreeSixty uint `protobuf:"varint,15"`
      Location uint `protobuf:"varint,23"`
      HDR uint `protobuf:"varint,25"`
      VR180 uint `protobuf:"varint,26"`
   } `protobuf:"bytes,2"`
}
