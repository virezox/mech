package mech

import (
   "fmt"
   "strings"
   "testing"
)

const doc = `
<head>
   <title>Umber</title>
   <link rel="icon" href="/umber/media/umber.png">
</head>
`

func TestHtmlRead(t *testing.T) {
   h := NewHtmlReader(strings.NewReader(doc))
   h.NextTag("title")
   fmt.Printf("%s\n", h.Bytes())
   h.NextAttr("rel", "icon")
   fmt.Println(h.Attr("href"))
}
