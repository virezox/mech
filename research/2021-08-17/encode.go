package encode

import (
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/html"
   "io"
)

var end = []byte{'\n'}

type Encoder struct {
   io.Writer
   tab string
}

func (e Encoder) Encode(r io.Reader) error {
   var tab []byte
   z := html.NewLexer(parse.NewInput(r))
   for {
      var err error
      switch t, data := z.Next(); t {
      case html.StartTagToken:
         err = e.write(tab, data)
         tab = append(tab, e.tab...)
      case html.AttributeToken:
         err = e.write(data)
      case html.StartTagCloseToken:
         err = e.write(data, end)
      case html.TextToken:
         err = e.write(tab, data, end)
      case html.EndTagToken:
         tab = tab[len(e.tab):]
         err = e.write(tab, data, end)
      case html.ErrorToken:
         return nil
      }
      if err != nil {
         return err
      }
   }
}

func (e Encoder) write(b ...[]byte) error {
   for _, s := range b {
      _, err := e.Write(s)
      if err != nil {
         return err
      }
   }
   return nil
}
