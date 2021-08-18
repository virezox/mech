package render

import (
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/html"
   "io"
)

var end = []byte{'\n'}

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
   return Lexer{
      html.NewLexer(parse.NewInput(r)),
   }
}

func (l Lexer) Render(w io.Writer, indent string) error {
   var ind []byte
   for {
      var err error
      switch t, data := l.Next(); t {
      case html.StartTagToken:
         err = write(w, ind, data)
         ind = append(ind, indent...)
      case html.AttributeToken:
         err = write(w, data)
      case html.StartTagCloseToken:
         err = write(w, data, end)
      case html.TextToken:
         err = write(w, ind, data, end)
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
