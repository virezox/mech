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
   vals := make(map[string][]byte)
   for {
      // state 1: break if EqToken
      var (
         k string
         inIdent bool
      )
      for {
         if tt, data := l.Next(); tt == js.ErrorToken {
            return vals
         } else if tt == js.EqToken {
            break
         } else if tt == js.IdentifierToken || tt == js.DotToken {
            if inIdent {
               k += string(data)
            } else {
               k = string(data)
               inIdent = true
            }
         } else {
            inIdent = false
         }
      }
      // state 2: break if !WhitespaceToken
      for {
         if tt, data := l.Next(); tt != js.WhitespaceToken {
            if tt == js.DivToken {
               tt, data = l.RegExp()
            }
            vals[k] = data
            break
         }
      }
      // state 3: break if SemicolonToken
      for {
         tt, data := l.Next()
         if tt == js.SemicolonToken {
            break
         }
         vals[k] = append(vals[k], data...)
      }
   }
}
