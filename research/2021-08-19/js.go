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

func (l lexer) values() map[string][]byte {
   var k string
   vals := make(map[string][]byte)
   for {
      switch tt, data := l.Next(); tt {
      case js.IdentifierToken:
         k = string(data)
      case
      js.CloseBraceToken,
      js.CloseBracketToken,
      js.ColonToken,
      js.CommaToken,
      js.DecimalToken,
      js.FalseToken,
      js.NullToken,
      js.OpenBraceToken,
      js.OpenBracketToken,
      js.StringToken,
      js.TrueToken:
         vals[k] = append(vals[k], data...)
      case js.ErrorToken:
         return vals
      }
   }
}
