package main

import (
   "fmt"
   "golang.org/x/net/html"
   "os"
   "strings"
)

func main() {
   f, err := os.Open("nyt.html")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   z := html.NewTokenizer(f)
   var indent string
   for {
      tt := z.Next()
      if tt == html.ErrorToken {
         break
      }
      if tt == html.EndTagToken {
         indent = indent[1:]
      }
      t := z.Token().String()
      if tt == html.TextToken && strings.TrimSpace(t) == "" {
         continue
      }
      fmt.Print(indent, t, "\n")
      if tt == html.StartTagToken {
         indent += " "
      }
   }
}
