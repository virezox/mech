package js

import (
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/js"
)

type Lexer struct {
   *js.Lexer
   *parse.Input
}

func NewLexer(b []byte) Lexer {
   z := parse.NewInputBytes(b)
   return Lexer{
      js.NewLexer(z), z,
   }
}

func (l Lexer) Values() map[string][]byte {
   vals := make(map[string][]byte)
   for {
      var (
         in bool
         k string
      )
      // state 1: break if EqToken
      for {
         tt, data := l.Next()
         if tt == js.ErrorToken {
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
      for {
         tt, data := l.Next()
         if tt == js.SemicolonToken {
            break
         }
         if tt == js.WhitespaceToken {
            continue
         }
         if tt == js.DivToken {
            tt, data = l.RegExp()
            if tt == js.ErrorToken {
               l.Rewind(0)
               tt, data = l.Next()
            }
         }
         vals[k] = append(vals[k], data...)
      }
   }
}
