package main

import (
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/js"
   "os"
)

func main() {
   f, err := os.Open("index.js")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   lex := js.NewLexer(parse.NewInput(f))
   for range [99]struct{}{} {
      t, data := lex.Next()
      if t == js.ErrorToken {
         break
      }
      if t == js.WhitespaceToken || t == js.LineTerminatorToken {
         continue
      }
      os.Stdout.Write(append(data, '\n'))
   }
}
