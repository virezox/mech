package main

import (
   "bytes"
   "fmt"
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/html"
   "io"
   "os"
)

var (
   _ = fmt.Print
   _ = os.Open
)

type Decoder struct {
   *html.Lexer
   attr map[string]string
}

func NewDecoder(r io.Reader) Decoder {
   return Decoder{
      html.NewLexer(parse.NewInput(r)), make(map[string]string),
   }
}

func (d Decoder) TagSelector(name string) bool {
   return true
}

func (d *Decoder) AttrSelector(key, val string) bool {
   for {
      t, _ := d.Next()
      if t == html.ErrorToken {
         break
      }
      if t == html.StartTagToken {
         d.attr = make(map[string]string)
      }
      if t == html.AttributeToken {
         k, v := d.Text(), bytes.Trim(d.AttrVal(), `'"`)
         d.attr[string(k)] = string(v)
      }
      if t == html.StartTagCloseToken && d.attr[key] == val {
         return true
      }
   }
   return false
}

/*
StartTag "<span"
Attribute " class='user'"
StartTagClose ">"
Text "John Doe"
EndTag "</span>"
*/
const span = "<span class='user'>John Doe</span>"

func main() {
   l := html.NewLexer(parse.NewInputString(span))
   for {
      tt, data := l.Next()
      if tt == html.ErrorToken {
         break
      }
      fmt.Printf("%v %q\n", tt, data)
   }
}
