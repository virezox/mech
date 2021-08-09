package main

import (
   "encoding/xml"
   "fmt"
   "io"
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
   for {
      token, err := d.Token()
      if err == io.EOF {
         break
      } else if err != nil {
         panic(err)
      }
      fmt.Printf("%+v\n", token)
   }
}
