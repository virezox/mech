package main

import (
    "encoding/xml"
    "fmt"
)

func main() {
   f, err := os.Open("inner.html")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   type node struct {
      Attr     []xml.Attr
      XMLName  xml.Name
      Children []node `xml:",any"`
      Text     string `xml:",chardata"`
   }
   var x node
   xml.NewDecoder(f).Decode(&x)
   fmt.Println(x)
   //buf, _ := xml.MarshalIndent(x, "", "\t")
}
