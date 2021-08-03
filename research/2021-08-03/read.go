package main

import (
   "github.com/89z/mech/musicbrainz"
   "github.com/89z/mech/youtube"
   "math"
)

func relativeDifference(x, y float64) float64 {
   if x < y {
      return (y - x) / y
   }
   return (x - y) / x
}

func diff(t musicbrainz.Track, i youtube.Item) (float64, error) {
   // duration MB
   dMB := t.Duration()
   // duration YT
   dYT, err := i.Duration()
   if err != nil {
      return 0, err
   }
   // duration difference
   dDiff := 100 * relativeDifference(
      float64(dMB), float64(dYT),
   )
   // image HQ1
   p := youtube.Picture{Base: "hq1", Format: youtube.JPG}
   hq1, err := size(p, i)
   if err != nil {
      return 0, err
   }
   // image HQ2
   p = youtube.Picture{Base: "hq2", Format: youtube.JPG}
   hq2, err := size(p, i)
   if err != nil {
      return 0, err
   }
   // size difference
   sDiff := 100 * relativeDifference(
      float64(hq1), float64(hq2),
   )
   // return
   return math.Pow(dDiff, 2) + math.Pow(sDiff, 2), nil
}
