package main

import (
   "fmt"
   "golang.org/x/net/html"
   "io"
   "os"
)

var (
   _ = fmt.Print
   _ = io.WriteString
)

func main() {
   r, err := os.Open("in.html")
   if err != nil {
      panic(err)
   }
   defer r.Close()
   w, err := os.Create("out.html")
   if err != nil {
      panic(err)
   }
   defer w.Close()
   z := html.NewTokenizer(r)
   for {
      tt := z.Next()
      if tt == html.ErrorToken {
         break
      }
      // this breaks with a single write for some reason
      w.Write(z.Raw())
      w.WriteString("\n")
   }
}
