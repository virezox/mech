package decode

import (
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/html"
   "io"
   "strings"
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

func NewDecoder(r io.Reader) Decoder {
   return Decoder{
      Lexer: html.NewLexer(parse.NewInput(r)),
   }
}

func (d Decoder) Attribute(key string) string {
   for _, a := range d.Attr {
      if a.Key == key {
         return strings.Trim(a.Val, `'"`)
      }
   }
   return ""
}

func (d *Decoder) NextAttr(key, val string) bool {
   for {
      t, _ := d.Next()
      if t == html.ErrorToken {
         break
      }
      if t == html.StartTagToken {
         d.Attr = nil
      }
      if t == html.AttributeToken {
         d.Attr = append(d.Attr, Attribute{
            string(d.Text()), string(d.AttrVal()),
         })
      }
      if t == html.StartTagCloseToken {
         if v := d.Attribute(key); v == val {
            return true
         }
      }
   }
   return false
}

func (d *Decoder) NextTag(name string) bool {
   for {
      t, _ := d.Next()
      if t == html.ErrorToken {
         break
      }
      if t == html.StartTagToken {
         if d.Data = string(d.Text()); d.Data == name {
            return true
         }
      }
   }
   return false
}
