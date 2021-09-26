package main

import (
   "fmt"
   "github.com/89z/parse/html"
   "os"
)

func main() {
   f, err := os.Open("index.html")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   l := html.NewLexer(f)
   for l.NextTag("script") {
      fmt.Printf("%q\n\n", l.Bytes())
   }
}
