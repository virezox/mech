package mech

import (
   "bytes"
   "golang.org/x/net/html"
   "io"
)

var VoidElement = map[string]bool{
   "br": true,
   "img": true,
   "input": true,
   "link": true,
   "meta": true,
}

type Encoder struct {
   io.Writer
   indent string
}

func NewEncoder(w io.Writer) Encoder {
   return Encoder{Writer: w}
}

func (e Encoder) Encode(r io.Reader) error {
   var indent []byte
   b := new(bytes.Buffer)
   z := html.NewTokenizer(r)
   for {
      t := z.Next()
      if t == html.ErrorToken {
         break
      }
      if t == html.EndTagToken {
         indent = indent[len(e.indent):]
      }
      raw := z.Raw()
      if t == html.TextToken && bytes.TrimSpace(raw) == nil {
         continue
      }
      b.Write(indent)
      b.Write(raw)
      b.WriteByte('\n')
      _, err := b.WriteTo(e.Writer)
      if err != nil {
         return err
      }
      if t == html.StartTagToken && !VoidElement[z.Token().Data] {
         indent = append(indent, e.indent...)
      }
   }
   return nil
}

func (e *Encoder) SetIndent(indent string) {
   e.indent = indent
}
