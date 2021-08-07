package main

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech/musicbrainz"
   "github.com/89z/mech/youtube"
   "os"
   "time"
)

func variance(data []float64) float64 {
   var count, mean, M2 float64
   for _, x := range data {
      count += 1
      delta := x - mean
      mean += delta / count
      M2 += delta * (x - mean)
   }
   return M2 / (count - 1)
}

func distance(x, y int, v []float64, data [][]float64) float64 {
   var sum float64
   for i, r := range data {
      delta := r[x] - r[y]
      sum += delta * delta / v[i]
   }
   return sum
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
