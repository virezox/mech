package decode

import (
   "github.com/tdewolff/parse/v2/html"
)

type Attribute struct {
   Key string
   Val string
}

type Decoder struct {
   *html.Lexer
   Data string
   Attr []Attribute
}
