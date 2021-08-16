package main

import (
   "bytes"
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/html"
   "io"
   "os"
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

// Move to the next element with the given tag. Set "Data" to the tag name, and
// set "Attr" to the element attributes.
func (d *Decoder) NextTag(name string) bool {
   return d.element(tag{name})
}

// Move to the next text node. Set "Data" to the content, and set "Attr" to
// nil.
func (d *Decoder) NextText() bool {
   for {
      t, b := d.Next()
      if t == html.ErrorToken {
         break
      }
      if t == html.TextToken {
         d.Data, d.Attr = string(b), nil
         return true
      }
   }
   return false
}

type attribute struct {
   key string
   value string
}

func main() {
   f, err := os.Open("index.html")
   if err != nil {
      panic(err)
   }
   defer f.Close()
}

type tag struct {
   name string
}

// Move to the next element with the given attribute. Set "Data" to the tag
// name, and set "Attr" to the element attributes.
func (d *Decoder) NextAttr(key, val string) bool {
   return d.element(attribute{key, val})
}

////////////////////////////////////////////////////////////////////////////////

type exiter interface {
   exit(*Decoder) bool
}

func (n tag) exit(d *Decoder) bool {
   return d.Data == n.name
}

func (a attribute) exit(d *Decoder) bool {
   value, ok := d.Attr[a.key]
   return ok && value == a.value
}

func (d *Decoder) element(e exiter) bool {
   for {
      t, _ := d.Next()
      if t == html.ErrorToken {
         break
      }
      text := string(d.Text())
      if t == html.StartTagToken {
         d.Data, d.Attr = text, make(map[string]string)
      }
      if t == html.AttributeToken {
         d.Attr[text] = string(bytes.Trim(d.AttrVal(), `'"`))
      }
      if t == html.StartTagCloseToken {
         if e.exit(d) {
            return true
         }
      }
   }
   return false
}
