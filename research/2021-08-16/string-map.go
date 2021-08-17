package decode

import (
   "github.com/tdewolff/parse/v2/html"
   "github.com/tdewolff/parse/v2"
   "io"
)

type Decoder struct {
   *html.Lexer
   Data string
   Attr map[string]string
}

func NewDecoder(r io.Reader) Decoder {
   return Decoder{
      Lexer: html.NewLexer(parse.NewInput(r)),
   }
}
