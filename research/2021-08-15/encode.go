package main

import (
   "bytes"
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/html"
   "os"
)

func main() {
   f, err := os.Open("index.html")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   lex := html.NewLexer(parse.NewInput(f))
   for {
      t, raw := lex.Next()
      if t == html.ErrorToken {
         break
      }
      raw = bytes.TrimSpace(raw)
      if t == html.TextToken && raw == nil {
         continue
      }
      os.Stdout.Write(append(raw, '\n'))
   }
}
