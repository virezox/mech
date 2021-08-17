package html

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

func TestDecode(t *testing.T) {
   d := NewDecoder(strings.NewReader(doc))
   d.NextTag("title")
   fmt.Printf("%s\n", d.Bytes())
   d.NextAttr("rel", "icon")
   fmt.Println(d.Attr("href"))
}
