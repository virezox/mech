package encode

import (
   "bytes"
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/html"
   "io"
)

var Void = map[string]bool{
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

func (e *Encoder) SetIndent(indent string) {
   e.indent = indent
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
      if t == html.StartTagToken && !Void[string(z.Text())] {
         indent = append(indent, e.indent...)
      }
   }
}
