package html

import (
   "bytes"
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/html"
   "io"
)

type Encoder struct {
   io.Writer
   tab string
}

func NewEncoder(w io.Writer) Encoder {
   return Encoder{Writer: w}
}

func (e *Encoder) SetIndent(tab string) {
   e.tab = tab
}

func (e Encoder) Encode(r io.Reader) error {
   var tab []byte
   b := new(bytes.Buffer)
   z := html.NewLexer(parse.NewInput(r))
   for {
      t, data := z.Next()
      if t == html.ErrorToken {
         return nil
      }
      if t == html.EndTagToken {
         tab = tab[len(e.tab):]
      }
      if t == html.TextToken && bytes.TrimSpace(data) == nil {
         continue
      }
      b.Write(tab)
      b.Write(data)
      b.WriteByte('\n')
      if _, err := b.WriteTo(e.Writer); err != nil {
         return err
      }
      if t == html.StartTagToken && !void[string(z.Text())] {
         tab = append(tab, e.tab...)
      }
   }
}

var void = map[string]bool{
   "br": true,
   "img": true,
   "input": true,
   "link": true,
   "meta": true,
}
