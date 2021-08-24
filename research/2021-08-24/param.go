package main

import (
   "encoding/base64"
   "fmt"
   "github.com/89z/mech/youtube"
   p "google.golang.org/protobuf/testing/protopack"
)

func encode(m p.Message) string {
   return "youtube.com/results?search_query=autechre&sp=" +
   base64.StdEncoding.EncodeToString(m.Marshal())
}

func main() {
   for k, v := range youtube.Params["FEATURES"] {
      val := encode(v)
      fmt.Print(k, "\n", val, "\n\n")
   }
}
