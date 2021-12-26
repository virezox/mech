package twitter

import (
   "fmt"
   "testing"
)

const (
   spaceID = "1OdKrBnaEPXKX"
   statusID = 1470124083547418624
)

GET /i/api/1.1/live_video_stream/status/28_1468947428984328193?client=web&use_syndication_guest_id=false&cookie_set_host=twitter.com HTTP/1.1
Host: twitter.com
authorization: Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA
x-guest-token: 1474894925598703618
var guest = &Guest{"1475108770955022337"}

func TestSpace(t *testing.T) {
   space, err := NewSpace(guest, spaceID)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", space)
}

func TestStatus(t *testing.T) {
   stat, err := NewStatus(guest, statusID)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", stat)
}
