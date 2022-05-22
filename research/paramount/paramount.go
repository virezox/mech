package main

import (
   "encoding/xml"
   "fmt"
   "os"
)

func main() {
   file, err := os.Open("ignore.xml")
   if err != nil {
      panic(err)
   }
   defer file.Close()
   med := new(Media)
   if err := xml.NewDecoder(file).Decode(med); err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", med)
}

type Media struct {
   Body struct {
      Seq  struct {
         Video []Video `xml:"video"`
      } `xml:"seq"`
   } `xml:"body"`
}

type Video struct {
   Title string `xml:"title,attr"`
   Src string `xml:"src,attr"`
   Param []struct {
      Name string `xml:"name,attr"`
      Value string `xml:"value,attr"`
   } `xml:"param"`
}
