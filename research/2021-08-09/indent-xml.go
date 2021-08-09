package main

import (
   "encoding/xml"
   "os"
)

func main() {
   f, err := os.Open("outer.html")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   d := xml.NewDecoder(f)
   d.Strict = false
   type node struct {
      Children []node `xml:",any"`
      Text string `xml:",chardata"`
      XMLName xml.Name
   }
   var x node
   d.Decode(&x)
   enc := xml.NewEncoder(os.Stdout)
   enc.Indent("", " ")
   enc.Encode(x)
}
