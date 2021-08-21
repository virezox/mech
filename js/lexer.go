package js

import (
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/js"
)

// godocs.io/github.com/tdewolff/parse/v2/js#Lexer
type Lexer struct {
   *js.Lexer
   *parse.Input
}

// godocs.io/github.com/tdewolff/parse/v2/js#NewLexer
func NewLexer(b []byte) Lexer {
   z := parse.NewInputBytes(b)
   return Lexer{
      js.NewLexer(z), z,
   }
}

// godocs.io/net/url#Values
func (l Lexer) Values() map[string][]byte {
   vals := make(map[string][]byte)
   for {
      var k string
      // state 1: break if EqToken
      var ident bool
      for {
         if tt, data := l.Next(); tt == js.ErrorToken {
            return vals
         } else if tt == js.EqToken {
            break
         } else if tt == js.IdentifierToken || tt == js.DotToken {
            if ident {
               k += string(data)
            } else {
               k = string(data)
               ident = true
            }
         } else {
            ident = false
         }
      }
      // state 2: break if SemicolonToken
      for {
         tt, data := l.Next()
         if tt == js.SemicolonToken {
            break
         } else if tt == js.WhitespaceToken {
            continue
         } else if tt == js.DivToken {
            if tt, data = l.RegExp(); tt == js.ErrorToken {
               l.Rewind(0)
               tt, data = l.Next()
            }
         }
         vals[k] = append(vals[k], data...)
      }
   }
}
