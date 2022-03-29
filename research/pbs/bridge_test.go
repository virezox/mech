package pbs

import (
   "fmt"
   "net/url"
   "testing"
)

const bridgeTest = "https://player.pbs.org/widget/partnerplayer/3016754074/"

func TestBridge(t *testing.T) {
   address, err := url.Parse(bridgeTest)
   if err != nil {
      t.Fatal(err)
   }
   bridge, err := NewBridge(address)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", bridge)
}
