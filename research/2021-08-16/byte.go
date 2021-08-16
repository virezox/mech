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

// Move to the next element with the given attribute. Set "Attr" to the element
// attributes.
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
            d.Text(), d.AttrVal(),
         })
      }
      if t == html.StartTagCloseToken {
         if v := d.Attribute(key); string(v) == val {
            return true
         }
      }
   }
   return false
}

// Move to the next element with the given tag. Set "Data" to the tag name, and
// set "Attr" to nil.
func (d *Decoder) NextTag(name string) bool {
   for {
      t, _ := d.Next()
      if t == html.ErrorToken {
         break
      }
      if t == html.StartTagToken {
         if d.Data = d.Text(); string(d.Data) == name {
            return true
         }
      }
   }
   return false
}

// This needs to be a separate function, as sometimes we are coming from tag,
// and sometimes we are coming from attribute.
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
