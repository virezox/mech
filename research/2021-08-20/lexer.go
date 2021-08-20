package js

import (
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/js"
)

type Lexer struct {
   *js.Lexer
}

func NewLexer(b []byte) Lexer {
   return Lexer{
      js.NewLexer(parse.NewInputBytes(b)),
   }
}

func (l Lexer) Values() map[string][]byte {
   vals := make(map[string][]byte)
   for {
      var in bool
      // state 1: break if EqToken
      var k string
      for {
         if tt, data := l.Next(); tt == js.ErrorToken {
            return vals
         } else if tt == js.EqToken {
            break
         } else if tt == js.IdentifierToken || tt == js.DotToken {
            if in {
               k += string(data)
            } else {
               k = string(data)
               in = true
            }
         } else {
            in = false
         }
      }
      // state 2: break if SemicolonToken
   }
}
