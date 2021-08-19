package js

import (
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/js"
)

type lexer struct {
   *js.Lexer
}

func newLexer(b []byte) lexer {
   return lexer{
      js.NewLexer(parse.NewInputBytes(b)),
   }
}
