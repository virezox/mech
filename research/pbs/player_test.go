package pbs

import (
   "fmt"
   "net/url"
   "testing"
)

const videoTest = "https://player.pbs.org/widget/partnerplayer/3016754074/"

func TestVideo(t *testing.T) {
   addr, err := url.Parse(videoTest)
   if err != nil {
      t.Fatal(err)
   }
   video, err := NewVideoBridge(addr)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", video)
}
