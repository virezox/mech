package decode

import (
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/html"
   "io"
   "strings"
)

func (d *Decoder) Data() string {
   for {
      switch d.TokenType {
      case html.ErrorToken:
         return ""
      case html.TextToken:
         return string(d.Text())
      }
      d.TokenType, _ = d.Next()
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

type Decoder struct {
   *html.Lexer
   html.TokenType
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


func (d *Decoder) NextAttr(key, val string) bool {
   for {
      switch d.TokenType, _ = d.Next(); d.TokenType {
      case html.ErrorToken:
         return false
      case html.StartTagToken:
         d.attr = nil
      case html.AttributeToken:
         d.attr[string(d.Text())] = string(d.AttrVal())
      case html.StartTagCloseToken:
         if v, ok := d.Attr(key); ok && v == val {
            return true
         }
      }
   }
}

