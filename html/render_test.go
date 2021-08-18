package html

import (
   "os"
   "strings"
   "testing"
)

const render = `
<meta charset="utf-8">
<h1>
<a href="/umber">Umber</a>
</h1>
<form target="_blank">
<input name="s" placeholder="search">
</form>
`

func TestRender(t *testing.T) {
   r := strings.NewReader(render)
   err := NewLexer(r).Render(os.Stdout, " ")
   if err != nil {
      t.Fatal(err)
   }
}
