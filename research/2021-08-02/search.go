package main

import (
   "fmt"
   "github.com/89z/mech/musicbrainz"
   "github.com/89z/mech/youtube"
   "time"
)

func main() {
   musicbrainz.Verbose = true
   youtube.Verbose = true
   d := 4*time.Minute + 42*time.Second
   t := musicbrainz.Track{
      Length: d.Milliseconds(),
   }
   s, err := youtube.NewSearch("oneohtrix point never returnal")
   if err != nil {
      panic(err)
   }
   for _, i := range s.Items() {
      r, err := newResult(t, i)
      if err != nil {
         panic(err)
      }
      fmt.Printf("%.3v %v %v\n", r.distance, r.VideoID(), r.Title())
   }
}
