package mech

import (
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/html"
   "io"
   "strings"
)

type Decoder struct {
   *html.Lexer
   html.TokenType
   data []byte
   attr map[string]string
}

func NewDecoder(r io.Reader) Decoder {
   return Decoder{
      Lexer: html.NewLexer(parse.NewInput(r)),
   }
}

func (d Decoder) Attr(key string) (string, bool) {
   val, ok := d.attr[key]
   if !ok {
      return "", false
   }
   return strings.Trim(val, `'"`), true
}

func (d *Decoder) Bytes() []byte {
   for {
      switch d.TokenType {
      case html.ErrorToken:
         return nil
      case html.TextToken:
         return d.data
      }
      d.TokenType, d.data = d.Next()
   }
}

func (d *Decoder) NextAttr(key, val string) bool {
   for {
      switch d.TokenType, _ = d.Next(); d.TokenType {
      case html.ErrorToken:
         return false
      case html.StartTagToken:
         d.attr = make(map[string]string)
      case html.AttributeToken:
         d.attr[string(d.Text())] = string(d.AttrVal())
      case html.StartTagCloseToken, html.StartTagVoidToken:
         if v, ok := d.Attr(key); ok && v == val {
            return true
         }
      }
   }
}

func (d *Decoder) NextTag(name string) bool {
   for {
      switch d.TokenType, _ = d.Next(); d.TokenType {
      case html.ErrorToken:
         return false
      case html.StartTagToken:
         if string(d.Text()) == name {
            return true
         }
      }
   }
}

func (d *Decoder) String() string {
   return string(d.Bytes())
}
