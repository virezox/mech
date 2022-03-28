package pbs

import (
   "fmt"
   "testing"
)

var addr = &url.URL{
   Scheme: "https",
   Host: "player.pbs.org",
   Path: "/widget/partnerplayer/3016754074/",
}

func TestJSON(t *testing.T) {
   video, err := NewVideoBridge(addr)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", video)
}
