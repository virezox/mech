package mech

import (
   "bytes"
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/html"
   "io"
)

var void = map[string]bool{
   "br": true,
   "img": true,
   "input": true,
   "link": true,
   "meta": true,
}

type HtmlWriter struct {
   io.Writer
   indent string
}

func NewHtmlWriter(w io.Writer) HtmlWriter {
   return HtmlWriter{Writer: w}
}

func (h *HtmlWriter) SetIndent(indent string) {
   h.indent = indent
}

func (h HtmlWriter) ReadFrom(r io.Reader) error {
   var indent []byte
   b := new(bytes.Buffer)
   z := html.NewLexer(parse.NewInput(r))
   for {
      t, data := z.Next()
      if t == html.ErrorToken {
         return nil
      }
      if t == html.EndTagToken {
         indent = indent[len(h.indent):]
      }
      if t == html.TextToken && bytes.TrimSpace(data) == nil {
         continue
      }
      b.Write(indent)
      b.Write(data)
      b.WriteByte('\n')
      if _, err := b.WriteTo(h.Writer); err != nil {
         return err
      }
      if t == html.StartTagToken && !void[string(z.Text())] {
         indent = append(indent, h.indent...)
      }
   }
}
