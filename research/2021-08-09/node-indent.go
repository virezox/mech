package main

import (
   "fmt"
   "golang.org/x/net/html"
   "os"
)

func main() {
   f, err := os.Open("outer.html")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   z := html.NewTokenizer(f)
   var indent string
   for {
      switch z.Next() {
      case html.ErrorToken:
         return
      case html.StartTagToken:
         indent += " "
      case html.EndTagToken:
         indent = indent[1:]
      }
      fmt.Print(indent, z.Token(), "\n")
   }
}
