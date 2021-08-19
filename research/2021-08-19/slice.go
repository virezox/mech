package js
import "github.com/tdewolff/parse/v2/js"

type declaration struct {
   key, val []byte
}

func (l lexer) declarations() []declaration {
   i := -1
   var decs []declaration
   for {
      switch tt, data := l.Next(); tt {
      case js.IdentifierToken:
         decs = append(decs, declaration{key: data})
         i++
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
         decs[i].val = append(decs[i].val, data...)
      case js.ErrorToken:
         return decs
      }
   }
}
