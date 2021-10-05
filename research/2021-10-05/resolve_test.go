package resolve
import "testing"

type test struct {
   in string
   typ rune
   id int
}

var tests = []test{
   {"1\tr:[\"nilZ0a1670971920x1633449063\"]\tt:1633449063", 'a', 1670971920},
   {"1\tr:[\"nilZ0t2809477874x1633469972\"]\tt:1633469972", 't', 2809477874},
}

func TestResolve(t *testing.T) {
   for _, test := range tests {
      typ, id, err := tralbum(test.in)
      if err != nil {
         t.Fatal(err)
      }
      if typ != test.typ {
         t.Fatal(typ)
      }
      if id != test.id {
         t.Fatal(id)
      }
   }
}
