package encode

import (
   "os"
   "strings"
   "testing"
)

const doc = `<meta charset="utf-8"><h1><a href="/umber">Umber</a></h1>`

func TestEncode(t *testing.T) {
   e := Encoder{os.Stdout, " "}
   e.Encode(strings.NewReader(doc))
}
