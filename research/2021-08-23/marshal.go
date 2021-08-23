package main

import (
   "fmt"
   "google.golang.org/protobuf/proto"
   "google.golang.org/protobuf/types/known/structpb"
)

func main() {
   person := map[string]interface{}{"month": 12, "day": 31}
   m, err := structpb.NewValue(person)
   if err != nil {
      panic(err)
   }
   b, err := proto.Marshal(m)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%q\n", b)
}
