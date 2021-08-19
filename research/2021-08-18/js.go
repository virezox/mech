package js

import (
   "github.com/tdewolff/parse/v2"
   "github.com/tdewolff/parse/v2/js"
)

type values map[string][]byte

func newValues(b []byte) values {
   var key string
   val := make(values)
   l := js.NewLexer(parse.NewInputBytes(b))
   for {
      switch tt, data := l.Next(); tt {
      case js.IdentifierToken:
         key = string(data)
      case js.CloseBraceToken, js.ColonToken, js.CommaToken, js.DecimalToken,
      js.OpenBraceToken, js.StringToken:
         val[key] = append(val[key], data...)
      case js.ErrorToken:
         return val
      }
   }
}
