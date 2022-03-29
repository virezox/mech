package pbs

import (
   "fmt"
   "fmt"
   "net/url"
   "testing"
   "time"
)

const playerTest = "https://player.pbs.org/widget/partnerplayer/3016754074/"

func TestVideo(t *testing.T) {
   addr, err := url.Parse(playerTest)
   if err != nil {
      t.Fatal(err)
   }
   video, err := NewVideoBridge(addr)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", video)
}
