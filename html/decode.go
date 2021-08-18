package html

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

// textContent
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

// getAttribute
func (d Decoder) GetAttr(key string) string {
   val := d.attr[key]
   return strings.Trim(val, `'"`)
}

// hasAttribute
func (d Decoder) HasAttr(key string) bool {
   _, ok := d.attr[key]
   return ok
}

// getElementsByClassName
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
         if d.HasAttr(key) && d.GetAttr(key) == val {
            return true
         }
      }
   }
}

// getElementsByTagName
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
