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
         } else if tt == js.DivToken && !in {
            tt, data = l.RegExp()
            in = true
         }
         vals[k] = append(vals[k], data...)
      }
   }
}
