package mech

import (
   "fmt"
   "golang.org/x/net/html"
   "os"
   "strings"
   "testing"
)

const span = `
<span class='user' id="user">John Doe</span>
`

func TestDecode(t *testing.T) {
   d := html.NewTokenizer(strings.NewReader(span))
   for {
      if d.Next() == html.ErrorToken {
         break
      }
      for _, a := range d.Token().Attr {
         fmt.Printf("%q %q\n", a.Key, a.Val)
      }
   }
}

func TestEncode(t *testing.T) {
   r, err := os.Open("index.html")
   if err != nil {
      panic(err)
   }
   defer r.Close()
   e := NewEncoder(os.Stdout)
   e.SetIndent(" ")
   e.Encode(r)
}
