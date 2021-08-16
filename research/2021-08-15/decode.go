package main

import (
   "fmt"
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/html"
   "io"
   "strings"
)

type Decoder struct {
   *html.Lexer
}

func NewDecoder(r io.Reader) Decoder {
   return Decoder{
      html.NewLexer(parse.NewInput(r)),
   }
}

const span = `
<span class='user' id="user">John Doe</span>
`

func main() {
   d := NewDecoder(strings.NewReader(span))
   for {
      tt, data := d.Next()
      if tt == html.ErrorToken {
         break
      }
      fmt.Printf("%v %q\n", tt, data)
   }
}

/*
StartTag "<span"
Attribute " class='user'"
Attribute " id=\"user\""
StartTagClose ">"

func Trim(s []byte, cutset string) []byte
*/
