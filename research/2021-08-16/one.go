package decode

import (
   "bytes"
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/html"
   "io"
)

type Attribute struct {
   Key []byte
   Val []byte
}

type Decoder struct {
   *html.Lexer
   Data []byte
   Attr []Attribute
}

func NewDecoder(r io.Reader) Decoder {
   return Decoder{
      Lexer: html.NewLexer(parse.NewInput(r)),
   }
}

func (d Decoder) Attribute(key string) []byte {
   for _, a := range d.Attr {
      if string(a.Key) == key {
         return bytes.Trim(a.Val, `'"`)
      }
   }
   return nil
}

// Move to the next text node. Set "Data" to the content, and set "Attr" to
// nil.
func (d *Decoder) NextText() bool {
   for {
      t, data := d.Next()
      if t == html.ErrorToken {
         break
      }
      if t == html.TextToken {
         d.Data, d.Attr = data, nil
         return true
      }
   }
   return false
}
