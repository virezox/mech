package html

import (
   "os"
   "strings"
   "testing"
)

const htmlWrite = `
<h1>
<a href="/umber">Umber</a>
</h1>
<form target="_blank">
<input name="s" placeholder="search">
</form>
`

func TestEncode(t *testing.T) {
   w := NewEncoder(os.Stdout)
   w.SetIndent(" ")
   if err := w.Encode(strings.NewReader(htmlWrite)); err != nil {
      t.Fatal(err)
   }
}
