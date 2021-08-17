package decode

import (
   "bytes"
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/html"
   "io"
)

type Attribute struct {
   Key, Val []byte
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

func (d Decoder) Attribute(key string) ([]byte, bool) {
   for _, a := range d.Attr {
      if string(a.Key) == key {
         return bytes.Trim(a.Val, `'"`), true
      }
   }
   return nil, false
}

func (d *Decoder) NextAttr(key, val string) bool {
   for {
      switch t, _ := d.Next(); t {
      case html.ErrorToken:
         return false
      case html.StartTagToken:
         d.Attr = nil
      case html.AttributeToken:
         d.Attr = append(d.Attr, Attribute{
            d.Text(), d.AttrVal(),
         })
      case html.StartTagCloseToken:
         if v, ok := d.Attribute(key); ok && string(v) == val {
            return true
         }
      }
   }
   return false
}

func (d *Decoder) NextTag(name string) bool {
   for {
      switch t, _ := d.Next(); t {
      case html.ErrorToken:
         return false
      case html.StartTagToken:
         if d.Data = d.Text(); string(d.Data) == name {
            return true
         }
      }
   }
   return false
}

func (d *Decoder) NextText() bool {
   for {
      switch t, data := d.Next(); t {
      case html.ErrorToken:
         return false
      case html.TextToken:
         d.Data = data
         return true
      }
   }
   return false
}
