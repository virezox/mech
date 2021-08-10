package main

import (
   "fmt"
   "golang.org/x/net/html"
   "io"
   "os"
)

func token(r io.Reader) []string {
   var data []string
   z := html.NewTokenizer(r)
   for {
      if z.Next() == html.ErrorToken {
         break
      }
      for _, a := range z.Token().Attr {
         if a.Key == "type" && a.Val == "application/ld+json" {
            z.Next()
            data = append(data, z.Token().Data)
         }
      }
   }
   return data
}
