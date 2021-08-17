package mech

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

func TestHtmlWrite(t *testing.T) {
   w := NewHtmlWriter(os.Stdout)
   w.SetIndent(" ")
   if err := w.ReadFrom(strings.NewReader(htmlWrite)); err != nil {
      t.Fatal(err)
   }
}
