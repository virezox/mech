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
   s, err := youtube.NewSearch("oneohtrix point never returnal returnal")
   if err != nil {
      panic(err)
   }
   for _, i := range s.Items() {
      r, err := newResult(t, i)
      if err != nil {
         panic(err)
      }
      fmt.Printf(
         "di %.3f du %.3f s %.3f %v %v\n",
         r.distance, r.duration, r.size, r.VideoID(), r.Title(),
      )
   }
}
