package pbs

import (
   "fmt"
   "net/url"
   "testing"
)

const widget = "https://player.pbs.org/widget/partnerplayer/3016754074/"

func TestWidget(t *testing.T) {
   addr, err := url.Parse(widget)
   if err != nil {
      t.Fatal(err)
   }
   wid, err := NewWidget(addr)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", wid)
}
