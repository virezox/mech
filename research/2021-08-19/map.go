package js
import "github.com/tdewolff/parse/v2/js"

type values map[string][]byte

func (l lexer) values() values {
   var key string
   val := make(values)
   for {
      switch tt, data := l.Next(); tt {
      case js.IdentifierToken:
         key = string(data)
      case
      js.CloseBraceToken,
      js.CloseBracketToken,
      js.ColonToken,
      js.CommaToken,
      js.DecimalToken,
      js.FalseToken,
      js.NullToken,
      js.OpenBraceToken,
      js.OpenBracketToken,
      js.StringToken,
      js.TrueToken:
         val[key] = append(val[key], data...)
      case js.ErrorToken:
         return val
      }
   }
}
