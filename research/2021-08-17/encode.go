package main

import (
   "bytes"
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/html"
   "io"
   "os"
   "strings"
)

type Encoder struct {
   io.Writer
   indent string
}

func (e Encoder) Encode(r io.Reader) error {
   var indent []byte
   b := new(bytes.Buffer)
   z := html.NewLexer(parse.NewInput(r))
   for {
      t, data := z.Next()
      if t == html.ErrorToken {
         return nil
      }
      if t == html.EndTagToken {
         indent = indent[len(e.indent):]
      }
      if t == html.TextToken && bytes.TrimSpace(data) == nil {
         continue
      }
      b.Write(indent)
      b.Write(data)
      b.WriteByte('\n')
      if _, err := b.WriteTo(e.Writer); err != nil {
         return err
      }
      if t == html.StartTagToken {
         indent = append(indent, e.indent...)
      }
   }
}

func main() {
   e := Encoder{os.Stdout, " "}
   e.Encode(strings.NewReader(`<h1><a href="/umber">Umber</a></h1>`))
}
