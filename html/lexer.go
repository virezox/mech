package html

import (
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/html"
   "io"
)

// pkg.go.dev/github.com/tdewolff/parse/v2/html#Lexer
type Lexer struct {
   *html.Lexer
   html.TokenType
   data []byte
   attr map[string]string
}

// pkg.go.dev/github.com/tdewolff/parse/v2/html#NewLexer
func NewLexer(r io.Reader) Lexer {
   z := parse.NewInput(r)
   return Lexer{
      Lexer: html.NewLexer(z),
   }
}
