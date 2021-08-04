package main

import (
   "encoding/json"
   "fmt"
   "github.com/89z/mech/musicbrainz"
   "github.com/89z/mech/youtube"
   "math"
   "os"
)

func relativeDifference(x, y float64) float64 {
   if x < y {
      return (y - x) / y
   }
   return (x - y) / x
}

type dimensions []struct {
   p float64
   q float64
}

func (s dimensions) euclideanDistance() float64 {
   var f float64
   for _, e := range s {
      d := relativeDifference(e.p, e.q)
      f += math.Pow(d, 2)
   }
   return f
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
      for index, item := range t.Items {
         dMB := t.Duration()
         dYT, err := item.Duration()
         if err != nil {
            panic(err)
         }
         d := dimensions{
            {
               0, float64(index),
            },
            {
               float64(dMB), float64(dYT),
            },
         }
         fmt.Println(d.euclideanDistance())
      }
   }
}
