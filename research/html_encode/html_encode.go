package encode

import (
   "bytes"
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/html"
   "io"
   "os"
)

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
   z := html.NewLexer(parse.NewInput(r))
   for {
      t, raw := z.Next()
      switch t {
      case html.ErrorToken:
         return nil
      case html.TextToken:
         if bytes.TrimSpace(raw) == nil {
            continue
         }
      }
      os.Stdout.Write(append(raw, '\n'))
   }
}
