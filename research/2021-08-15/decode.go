package main

import (
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/html"
   "io"
   "os"
)

type Decoder struct {
   *html.Lexer
}

func NewDecoder(r io.Reader) Decoder {
   return Decoder{
      html.NewLexer(parse.NewInput(r)),
   }
}

func main() {
   f, err := os.Open("index.html")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   NewDecoder(f)
}
