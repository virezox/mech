package decode

import (
   "github.com/tdewolff/parse/v2/html"
   "github.com/tdewolff/parse/v2"
   "io"
   "strings"
)

type Decoder struct {
   *html.Lexer
   Data string
   Attr map[string]string
}

func NewDecoder(r io.Reader) Decoder {
   return Decoder{
      Lexer: html.NewLexer(parse.NewInput(r)),
   }
}

func (d Decoder) Attribute(key string) (string, bool) {
   val, ok := d.Attr[key]
   if !ok {
      return "", false
   }
   return strings.Trim(val, `'"`), true
}

func (d *Decoder) NextAttr(key, val string) bool {
   for {
      switch t, _ := d.Next(); t {
      case html.ErrorToken:
         return false
      case html.StartTagToken:
         d.Attr = nil
      case html.AttributeToken:
         d.Attr[string(d.Text())] = string(d.AttrVal())
      case html.StartTagCloseToken:
         if v, ok := d.Attribute(key); ok && v == val {
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
         if d.Data = string(d.Text()); d.Data == name {
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
         d.Data = string(data)
         return true
      }
   }
   return false
}
