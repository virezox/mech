package main

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech/musicbrainz"
   "github.com/89z/mech/youtube"
   "math"
   "os"
   "sort"
   "time"
)

func euclideanDistance(hq1, hq2 int64) (float64, error) {
   var dDiff float64 = 1
   // size difference
   sDiff := 100 * relativeDifference(
      float64(hq1), float64(hq2),
   )
   // return
   return math.Pow(dDiff, 2) + math.Pow(sDiff, 2), nil
}

func relativeDifference(x, y float64) float64 {
   if x < y {
      return (y - x) / y
   }
   return (x - y) / x
}

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
   var tracks []track
   json.NewDecoder(f).Decode(&tracks)
   for _, t := range tracks {
      var diffs []diff
      for _, i := range t.Items {
         // duration MB
         dMB := t.Duration()
         // duration YT
         dYT, err := i.Duration()
         if err != nil {
            panic(err)
         }
         // duration difference
         diffs = append(diffs, diff{
            durationDifference(dMB, dYT), i.Item,
         })
      }
      sort.Slice(diffs, func(a, b int) bool {
         return diffs[a].duration < diffs[b].duration
      })
      fmt.Print("\n", t.Title, "\n")
      for _, d := range diffs {
         fmt.Println(d.duration, d.VideoID(), d.Title())
      }
   }
}

type diff struct {
   duration time.Duration
   youtube.Item
}

type track struct {
   musicbrainz.Track
   Items []item
}

type item struct {
   youtube.Item
   HQ1 int64
   HQ2 int64
}
