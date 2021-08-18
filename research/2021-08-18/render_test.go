package render

import (
   "os"
   "strings"
   "testing"
)

const doc = `<meta charset="utf-8"><h1><a href="/umber">Umber</a></h1>`

func TestRender(t *testing.T) {
   r := strings.NewReader(doc)
   err := NewLexer(r).Render(os.Stdout, " ")
   if err != nil {
      t.Fatal(err)
   }
}
