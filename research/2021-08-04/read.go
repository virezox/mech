package main

import (
   "encoding/json"
   "github.com/89z/mech/musicbrainz"
   "github.com/89z/mech/youtube"
   "os"
   "time"
)

func durationDifference(x, y time.Duration) time.Duration {
   if x < y {
      return y - x
   }
   return x - y
}

func main() {
   f, err := os.Open("returnal.json")
   if err != nil {
      panic(err)
   }
   defer f.Close()
   var tracks []struct {
      musicbrainz.Track
      Items []struct {
         youtube.Item
         HQ1 int64
         HQ2 int64
      }
   }
   json.NewDecoder(f).Decode(&tracks)
   for _, t := range tracks {
      var points []point
      for i, item := range t.Items {
         dMB := t.Duration()
         dYT, err := item.Duration()
         if err != nil {
            panic(err)
         }
         points = append(points, point{
            i, durationDifference(dMB, dYT),
         })
      }
      // normalize
   }
}

type point struct {
   index int
   duration time.Duration
}
