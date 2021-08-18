package html

import (
   "bytes"
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/html"
   "io"
)

var end = []byte{'\n'}

var void = map[string]bool{
   "br": true,
   "img": true,
   "input": true,
   "link": true,
   "meta": true,
}

func write(w io.Writer, b ...[]byte) error {
   for _, s := range b {
      _, err := w.Write(s)
      if err != nil {
         return err
      }
   }
   return nil
}

type Lexer struct {
   *html.Lexer
}

func NewLexer(r io.Reader) Lexer {
   z := parse.NewInput(r)
   return Lexer{
      html.NewLexer(z),
   }
}

func (l Lexer) Render(w io.Writer, indent string) error {
   var ind []byte
   for {
      var err error
      switch t, data := l.Next(); t {
      case html.StartTagToken:
         err = write(w, ind, data)
         if !void[string(l.Text())] {
            ind = append(ind, indent...)
         }
      case html.AttributeToken:
         err = write(w, data)
      case html.StartTagCloseToken:
         err = write(w, data, end)
      case html.TextToken:
         if bytes.TrimSpace(data) != nil {
            err = write(w, ind, data, end)
         }
      case html.EndTagToken:
         ind = ind[len(indent):]
         err = write(w, ind, data, end)
      case html.ErrorToken:
         return nil
      }
      if err != nil {
         return err
      }
   }
}
