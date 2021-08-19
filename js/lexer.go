package js

import (
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/js"
)

// pkg.go.dev/github.com/tdewolff/parse/v2/js#Lexer
type Lexer struct {
   *js.Lexer
}

// pkg.go.dev/github.com/tdewolff/parse/v2/js#NewLexer
func NewLexer(b []byte) Lexer {
   return Lexer{
      js.NewLexer(parse.NewInputBytes(b)),
   }
}

// pkg.go.dev/net/url#Values
func (l Lexer) Values() map[string][]byte {
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
